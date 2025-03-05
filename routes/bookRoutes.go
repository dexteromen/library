package routes

import (
	"library/controllers"
	"library/middlewares"

	"github.com/gin-gonic/gin"
)

func bookRoutes(router *gin.Engine) {
	// //Without auth
	// router.GET("/books", controllers.GetBooks)             // Get All Books
	// router.POST("/book", controllers.CreateBook)           // Create Book
	// router.PUT("/book/:id", controllers.UpdateBookByID)    // Update Book
	// router.DELETE("/book/:id", controllers.DeleteBookByID) // Delete Book
	// router.GET("/book/:id", controllers.GetBookByID)       // Get Book By ID

	router.GET("/search", controllers.SearchBooks)   // Search Books
	router.GET("/books", controllers.GetBooks)       // Get All Books
	router.GET("/book/:id", controllers.GetBookByID) // Get Book By ID
	bookGroup := router.Group("/")
	bookGroup.Use(middlewares.AuthMiddleware())
	{
		// Admin Routes
		bookGroup.POST("/book", middlewares.RoleMiddleware("admin"), controllers.CreateBook)           // Create Book
		bookGroup.PUT("/book/:id", middlewares.RoleMiddleware("admin"), controllers.UpdateBookByID)    // Update Book
		bookGroup.DELETE("/book/:id", middlewares.RoleMiddleware("admin"), controllers.DeleteBookByID) // Delete Book

		// // User Routes
		// router.POST("/request", middlewares.RoleMiddleware("user"), controllers.RequestBook)

		// // Approver Routes
		// router.PUT("/approve/:id", middlewares.RoleMiddleware("approver"), controllers.ApproveRequestById)

	}

	// router.POST("/request", controllers.RequestBook)
	// router.POST("/issue", controllers.IssueBook)
	// router.POST("/return", controllers.ReturnBook)
}
