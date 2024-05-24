package rpc

import cmthttp "github.com/cometbft/cometbft/rpc/client/http"

type Client struct {
	RPCClient *cmthttp.HTTP
	host string
}
