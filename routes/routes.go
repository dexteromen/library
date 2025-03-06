package routes

import (
	"library/controllers"
	"library/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	//Auth Routes
	router.POST("/signup", controllers.SignUp)
	router.POST("/signin", controllers.SignIn)
	router.POST("/signout", controllers.SignOut)

	router.GET("/users", controllers.GetUsers)
	router.GET("/user/:id", controllers.GetUserById)
	router.PUT("/user/:id", controllers.UpdateUserById)
	router.DELETE("/user/:id", controllers.DeleteUserById)

	// Library Routes
	router.GET("/library", controllers.GetLibraries) // Get all Libraries
	libraryGroup := router.Group("/library")
	libraryGroup.Use(middlewares.AuthMiddleware())
	{
		libraryGroup.POST("/", controllers.CreateLibrary)                                                    // Create Library
		libraryGroup.PUT("/:id", middlewares.RoleMiddleware("owner"), controllers.UpdateLibrary)             // Update Library
		libraryGroup.DELETE("/:id", middlewares.RoleMiddleware("owner", "admin"), controllers.DeleteLibrary) // Delete Library
	}

	// Book Routes
	router.GET("/search", controllers.SearchBooks)   // Search Books by Title, Authors, Publisher
	router.GET("/books", controllers.GetBooks)       // Get All Books
	router.GET("/book/:id", controllers.GetBookByID) // Get Book By ID
	bookGroup := router.Group("/book")
	bookGroup.Use(middlewares.AuthMiddleware())
	{
		bookGroup.POST("/", middlewares.RoleMiddleware("owner"), controllers.CreateBook)                   // Create Book By Owner Only Because Owner has LibID
		bookGroup.PUT("/:id", middlewares.RoleMiddleware("owner"), controllers.UpdateBookByID)             // Update Book By Owner Only Because Owner has LibID
		bookGroup.DELETE("/:id", middlewares.RoleMiddleware("owner", "admin"), controllers.DeleteBookByID) // Delete Book By Owner and Admin Only not by Reader
	}

	// Request Routes
	requestGroup := router.Group("/requests")
	requestGroup.Use(middlewares.AuthMiddleware())
	{
		requestGroup.GET("/", controllers.GetRequests)
		requestGroup.POST("/", middlewares.RoleMiddleware("reader"), controllers.CreateRequest)
		requestGroup.PUT("/:id/approve", middlewares.RoleMiddleware("admin"), controllers.ApproveRequest)
	}

	// Issue Routes
	issueGroup := router.Group("/issues")
	issueGroup.Use(middlewares.AuthMiddleware())
	{
		issueGroup.POST("/", middlewares.RoleMiddleware("admin"), controllers.CreateIssue)
		issueGroup.PUT("/:id/return", middlewares.RoleMiddleware("owner"), controllers.ReturnBook)
		issueGroup.GET("/", controllers.GetIssues)
	}
}
