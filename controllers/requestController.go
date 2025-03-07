package controllers

import (
	"library/config"
	"library/models"
	"library/utils"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Get All Requests
func GetRequests(c *gin.Context) {
	var requests []models.RequestEvent
	config.DB.Find(&requests)

	//if requwsts is empty
	if len(requests) == 0 {
		utils.RespondJSON(c, http.StatusNotFound, "No Requests Found", nil)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "All Requests", requests)
}

// Create Request
func CreateRequest(c *gin.Context) {
	var request models.RequestEvent
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Cannot Bind JSON Data", gin.H{"error": err.Error()})
		return
	}

	//Finding Book by ISBN
	var book models.BookInventory
	if err := config.DB.Where("isbn = ?", request.ISBN).First(&book).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	//Checking if book is available
	if book.AvailableCopies <= 0 {
		utils.RespondJSON(c, http.StatusConflict, "Request Cannot be made !!, Book is not available.", nil)
		return
	}

	// Because reader is creating the request
	// Assuming middleware sets user_id in context
	readerID, exists := c.Get("user_id")
	if !exists {
		utils.RespondJSON(c, http.StatusUnauthorized, "Reader ID not found in context", nil)
		return
	}
	readerIDUint, ok := readerID.(uint)
	if !ok {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to cast Reader Id", nil)
		return
	}

	request.ReaderID = readerIDUint

	//checking if user has already requested for the book
	var existingRequest models.RequestEvent
	if err := config.DB.Where("isbn = ? AND reader_id = ? AND approval_date IS NULL", request.ISBN, request.ReaderID).First(&existingRequest).Error; err == nil {
		utils.RespondJSON(c, http.StatusConflict, "Request already exists", nil)
		return
	}

	request.ReaderID = readerIDUint
	request.RequestDate = time.Now().Format("2006-01-02 15:04:05")
	request.RequestType = "Borrow"
	request.IssueStatus = "Pending"

	if err := config.DB.Create(&request).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Cannot create request", gin.H{"error": err.Error()})
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Request Created", request)
}

// Approve Request
func ApproveRequest(c *gin.Context) {
	var request models.RequestEvent
	id := c.Param("id")

	if err := config.DB.First(&request, id).Error; err != nil {
		// c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		utils.RespondJSON(c, http.StatusNotFound, "Request not found", nil)
		return
	}

	// // Assuming middleware sets user_id in context
	// approverID := c.GetUint("user_id")
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
	request.IssueStatus = "Approved"

	if err := config.DB.Save(&request).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
		return
	}

	utils.RespondJSON(c, http.StatusOK, "All Requests", request)
}

// Approve And Issue Request
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
