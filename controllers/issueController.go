package controllers

import (
	"library/config"
	"library/models"
	// "library/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create Issue
func CreateIssue(c *gin.Context) {
	var issue models.IssueRegistery
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	issue.IssueDate = time.Now()

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

	approverID := c.GetUint("user_id") // Assuming user_id is set by middleware
	now := time.Now()
	issue.ReturnDate = &now
	issue.ReturnApproverID = &approverID
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