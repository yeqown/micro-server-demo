package grpc

import (
	"context"

	protogen "github.com/yeqown/micro-server-demo/api/protogen"
)

func (s *gRPCServer) Echo(ctx context.Context, form *protogen.FooForm) (*protogen.FooResponse, error) {
	bar := form.Foo
	return &protogen.FooResponse{Bar: bar}, nil
}
