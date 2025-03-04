package routes

import (
	"library/controllers"
	"library/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/signup", controllers.SignUp)
	router.POST("/signin", controllers.SignIn)
	router.POST("/signout", controllers.SignOut)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/admin", controllers.AdminIndex)
		protected.GET("/user", controllers.UserIndex)
	}
}
