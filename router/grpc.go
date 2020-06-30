package router

import (
	"context"
	"fmt"
	"log"
	"net"

	protogen "github.com/yeqown/micro-server-demo/api/protogen"
	"github.com/yeqown/micro-server-demo/global"
	"github.com/yeqown/micro-server-demo/internal/modules/demo/usecase"

	logger "github.com/yeqown/infrastructure/framework/logrus-logger"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

func (s *gRPCServer) Echo(ctx context.Context, form *protogen.FooForm) (*protogen.FooResponse, error) {
	tokenInfo := ctx.Value(tokenKey)
	log.Printf("get tokenInfo: %v", tokenInfo)

	bar := form.Foo
	return &protogen.FooResponse{Bar: bar}, nil
}

type gRPCServer struct {
	port  int
	fooUC usecase.FooUsecase
}

// NewgRPC .
func NewgRPC(port int) *gRPCServer {
	return &gRPCServer{
		port:  port,
		fooUC: usecase.NewFooUsecase(global.Repos.FooRepo),
	}
}

var (
// logrusLogger = new(logrus.Logger)
// customFunc   grpc_logrus.CodeToLevel
)

func (s *gRPCServer) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	logrusEntry := logger.Log.WithField("service", "micro-server-demo")
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_logrus.Option{
		// grpc_logrus.WithLevels(grpc_logrus.customFunc),
		grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel),
	}
	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	grpcSrv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			// grpc_prometheus.StreamServerInterceptor,
			// grpc_auth.StreamServerInterceptor(myAuthFunction),
			grpc_logrus.StreamServerInterceptor(logrusEntry, opts...),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			// grpc_prometheus.UnaryServerInterceptor,
			// grpc_auth.UnaryServerInterceptor(myAuthFunction),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	reflection.Register(grpcSrv)
	protogen.RegisterFooServer(grpcSrv, s)

	return grpcSrv.Serve(l)
}

func parseToken(token string) (struct{}, error) {
	return struct{}{}, nil
}

func userClaimFromToken(struct{}) string {
	return "foobar"
}

var tokenKey = "tokenInfo"

func myAuthFunction(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "schema")
	if err != nil {
		return nil, err
	}
	tokenInfo, err := parseToken(token)
	if err != nil {
		logger.Log.Warnf("router.myAuthFunction failed to parseToken, err=%v", err)
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	// grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))
	newCtx := context.WithValue(ctx, tokenKey, tokenInfo)
	return newCtx, nil
}
