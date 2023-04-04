package config

import (
	"github.com/gin-gonic/gin"

	"gin_docker/handler"
)

func Routing(e *gin.Engine) {
	e.GET("/", handler.Index)
}
