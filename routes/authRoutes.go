package routes

import (
	"library/controllers"
	// "library/middlewares"

	"github.com/gin-gonic/gin"
)

func authRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.SignUp)
	router.POST("/signin", controllers.SignIn)
	router.POST("/signout", controllers.SignOut)
	
	router.GET("/users", controllers.GetUsers)
	router.GET("/user/:id", controllers.GetUserById)
	router.PUT("/user/:id", controllers.UpdateUserById)
	router.DELETE("/user/:id", controllers.DeleteUserById)
	
	// router.POST("/forgotPassword", controllers
	// router.POST("/resetPassword", controllers.ResetPassword)
	// router.POST("/changePassword", controllers.ChangePassword)
	// router.POST("/changeRole", controllers.ChangeRole)
	// router.POST("/changeLibID", controllers.ChangeLibID)
	// router.POST("/changeContactNumber", controllers.ChangeContactNumber)
	// router.POST("/changeName", controllers.ChangeName)
	// router.POST("/changeEmail", controllers.ChangeEmail)
	// router.POST("/changePassword", controllers.ChangePassword)

	// protected := router.Group("/")
	// protected.Use(middlewares.AuthMiddleware())
	// {
	// 	protected.GET("/admin", controllers.AdminIndex)
	// 	protected.GET("/user", controllers.UserIndex)
	// }

	// // Protected Routes (Require Authentication)
	// router.Use(middlewares.AuthMiddleware())
}