package utils

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Wallet struct {
	config *sdkTypes.Config
	Acc sdkTypes.AccAddress
	Val sdkTypes.ValAddress
	Cons sdkTypes.ConsAddress
}

func getPrefix(addr string) string {
	i := strings.Index(addr, "1")
	if i == -1 {
		msg := fmt.Sprintf("Address converting error => %s", addr)
		panic(errors.New(msg))
	}
	return addr[:i]
}

func NewWallet(_ context.Context, addr string) (*Wallet, error) {
	prefix := getPrefix(addr)
	config := sdkTypes.NewConfig()
	config.SetBech32PrefixForAccount(prefix, prefix+"pub")
	config.SetBech32PrefixForValidator(prefix+"valoper", prefix+"valoperpub")
	config.SetBech32PrefixForConsensusNode(prefix+"valcons", prefix+"valvalconspub")

	address, err := sdkTypes.GetFromBech32(addr, prefix)
	if err != nil {
		return nil, err
	}
	hex := hex.EncodeToString(address)

	acc, err := sdkTypes.AccAddressFromHexUnsafe(hex)
	if err != nil {
		return nil, err
	}
	val, err := sdkTypes.ValAddressFromHex(hex)
	if err != nil {
		return nil, err
	}
	cons, err := sdkTypes.ConsAddressFromHex(hex)
	if err != nil {
		return nil, err
	}

	wallet := &Wallet{
		config,
		acc,
		val,
		cons,
	}
	return wallet, nil
}

func (w *Wallet) PrintAcc() string {
	sdkTypes.GetConfig().SetBech32PrefixForAccount(w.config.GetBech32AccountAddrPrefix(), w.config.GetBech32AccountPubPrefix())
	return w.Acc.String()
}

func (w *Wallet) PrintValoper() string {
	sdkTypes.GetConfig().SetBech32PrefixForValidator(w.config.GetBech32ValidatorAddrPrefix(), w.config.GetBech32ValidatorPubPrefix())
	return w.Val.String()
}

func (w *Wallet) PrintCons() string {
	sdkTypes.GetConfig().SetBech32PrefixForConsensusNode(w.config.GetBech32ConsensusAddrPrefix(), w.config.GetBech32ConsensusPubPrefix())
	return w.Cons.String()
}
