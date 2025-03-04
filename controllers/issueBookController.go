package controllers

import (
	"library/config"
	"library/models"
	// "library/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// IssueRequest struct
type IssueRequest struct {
	ISBN               string `json:"isbn" binding:"required"`
	ReaderID           uint   `json:"reader_id" binding:"required"`
	IssueApproverID    uint   `json:"issue_approver_id" binding:"required"`
	ExpectedReturnDate string `json:"expected_return_date" binding:"required"`
}

// IssueBook handles book issuance
func IssueBook(c *gin.Context) {
	var request IssueRequest

	// Bind JSON input
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Convert expected return date
	expectedReturnDate, err := time.Parse("2006-01-02", request.ExpectedReturnDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	// Validate if book exists
	var book models.BookInventory
	if err := config.DB.Where("isbn = ?", request.ISBN).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Validate available copies
	if book.AvailableCopies <= 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "No copies available"})
		return
	}

	// Create issue record
	issue := models.IssueRegistery{
		ISBN:               request.ISBN,
		ReaderID:           request.ReaderID,
		IssueApproverID:    request.IssueApproverID,
		IssueStatus:        "Issued",
		IssueDate:          time.Now(),
		ExpectedReturnDate: expectedReturnDate,
	}

	// Save to database
	if err := config.DB.Create(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to issue book"})
		return
	}

	// Update available copies
	book.AvailableCopies--
	config.DB.Save(&book)

	c.JSON(http.StatusCreated, gin.H{"message": "Book issued successfully"})
}
