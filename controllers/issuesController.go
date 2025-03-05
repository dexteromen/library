package controllers

import (
	"library/config"
	"library/models"

	// "library/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// // Create Issue
// func CreateIssue(c *gin.Context) {
// 	var issue models.IssueRegistery
// 	if err := c.ShouldBindJSON(&issue); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	issue.IssueDate = time.Now()

// 	if err := config.DB.Create(&issue).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, issue)
// }

func CreateIssue(c *gin.Context) {
	var issue models.IssueRegistery
	
	// Bind JSON directly to the struct
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse Expected Return Date from JSON
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, issue.ExpectedReturnDate) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}
	issue.ExpectedReturnDate = parsedDate.Format(layout) // Assign parsed date

	issue.IssueDate = time.Now() // Automatically set issue date

	// Save to DB
	if err := config.DB.Create(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, issue)
}


// Return Book
func ReturnBook(c *gin.Context) {
	var issue models.IssueRegistery
	id := c.Param("id")

	if err := config.DB.First(&issue, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		return
	}

	// // Assuming middleware sets user_id in context
	// approverID := c.GetUint("user_id")
	approverID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	approverIDUint, ok := approverID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast user ID"})
		return
	}

	now := time.Now()
	issue.ReturnDate = &now
	issue.ReturnApproverID = &approverIDUint
	issue.IssueStatus = "Returned"

	if err := config.DB.Save(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, issue)
}

// Get All Issues
func GetIssues(c *gin.Context) {
	var issues []models.IssueRegistery
	config.DB.Find(&issues)
	c.JSON(http.StatusOK, issues)
}
