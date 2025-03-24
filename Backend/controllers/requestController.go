package controllers

import (
	"library/config"
	"library/models"
	"library/utils"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetRequests godoc
// @Summary Get all requests
// @Description Retrieve a list of all requests
// @Tags requests
// @Accept  json
// @Produce  json
// @Success 200   {object}  map[string]interface{}  "All requests retrieved successfully"
// @Failure 500   {object}  map[string]interface{}  "Failed to fetch requests"
// @Router /requests [get]
func GetRequests(c *gin.Context) {
	var requests []models.RequestEvent
	config.DB.Find(&requests)

	utils.RespondJSON(c, http.StatusOK, "All Requests", requests)
}

// CreateRequest godoc
// @Summary Create a new request
// @Description Create a new request for a book
// @Tags requests
// @Accept  json
// @Produce  json
// @Param   request  body      models.RequestEvent  true  "Request data"
// @Success 201   {object}  map[string]interface{}  "Request created successfully"
// @Failure 400   {object}  map[string]interface{}  "Cannot Bind JSON Data"
// @Failure 404   {object}  map[string]interface{}  "Book not found"
// @Failure 409   {object}  map[string]interface{}  "Request already exists or Book is not available"
// @Failure 500   {object}  map[string]interface{}  "Cannot create request"
// @Router /request [post]
func CreateRequest(c *gin.Context) {
	var request models.RequestEvent
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Cannot Bind JSON Data", gin.H{"error": err.Error()})
		return
	}

	// Finding the book by ISBN
	var book models.BookInventory
	if err := config.DB.Where("isbn = ?", request.ISBN).First(&book).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	// Checking if the book is available
	if book.AvailableCopies <= 0 {
		utils.RespondJSON(c, http.StatusConflict, "Request cannot be made! Book is not available.", nil)
		return
	}

	// Assuming middleware sets user_id in context
	readerID, _ := c.Get("user_id")
	readerIDUint, _ := readerID.(uint)

	// Checking if the user has already borrowed the book and hasn't returned it
	var borrowedRequest models.RequestEvent
	if err := config.DB.Where("isbn = ? AND reader_id = ? AND request_type = ? AND issue_status IN (?, ?)",
		request.ISBN, readerIDUint, "Borrow", "Approved", "Issued").First(&borrowedRequest).Error; err == nil {
		utils.RespondJSON(c, http.StatusConflict, "You must return the book before making another request for it.", nil)
		return
	}

	// Checking if there is an existing pending request
	var pendingRequest models.RequestEvent
	if err := config.DB.Where("isbn = ? AND reader_id = ? AND issue_status = ?", request.ISBN, readerIDUint, "Pending").First(&pendingRequest).Error; err == nil {
		utils.RespondJSON(c, http.StatusConflict, "You already have a pending request for this book.", nil)
		return
	}

	// Checking if the user has already requested the book and approval_date is NULL
	var existingRequest models.RequestEvent
	if err := config.DB.Where("isbn = ? AND reader_id = ? AND approval_date IS NULL", request.ISBN, readerIDUint).First(&existingRequest).Error; err == nil {
		utils.RespondJSON(c, http.StatusConflict, "Request already exists", nil)
		return
	}

	// Creating new request
	request.ReaderID = readerIDUint
	request.RequestDate = time.Now().Format("2006-01-02 15:04:05")
	request.RequestType = "Borrow"
	request.IssueStatus = "Pending"

	if err := config.DB.Create(&request).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create request", gin.H{"error": err.Error()})
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Request Created", request)
}

// func CreateRequest(c *gin.Context) {
// 	var request models.RequestEvent
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		utils.RespondJSON(c, http.StatusBadRequest, "Cannot Bind JSON Data", gin.H{"error": err.Error()})
// 		return
// 	}

// 	//Finding Book by ISBN
// 	var book models.BookInventory
// 	if err := config.DB.Where("isbn = ?", request.ISBN).First(&book).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusNotFound, "Book not found", nil)
// 		return
// 	}

// 	//Checking if book is available
// 	if book.AvailableCopies <= 0 {
// 		utils.RespondJSON(c, http.StatusConflict, "Request Cannot be made !!, Book is not available.", nil)
// 		return
// 	}

// 	// Because reader is creating the request
// 	// Assuming middleware sets user_id in context
// 	readerID, _ := c.Get("user_id")
// 	readerIDUint, _ := readerID.(uint)

// 	request.ReaderID = readerIDUint

