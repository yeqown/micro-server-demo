package router

import (
	"fmt"
	"log"

	demohtp "github.com/yeqown/micro-server-demo/internal/modules/demo/delivery/http"

	"github.com/gin-gonic/gin"
)

type httpServer struct {
	port int
	engi *gin.Engine
}

// NewHTTP .
func NewHTTP(port int) httpServer {
	srv := httpServer{
		port: port,
		engi: gin.New(),
	}

	return srv
}

func (s httpServer) Run() error {
	log.Printf("REST http server is running on: %d", s.port)
	return s.engi.Run(fmt.Sprintf(":%d", s.port))
}

// MountRouters .
func (s httpServer) mountRouters() {
	// use middlewares
	s.engi.Use(gin.Logger())
	s.engi.Use(gin.Recovery())
	// mount routers

	group := s.engi.Group("/v1")
	{
		fooHdl := demohtp.New()
		group.GET("/foo", fooHdl.Bar)
		group.GET("/echo", fooHdl.Echo)
	}
}
