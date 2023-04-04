package main

import (
	"github.com/gin-gonic/gin"

	"gin_docker/config"
)

func main() {
	engine := gin.Default()

	config.Routing(engine)

	engine.Run(":8080")
}
