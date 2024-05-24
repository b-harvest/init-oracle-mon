package grpc

import (
	"context"

	tx "github.com/cosmos/cosmos-sdk/types/tx"
)

func (c *Client) GetTx(ctx context.Context, hash string) (*tx.GetTxResponse, error) {
	resp, err := c.txServiceClient.GetTx(
		ctx,
		&tx.GetTxRequest{
			Hash: hash,
		},
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
