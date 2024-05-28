package app

import (
	"context"
	"fmt"

	"bharvest.io/init-oracle-mon/client/rpc"
	"bharvest.io/init-oracle-mon/store"
	"bharvest.io/init-oracle-mon/utils"

	"github.com/cometbft/cometbft/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	compression "github.com/skip-mev/slinky/abci/strategies/codec"
)

func (app *BaseApp) SubVoteExtension(ctx context.Context) {
	rpcClient, err := rpc.New(app.cfg.General.RPC)
	if err != nil {
		app.chErr <- err
		return
	}
	err = rpcClient.Connect(ctx)
	if err != nil {
		app.chErr <- err
		return
	}
	defer func() {
		err = rpcClient.Terminate(ctx)
		if err != nil {
			app.chErr <- err
		}
	}()


	// Prepare codecs for decoding
	commitCodec := compression.NewCompressionExtendedCommitCodec(
		compression.NewDefaultExtendedCommitCodec(),
		compression.NewZStdCompressor(),
	)

	chNewBlock, err := rpcClient.Subscribe(ctx, "tm.event = 'NewBlock'")
	if err != nil {
		app.chErr <- err
		return
	}
	for newBlock := range chNewBlock {
		state := store.NewState()

		block := newBlock.Data.(types.EventDataNewBlock).Block
		state.Height = block.Height
		utils.Info(fmt.Sprintf("New block height: %d", block.Height))

		// Extended Commit is always the first tx in a block
		commit, err := commitCodec.Decode(block.Txs[0])
		if err != nil {
			app.chErr <- err
			utils.Error(err)

			continue
		}

		// signedCnt for check double sign
		// 1: normal sign
		// except: double sign
		signedCnt := 0
		for _, vote := range commit.Votes {
			addr := vote.Validator.Address
			consAddr, err := sdktypes.Bech32ifyAddressBytes("initvalcons", addr)
			if err != nil {
				app.chErr <- err
				utils.Error(err)

				break
			}

			// Check list: double sign, block sign, oracle sign
			if consAddr == app.cfg.General.ConsAddr {
				signedCnt++

				// BlockIDFlagCommit  BlockIDFlag = 2
				if vote.BlockIdFlag == 2 {
					state.BlockSign = true
					utils.Info(fmt.Sprintf("Oracle block signed detected at %d", block.Height))
				}

				// TODO: We have to check more detail about vote extension
				// like pair and price
				if len(vote.VoteExtension) != 0 {
					state.OracleSign = true
					utils.Info(fmt.Sprintf("Oracle data submitted detected at %d", block.Height))
				}
			}
		}

		if signedCnt > 1 {
			state.OracleDoubleSign = true
			utils.Error(fmt.Errorf("Double sign detected at %d", block.Height))
		}

		if state.BlockSign && state.OracleSign && !state.OracleDoubleSign {
			state.Status = true
		} else {
			utils.Info(fmt.Sprintf("Something wrong with your oracle node : %+v", state))
			state.OracleMissCnt++
		}

		if block.Height % 30 == 0 {
			state.OracleMissCnt = 0
		}

		if state.OracleMissCnt > 10 {
			state.OracleMissCnt = 0
			utils.SendTg(fmt.Sprintf("Something wrong with your oracle node at %d", block.Height))
		}

		err = store.GlobalState.Enqueue(state)
		if err != nil {
			app.chErr <- err
			utils.Error(err)

			continue
		}
	}
}
