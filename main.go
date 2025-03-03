package main

import (
	"github.com/dexteromen/library/config"
	"github.com/dexteromen/library/models"
	"github.com/dexteromen/library/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	// config.DB.AutoMigrate(&models.Todo{})

	config.DB.AutoMigrate(
		&models.Library{},
		&models.User{},
		// &models.BookInventory{},
		// &models.RequestEvent{},
		// &models.IssueRegistry{},
	)

	// config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	// // Todo Routes
	// // routes.TodoRoutes(r)

	routes.Routes(r)


	r.Run()
}
