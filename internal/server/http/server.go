package http

import (
	"fmt"
	"log"

	"github.com/yeqown/micro-server-demo/internal/repository"
	"github.com/yeqown/micro-server-demo/internal/service"

	"github.com/gin-gonic/gin"
)

type httpServer struct {
	port   int
	engine *gin.Engine
}

// New .
func New(fooRepo repository.FooRepo, port int) *httpServer {
	fooUC := service.NewFooUsecase(fooRepo)

	fooHdl := fooHandler{
		uc: fooUC,
	}

	e := gin.New()

	group := e.Group("/v1")
	{
		group.GET("/foo", fooHdl.Bar)
		group.GET("/echo", fooHdl.Echo)
	}

	return &httpServer{
		port:   port,
		engine: e,
	}
}

func (s *httpServer) Run() error {
	log.Printf("REST http server is running on: %d", s.port)
	return s.engine.Run(fmt.Sprintf(":%d", s.port))
}
