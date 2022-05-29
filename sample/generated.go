// Code generated by gengshell DO NOT EDIT.
// versions:
//  gshell 0.0.1

package main

import (
	"context"
	"encoding/json"

	rpc "github.com/j-tokumori/gshell/sample/api"

	"github.com/j-tokumori/gshell"
)

// RegisterRPC ...
func RegisterRPC() gshell.Option {
	return gshell.RegisterRPCFactories(
		NewRegister,
	)
}

// Register ...
type Register rpc.RegisterArgs

// Call ...
func (r *Register) Call(ctx context.Context, conn gshell.Conn) (res *gshell.Response, err error) {
	client := rpc.NewAuthServiceClient(conn)
	args := rpc.RegisterArgs(*r)
	res = gshell.NewEmptyResponse()
	res.Reply, err = client.Register(ctx, &args, gshell.ResponseOptions(res)...)
	return res, err
}

// NewRegister ...
func NewRegister(in []byte) gshell.RPC {
	r := &Register{}
	err := json.Unmarshal(in, r)
	if err != nil {
		panic(err)
	}
	return r
}

// RegisterReply ...
func RegisterReply(c *gshell.Client) *rpc.RegisterReply {
	if res := c.Response("Register"); res != nil {
		return res.Reply.(*rpc.RegisterReply)
	}
	return nil
}