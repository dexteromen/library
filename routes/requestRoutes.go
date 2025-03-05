package routes

import (
	"library/controllers"
	"library/middlewares"

	"github.com/gin-gonic/gin"
)

func requestRoutes(r *gin.Engine) {

	// r.Use(middlewares.AuthMiddleware())

	// requestGroup := r.Group("/requests")
	// {
	// 	requestGroup.POST("/", controllers.CreateRequest)
	// 	requestGroup.PUT("/:id/approve", controllers.ApproveRequest)
	// 	requestGroup.GET("/", controllers.GetRequests)
	// }

	requestGroup := r.Group("/requests")
	requestGroup.Use(middlewares.AuthMiddleware())
	{
		requestGroup.POST("/", controllers.CreateRequest)
		requestGroup.PUT("/:id/approve", middlewares.RoleMiddleware("owner"), controllers.ApproveRequest)
		requestGroup.GET("/", controllers.GetRequests)
	}

}
