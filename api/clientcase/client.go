package clientcase

import (
	"context"
	"fmt"
	"log"
	"time"

	protogen "github.com/yeqown/micro-server-demo/api/protogen"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	_hdl    *handler
	timeout = 5 * time.Second
)

// Initial .
func Initial(rpcAddr string) error {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(rpcAddr, opts)
	if err != nil {
		log.Printf("[Error] could not dial: %s with err: %v", rpcAddr, err)
	}

	_hdl = &handler{
		rpcAddr:          rpcAddr,
		client:           protogen.NewFooClient(conn),
		lastGrpcReqError: err,
	}

	return nil
}

type handler struct {
	rpcAddr          string             // rpc configs
	client           protogen.FooClient // grpc client var
	lastGrpcReqError error              // record last rpc all error
}

func (h *handler) connect() {
	// if h.client != nil && h.lastGrpcReqError != nil {
	// 	opts := grpc.WithInsecure()
	// 	conn, err := grpc.Dial(h.rpcAddr, opts)
	// 	if err != nil {
	// 		log.Printf("[Error] could not dial: %s with err: %v", h.rpcAddr, err)
	// 		return
	// 	}

	// 	log.Printf("usersvc.client.connect called")
	// 	h.client = protogen.NewFooClient(conn)
	// }

	// Changed: do nothing should not retry manually
}

// Echo .
func Echo(in *protogen.FooForm) (*protogen.FooResponse, error) {
	_hdl.connect()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	nCtx := metautils.NiceMD(
		metadata.Pairs("authorization", fmt.Sprintf(
			"%s %v", "schema", "tokenthis")),
	).ToOutgoing(ctx)

	resp, err := _hdl.client.Echo(nCtx, in)
	_hdl.lastGrpcReqError = err
	return resp, err
}
