package http

import (
	"net/http"

	"github.com/yeqown/micro-server-demo/global"
	"github.com/yeqown/micro-server-demo/internal/modules/demo/usecase"

	"github.com/gin-gonic/gin"
)

type fooHandler struct {
	uc usecase.FooUsecase
}

// New .
func New() fooHandler {
	return fooHandler{
		uc: usecase.NewFooUsecase(global.Repos.FooRepo),
	}
}

func (hdl fooHandler) Bar(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "bar",
	})
}

type echoForm struct {
	Echo string `form:"echo" binding:"required"`
}

func (hdl fooHandler) Echo(c *gin.Context) {
	var (
		form = new(echoForm)
	)

	if err := c.ShouldBind(form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "param wrong",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": form.Echo})
}
