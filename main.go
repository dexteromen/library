package main

import (
	"github.com/dexteromen/library/initializers"
	"github.com/dexteromen/library/models"
	"github.com/dexteromen/library/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	// initializers.DB.AutoMigrate(&models.Todo{})
	
	initializers.DB.AutoMigrate(
		&models.Library{}, 
		&models.User{}, 
		&models.BookInventory{}, 
		&models.RequestEvent{}, 
		&models.IssueRegistry{},
	)

	r := gin.Default()

	// // Todo Routes
	// routes.TodoRoutes(r)
	routes.Routes(r)

	r.Run()
}
