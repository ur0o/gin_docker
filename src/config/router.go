package config

import (
	"github.com/gin-gonic/gin"

	"gin_docker/handlers"
)

func Init(e *gin.Engine) {
	initRouting(e)
}

func initRouting(e * gin.Engine) {
	e.GET("/", handlers.Index)
}