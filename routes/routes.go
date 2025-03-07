package routes

import (
	"library/controllers"
	"library/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	//Auth Routes
	router.POST("/signup", controllers.SignUp)   // Create User
	router.POST("/signin", controllers.SignIn)   // Login User
	router.POST("/signout", controllers.SignOut) // Logout User

	router.GET("/users", controllers.GetUsers)             // Get all Users
	router.GET("/user/:id", controllers.GetUserById)       // Get User By ID
	router.PUT("/user/:id", controllers.UpdateUserById)    // Update User By ID
	router.DELETE("/user/:id", controllers.DeleteUserById) // Delete User By ID

	// Library Routes
	router.GET("/library", controllers.GetLibraries) // Get all Libraries
	libraryGroup := router.Group("/")
	libraryGroup.Use(middlewares.AuthMiddleware())
	{
		libraryGroup.POST("/library/", controllers.CreateLibrary)                                                    // Create Library
		libraryGroup.PUT("/library/:id", middlewares.RoleMiddleware("owner"), controllers.UpdateLibrary)             // Update Library
		libraryGroup.DELETE("/library/:id", middlewares.RoleMiddleware("owner", "admin"), controllers.DeleteLibrary) // Delete Library
	}

	// Book Routes
	router.GET("/search", controllers.SearchBooks)   // Search Books by Title, Authors, Publisher
	router.GET("/books", controllers.GetBooks)       // Get All Books
	router.GET("/book/:id", controllers.GetBookByID) // Get Book By ID
	bookGroup := router.Group("/")
	bookGroup.Use(middlewares.AuthMiddleware())
	{
		bookGroup.POST("/book", middlewares.RoleMiddleware("owner"), controllers.CreateBook)                    // Create Book By Owner Only Because Owner has LibID
		bookGroup.PUT("/book/:id", middlewares.RoleMiddleware("owner"), controllers.UpdateBookByID)             // Update Book By Owner Only Because Owner has LibID
		bookGroup.DELETE("/book/:id", middlewares.RoleMiddleware("owner", "admin"), controllers.DeleteBookByID) // Delete Book By Owner and Admin Only not by Reader
	}

	// Request Issues Return Routes
	requestAndIssuesGroup := router.Group("/")
	requestAndIssuesGroup.Use(middlewares.AuthMiddleware())
	{
		requestAndIssuesGroup.GET("/requests", controllers.GetRequests)                                                          // Get all requests
		requestAndIssuesGroup.GET("/issues", controllers.GetIssues)                                                              // Get all issues
		requestAndIssuesGroup.POST("/request", middlewares.RoleMiddleware("reader"), controllers.CreateRequest)                  // Create Request
		requestAndIssuesGroup.PUT("/approve-issue/:id", middlewares.RoleMiddleware("admin"), controllers.ApproveAndIssueRequest) // Approve Request and Issue Book
		requestAndIssuesGroup.PUT("/return/:id", middlewares.RoleMiddleware("reader"), controllers.ReturnBook)                   // Return book by isbn
		// requestAndIssuesGroup.PUT("/approve/:id", middlewares.RoleMiddleware("admin"), controllers.ApproveRequest)
	}
}
