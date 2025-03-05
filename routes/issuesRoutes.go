package routes

import (
	"library/controllers"
	"library/middlewares"

	"github.com/gin-gonic/gin"
)

func issuesRoutes(r *gin.Engine) {

	// r.Use(middlewares.AuthMiddleware())

	// issueGroup := r.Group("/issues")
	// {
	// 	issueGroup.POST("/", controllers.CreateIssue)
	// 	issueGroup.PUT("/:id/return", controllers.ReturnBook)
	// 	issueGroup.GET("/", controllers.GetIssues)
	// }

	issueGroup := r.Group("/issues")
	issueGroup.Use(middlewares.AuthMiddleware())
	{
		issueGroup.POST("/", middlewares.RoleMiddleware("admin"), controllers.CreateIssue)
		issueGroup.PUT("/:id/return", middlewares.RoleMiddleware("approver"), controllers.ReturnBook)
		issueGroup.GET("/", controllers.GetIssues)
	}

}
