package main

import (
	"library/config"
	"library/models"
	"library/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	config.DB.AutoMigrate(
		&models.User{},
	)

	r := gin.Default()

	routes.Routes(r)


	r.Run()
}
