package rpc

import (
	"context"

	"bharvest.io/init-oracle-mon/utils"
	cmthttp "github.com/cometbft/cometbft/rpc/client/http"
)

func New(host string) (*Client, error) {
	result := &Client{
		host: host,
	}
	client, err := cmthttp.New(result.host, "/websocket")
	if err != nil {
		return nil, err
	}
	result.RPCClient = client

	return result, nil
}

func (c *Client) Connect(ctx context.Context) error {
	// For websocket connection
	err := c.RPCClient.Start()
	if err != nil {
		return err
	}

	utils.Info("RPC connected")
	return nil
}

func (c *Client) Terminate(_ context.Context) error {
	// For websocket connection
	err := c.RPCClient.Stop()
	if err != nil {
		return err
	}

	utils.Info("RPC connection terminated")
	return nil
}
