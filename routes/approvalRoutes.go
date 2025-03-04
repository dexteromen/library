package routes

import (
	"library/controllers"
	// "library/middlewares"

	"github.com/gin-gonic/gin"
)

func issueApprovalRoutes(r *gin.Engine) {
	// // Middleware (if needed)
	// r.Use(middleware.AuthMiddleware())

	requestGroup := r.Group("/requests")
	{
		requestGroup.POST("/", controllers.CreateRequest)
		requestGroup.PUT("/:id/approve", controllers.ApproveRequest)
		requestGroup.GET("/", controllers.GetRequests)
	}

	issueGroup := r.Group("/issues")
	{
		issueGroup.POST("/", controllers.CreateIssue)
		issueGroup.PUT("/:id/return", controllers.ReturnBook)
		issueGroup.GET("/", controllers.GetIssues)
	}

}
