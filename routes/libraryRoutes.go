package routes

import (
	"library/controllers"
	"library/middlewares"

	"github.com/gin-gonic/gin"
)

func libraryRoutes(router *gin.Engine) {

	// //Without auth
	// router.GET("/library", controllers.GetLibraries)
	// router.POST("/library", controllers.CreateLibrary)
	// router.PUT("/library/:id", controllers.UpdateLibrary)
	// router.DELETE("/library/:id", controllers.DeleteLibrary)

	router.GET("/library", controllers.GetLibraries) // Get all Libraries
	libraryGroup := router.Group("/")
	libraryGroup.Use(middlewares.AuthMiddleware())
	{
		libraryGroup.POST("/library", middlewares.RoleMiddleware("admin"), controllers.CreateLibrary)       // Create Library
		libraryGroup.PUT("/library/:id", middlewares.RoleMiddleware("admin"), controllers.UpdateLibrary)    // Update Library
		libraryGroup.DELETE("/library/:id", middlewares.RoleMiddleware("admin"), controllers.DeleteLibrary) // Delete Library
	}

}
