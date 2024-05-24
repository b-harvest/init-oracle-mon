package grpc

import (
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	"google.golang.org/grpc"
)

type Client struct {
	host string
	conn *grpc.ClientConn
	txServiceClient tx.ServiceClient
}
