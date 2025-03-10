// package main

// import (
// 	"library/config"
// 	"library/routes"

// 	"github.com/gin-gonic/gin"
// )

// func init() {
// 	config.LoadEnvVariables()
// 	config.ConnectDB()
// }

// func main() {
// 	r := gin.Default()
// 	routes.Routes(r)
// 	r.Run()
// }

package main

import (
	"library/config"
	"library/routes"

	_ "library/docs" // Import generated Swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Library Management API
// @version 1.0
// @description This is a library management system API with JWT authentication.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url http://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api/v1

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	r := gin.Default()

	// Register routes
	routes.Routes(r)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	r.Run()
}
