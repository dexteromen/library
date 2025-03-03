package routes

import (
	"library/controllers"
	"library/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	/*
		// r.GET("/user", controllers.GetAllUsers)
		// r.POST("/user", controllers.CreateUser)
		// r.GET("/user/:id", controllers.GetUserById)
		// r.PUT("/user/:id", controllers.UserUpdateDetails)

		// userGroup := r.Group("/todos")
		// {
		// 	// Todo Routes
		// 	userGroup.POST("", controllers.TodosCreate)       // Create
		// 	userGroup.GET("", controllers.TodosIndex)         // Read all
		// 	userGroup.GET("/:id", controllers.TodosShow)      // Read One
		// 	userGroup.PUT("/:id", controllers.TodosUpdate)    // Update
		// 	userGroup.DELETE("/:id", controllers.TodosDelete) // Delete
		// }

		// Routes for BookInventory CRUD operations
		// r.GET("/books", controllers.GetAllBooks)
		// r.GET("/books/:isbn", controllers.GetBookByISBN)
		// r.POST("/books", controllers.CreateBook)
		// r.PUT("/books/:isbn", controllers.UpdateBook)
		// r.DELETE("/books/:isbn", controllers.DeleteBook)

		// r.GET("/libraries", controllers.GetAllLibraries)
		// r.GET("/users", controllers.GetAllUsers)
		// r.POST("/users", controllers.CreateUser)
		// r.POST("/requests", controllers.CreateRequest)
		// r.POST("/issue", controllers.IssueBook)

		// Register routes
		// router.POST("/users", controllers.CreateUser)
		// router.GET("/users", controllers.GetAllUsers) //New Created
		// router.GET("/users/:id", controllers.GetUser)
		// router.PUT("/users/:id", controllers.UpdateUser)
		// router.DELETE("/users/:id", controllers.DeleteUser)

		// router.POST("/libraries", controllers.CreateLibrary)
		// router.GET("/libraries", controllers.GetAllLibrary) //New created
		// router.GET("/libraries/:id", controllers.GetLibrary)
		// router.PUT("/libraries/:id", controllers.UpdateLibrary)
		// router.DELETE("/libraries/:id", controllers.DeleteLibrary)

	*/
	
	/*
	// router.POST("/books", controllers.CreateBook)
	// router.GET("/books/:isbn", controllers.GetBook)
	// router.PUT("/books/:isbn", controllers.UpdateBook)
	// router.DELETE("/books/:isbn", controllers.DeleteBook)

	// router.POST("/requestevents", controllers.CreateRequestEvent)

	// router.POST("/issueregistry", controllers.CreateIssueRegistry)

	// usersGroup := router.Group("/users")
	// {
	// 	usersGroup.POST("", controllers.CreateUser)       // Create
	// 	usersGroup.GET("", controllers.GetAllUsers)       // Read all
	// 	usersGroup.GET("/:id", controllers.GetUser)       // Read One
	// 	usersGroup.PUT("/:id", controllers.UpdateUser)    // Update
	// 	usersGroup.DELETE("/:id", controllers.DeleteUser) // Delete
	// }
	// librariesGroup := router.Group("/libraries")
	// {
	// 	librariesGroup.POST("/", controllers.CreateLibrary)
	// 	librariesGroup.GET("/", controllers.GetAllLibrary)
	// 	librariesGroup.GET("/:id", controllers.GetLibrary)
	// 	librariesGroup.PUT("/:id", controllers.UpdateLibrary)
	// 	librariesGroup.DELETE("/:id", controllers.DeleteLibrary)
	// }

	// api := router.Group("/api")
	// {
	// 	api.POST("/signup", controllers.SignUp)
	// 	api.POST("/signin", controllers.SignIn)
	// 	api.GET("/user", middleware.AuthMiddleware(), controllers.UserIndex)
	// 	api.GET("/admin", middleware.AuthMiddleware(), controllers.AdminIndex)
	// }
	*/

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