// 	//checking if user has already requested for the book
// 	var existingRequest models.RequestEvent
// 	if err := config.DB.Where("isbn = ? AND reader_id = ? AND approval_date IS NULL", request.ISBN, request.ReaderID).First(&existingRequest).Error; err == nil {
// 		utils.RespondJSON(c, http.StatusConflict, "Request already exists", nil)
// 		return
// 	}

// 	request.ReaderID = readerIDUint
// 	request.RequestDate = time.Now().Format("2006-01-02 15:04:05")
// 	request.RequestType = "Borrow"
// 	request.IssueStatus = "Pending"

// 	config.DB.Create(&request)
// 	utils.RespondJSON(c, http.StatusCreated, "Request Created", request)
// }

/*
// Approve Request
// func ApproveRequest(c *gin.Context) {
// 	var request models.RequestEvent
// 	id := c.Param("id")

// 	if err := config.DB.First(&request, id).Error; err != nil {
// 		// c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
// 		utils.RespondJSON(c, http.StatusNotFound, "Request not found", nil)
// 		return
// 	}

// 	// // Assuming middleware sets user_id in context
// 	// approverID := c.GetUint("user_id")
// 	approverID, exists := c.Get("user_id")
// 	if !exists {
// 		utils.RespondJSON(c, http.StatusUnauthorized, "User ID not found in context", nil)
// 		return
// 	}
// 	approverIDUint, ok := approverID.(uint)
// 	if !ok {
// 		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to cast user ID", nil)
// 		return
// 	}

// 	now := time.Now()
// 	request.ApprovalDate = &now
// 	request.ApproverID = &approverIDUint
// 	request.IssueStatus = "Approved"

// 	if err := config.DB.Save(&request).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
// 		return
// 	}

// 	utils.RespondJSON(c, http.StatusOK, "All Requests", request)
// }
*/

// ApproveAndIssueRequest godoc
// @Summary Approve and issue a book request
// @Description Approve a book request and issue the book to the user
// @Tags requests
// @Accept  json
// @Produce  json
// @Param   id     path      string  true  "Request ID"
// @Success 200   {object}  map[string]interface{}  "Book issued and approved successfully"
// @Failure 404   {object}  map[string]interface{}  "Request not found"
// @Failure 409   {object}  map[string]interface{}  "No available copies for this book"
// @Failure 500   {object}  map[string]interface{}  "Failed to update book inventory or issue book"
// @Router /approve-issue/{id} [put]
func ApproveAndIssueRequest(c *gin.Context) {
	var request models.RequestEvent
	id := c.Param("id")
	if err := config.DB.First(&request, id).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Request not found", nil)
		return
	}

	// Fetch book inventory
	var book models.BookInventory
	if err := config.DB.Where("isbn = ?", request.ISBN).First(&book).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}
	// Check if the book is available
	if book.AvailableCopies <= 0 {
		utils.RespondJSON(c, http.StatusConflict, "No available copies for this book", nil)
		return
	}
	// Update book inventory (decrease available copies)
	book.AvailableCopies--
	if err := config.DB.Save(&book).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to update book inventory", nil)
		return
	}

	// // Assuming middleware sets user_id in context
	approverID, exists := c.Get("user_id")
	if !exists {
		utils.RespondJSON(c, http.StatusUnauthorized, "User ID not found in context", nil)
		return
	}
	approverIDUint, ok := approverID.(uint)
	if !ok {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to cast user ID", nil)
		return
	}

	now := time.Now()
	request.ApprovalDate = &now
	request.ApproverID = &approverIDUint
	request.IssueStatus = "Approved And Issued"

	if err := config.DB.Save(&request).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
		return
	}

	// utils.RespondJSON(c, http.StatusOK, "All Requests", request)

	// Insert new record in IssueRegistry
	issue := models.IssueRegistery{
		ISBN:               request.ISBN,
		ReaderID:           request.ReaderID,
		IssueApproverID:    *request.ApproverID,
		IssueStatus:        "Issued",
		IssueDate:          time.Now().Format("2006-01-02 15:04:05"),          // in format "2006-01-02 15:04:05"
		ExpectedReturnDate: time.Now().AddDate(0, 0, 14).Format("2006-01-02"), // Default 2-week return period
	}
	if err := config.DB.Create(&issue).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to issue book", nil)
		return
	}
	utils.RespondJSON(c, http.StatusOK, "Book issued and approved successfully", nil)
}
