package routes

import (
	"library/controllers"
	"library/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Public Routes
	router.POST("/signup", controllers.SignUp)
	router.POST("/signin", controllers.SignIn)

	// Protected Routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware()) // Apply JWT middleware
	{
		protected.GET("/admin", controllers.AdminIndex)
		protected.GET("/user", controllers.UserIndex)
	}
}
