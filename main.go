package main

import (
	"library/config"
	"library/routes"

	_ "library/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.Routes(r)

	r.Run()
}
