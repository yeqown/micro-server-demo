package grpc

import (
	"context"
	"log"

	protogen "github.com/yeqown/micro-server-demo/api/protogen"
)

func (s *gRPCServer) Echo(ctx context.Context, form *protogen.FooForm) (*protogen.FooResponse, error) {
	tokenInfo := ctx.Value(tokenKey)
	log.Printf("get tokenInfo: %v", tokenInfo)

	bar := form.Foo
	return &protogen.FooResponse{Bar: bar}, nil
}
