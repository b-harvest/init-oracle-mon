package rpc

import (
	"context"

	cmttypes "github.com/cometbft/cometbft/rpc/core/types"
)

func (c *Client) Subscribe(ctx context.Context, query string) (<-chan cmttypes.ResultEvent, error) {
	resultEvent, err := c.RPCClient.Subscribe(ctx, "subscribe", query)
	if err != nil {
		return nil, err
	}

	return resultEvent, nil
}
