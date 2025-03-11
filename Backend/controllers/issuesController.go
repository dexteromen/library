package controllers

import (
	"library/config"
	"library/models"

	"library/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetIssues godoc
// @Summary Get all issues
// @Description Retrieve a list of all issues
// @Tags issues
// @Accept  json
// @Produce  json
// @Success 200   {object}  map[string]interface{}  "All issues retrieved successfully"
// @Failure 500   {object}  map[string]interface{}  "Failed to fetch issues"
// @Router /issues [get]
func GetIssues(c *gin.Context) {
	var issues []models.IssueRegistery
	config.DB.Find(&issues)
	utils.RespondJSON(c, http.StatusOK, "All Issues Retrieved.", issues)
}

func IssueBook(c *gin.Context) {
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
	IssueApproverID, exists := c.Get("user_id")
	if !exists {
		utils.RespondJSON(c, http.StatusUnauthorized, "User ID not found in context", nil)
		return
	}

	//Update RequestEvent
	// request.ApprovalDate = time.Now()
	// request.ApproverID = &IssueApproverIDUint
	request.IssueStatus = "Issued"
	if err := config.DB.Save(&request).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
		return
	}

	// Insert new record in IssueRegistry
	issue := models.IssueRegistery{
		ISBN:               request.ISBN,
		ReaderID:           request.ReaderID,
		IssueApproverID:    IssueApproverID.(uint),
		IssueStatus:        "Issued",
		IssueDate:          time.Now().Format("2006-01-02 15:04:05"),          // in format "2006-01-02 15:04:05"
		ExpectedReturnDate: time.Now().AddDate(0, 0, 14).Format("2006-01-02"), // Default 2-week return period
	}

	if err := config.DB.Create(&issue).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to issue book", nil)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Book issued successfully", issue)
}

// Return Book by isbn and reader_id from isuue_registery

// ReturnBook godoc
// @Summary Return a book
// @Description Return a book by ISBN and reader ID from issue registry
// @Tags issues
// @Accept  json
// @Produce  json
// @Param   id     path      string  true  "Book ISBN"
// @Success 200   {object}  map[string]interface{}  "Book returned successfully"
// @Failure 404   {object}  map[string]interface{}  "Issue record not found"
// @Failure 409   {object}  map[string]interface{}  "Book already returned"
// @Failure 401   {object}  map[string]interface{}  "User ID not found in context"
// @Failure 500   {object}  map[string]interface{}  "Failed to update issue record or book inventory"
// @Router /return/{id} [put]
func ReturnBook(c *gin.Context) {
	id := c.Param("id")

	var issue models.IssueRegistery
	if err := config.DB.Where("isbn = ?", id).First(&issue).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Issue record not found", nil)
		return
	}

	// Check if the book is already returned
	if issue.IssueStatus == "Returned" {
		utils.RespondJSON(c, http.StatusConflict, "Book already returned", nil)
		return
	}

	// Check if readerid is same as the logged in user
	readerID, exists := c.Get("user_id")
	if !exists {
		utils.RespondJSON(c, http.StatusUnauthorized, "User ID not found in context", nil)
		return
	}
	readerIDUint, ok := readerID.(uint)
	if !ok {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to cast user ID", nil)
		return
	}
	if issue.ReaderID != readerIDUint {
		utils.RespondJSON(c, http.StatusUnauthorized, "You are not authorized to return this book", nil)
		return
	}

	// Update issue record with return date
	now := time.Now()
	issue.ReturnDate = now.Format("2006-01-02 15:04:05")
	issue.IssueStatus = "Returned"

	if err := config.DB.Save(&issue).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to update issue record", nil)
		return
	}

	// Update book inventory (increase available copies)
	var book models.BookInventory
	if err := config.DB.Where("isbn = ?", issue.ISBN).First(&book).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Book inventory record not found", nil)
		return
	}

	book.AvailableCopies++
	if err := config.DB.Save(&book).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to update book inventory", nil)
		return
	}

	//update the requestType in requestEvent
	var request models.RequestEvent
	if err := config.DB.Where("isbn = ?", issue.ISBN).First(&request).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Request record not found", nil)
		return
	}
	request.RequestType = "Return"
	if err := config.DB.Save(&request).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to update request record", nil)
		return
	}

	// utils.RespondJSON(c, http.StatusOK, "Book returned successfully", gin.H{"issue": issue, "book": book, "request": request})
	utils.RespondJSON(c, http.StatusOK, "Book returned successfully", nil)
}
