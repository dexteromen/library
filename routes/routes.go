package routes

import (
	// "library/config"
	"library/controllers"
	"library/middlewares"

	// "library/models"
	// "library/utils"
	// "net/http"
	// "time"

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

	// readerGroup := router.Group("/")
	// readerGroup.POST("/books/request", RaiseIssueRequest)

	// adminGroup := router.Group("/")
	// adminGroup.GET("/books/requests", ListIssueRequests)
	// adminGroup.PUT("/books/requests/:request_id", ApproveRejectIssueRequest)
}

// type IssueRequest struct {
// 	ID            uint      `gorm:"primaryKey" json:"id"`
// 	BookID        uint      `json:"book_id"`
// 	Email         string    `json:"email"`
// 	Status        string    `json:"status"` // pending, approved, rejected
// 	ApproverEmail string    `json:"approver_email,omitempty"`
// 	ApprovedAt    time.Time `json:"approved_at,omitempty"`
// 	CreatedAt     time.Time `json:"created_at"`
// }
// type IssueRegistry struct {
// 	ID         uint       `gorm:"primaryKey" json:"id"`
// 	BookID     uint       `json:"book_id"`
// 	UserEmail  string     `json:"user_email"`
// 	IssueDate  time.Time  `json:"issue_date"`
// 	ReturnDate *time.Time `json:"return_date,omitempty"` // Null if not returned yet
// }

// func RaiseIssueRequest(c *gin.Context) {
// 	// Get user email from context (set in AuthMiddleware)
// 	email, exists := c.Get("email")
// 	if !exists {
// 		utils.RespondJSON(c, http.StatusUnauthorized, "Unauthorized access", nil)
// 		return
// 	}

// 	// Parse request body
// 	var request struct {
// 		BookID uint `json:"book_id"`
// 	}
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		utils.RespondJSON(c, http.StatusBadRequest, "Invalid request data", nil)
// 		return
// 	}

// 	// Check if book exists and count available copies
// 	var book models.Book
// 	if err := config.DB.Where("id = ?", request.BookID).First(&book).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusNotFound, "Book not found", nil)
// 		return
// 	}

// 	// Count available copies (not issued)
// 	var issuedCount int64
// 	config.DB.Model(&models.IssueRegistry{}).
// 		Where("book_id = ? AND return_date IS NULL", request.BookID).
// 		Count(&issuedCount)

// 	availableCopies := book.TotalCopies - int(issuedCount)
// 	if availableCopies <= 0 {
// 		utils.RespondJSON(c, http.StatusConflict, "Book is not available", nil)
// 		return
// 	}

// 	// Create issue request
// 	issueRequest := models.IssueRequest{
// 		BookID:    request.BookID,
// 		Email:     email.(string),
// 		Status:    "pending", // Default status
// 		CreatedAt: time.Now(),
// 	}

// 	if err := config.DB.Create(&issueRequest).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusInternalServerError, "Could not create issue request", nil)
// 		return
// 	}

// 	utils.RespondJSON(c, http.StatusOK, "Issue request submitted successfully", issueRequest)
// }

// func ListIssueRequests(c *gin.Context) {
// 	var issueRequests []models.IssueRequest
// 	if err := config.DB.Find(&issueRequests).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusInternalServerError, "Could not fetch issue requests", nil)
// 		return
// 	}
// 	utils.RespondJSON(c, http.StatusOK, "Issue requests retrieved successfully", issueRequests)
// }

// func ApproveRejectIssueRequest(c *gin.Context) {
// 	// Get admin email from context (set in AuthMiddleware)
// 	adminEmail, exists := c.Get("email")
// 	if !exists {
// 		utils.RespondJSON(c, http.StatusUnauthorized, "Unauthorized access", nil)
// 		return
// 	}

// 	// Parse request parameters
// 	requestID := c.Param("request_id")
// 	var requestBody struct {
// 		Status string `json:"status"` // "approved" or "rejected"
// 	}
// 	if err := c.ShouldBindJSON(&requestBody); err != nil {
// 		utils.RespondJSON(c, http.StatusBadRequest, "Invalid request data", nil)
// 		return
// 	}

// 	// Validate status
// 	if requestBody.Status != "approved" && requestBody.Status != "rejected" {
// 		utils.RespondJSON(c, http.StatusBadRequest, "Invalid status value", nil)
// 		return
// 	}

// 	// Find the issue request
// 	var issueRequest models.IssueRequest
// 	if err := config.DB.Where("id = ?", requestID).First(&issueRequest).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusNotFound, "Issue request not found", nil)
// 		return
// 	}

// 	// Update the issue request status
// 	issueRequest.Status = requestBody.Status
// 	issueRequest.ApproverEmail = adminEmail.(string)
// 	issueRequest.ApprovedAt = time.Now()

// 	if err := config.DB.Save(&issueRequest).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusInternalServerError, "Could not update issue request", nil)
// 		return
// 	}

// 	// If approved, register book issue
// 	if requestBody.Status == "approved" {
// 		issueRegistry := models.IssueRegistry{
// 			BookID:     issueRequest.BookID,
// 			UserEmail:  issueRequest.Email,
// 			IssueDate:  time.Now(),
// 			ReturnDate: nil, // Not returned yet
// 		}

// 		if err := config.DB.Create(&issueRegistry).Error; err != nil {
// 			utils.RespondJSON(c, http.StatusInternalServerError, "Could not create issue registry", nil)
// 			return
// 		}
// 	}

// 	utils.RespondJSON(c, http.StatusOK, "Issue request processed successfully", issueRequest)
// }
