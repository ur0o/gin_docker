package main

import (
	"github.com/gin-gonic/gin"

	"gin_docker/config"
	"gin_docker/database"
	"gin_docker/models"
)

func main() {
	engine := gin.Default()

	database.Init()
	defer database.Close()
	models.Migrate()
	config.Init(engine)
	engine.Static("/assets", "./assets")
	engine.LoadHTMLGlob("templates/*")

	engine.Run(":8080")
}
