package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/programzheng/go-auth/config"
)

var env = config.New()

func RequestDebug(request interface{}) {
	if env.GetString("API_DEBUG") == "true" {
		fmt.Printf("RequestDebug:%v\n", request)
	}
}

func GinBind(c *gin.Context, request interface{}) error {
	if err := c.Bind(&request); err != nil {
		return err
	}
	RequestDebug(request)

	return nil
}
