package http

import (
	"github.com/gin-gonic/gin"
)

type httpServer struct {
	port   int
	engine *gin.Engine
}

func (s *httpServer) Run() {

}
