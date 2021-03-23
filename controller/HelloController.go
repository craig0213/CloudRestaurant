package controller

import (
	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func (hello *HelloController) Router(e *gin.Engine) {
	e.GET("/hello", hello.Hello)
}

func (hello *HelloController) Hello(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"message": "hello cloudrestaurant",
		"context": "test",
	})
}
