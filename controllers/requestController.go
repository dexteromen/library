package controllers

import (
	"library/config"
	"library/models"

	// "library/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create Request
func CreateRequest(c *gin.Context) {
	var request models.RequestEvent
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.RequestDate = time.Now()

	if err := config.DB.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, request)
}

// Approve Request
func ApproveRequest(c *gin.Context) {
	var request models.RequestEvent
	id := c.Param("id")

	if err := config.DB.First(&request, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
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
	request.ApprovalDate = &now
	request.ApproverID = &approverIDUint

	if err := config.DB.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

// Get All Requests
func GetRequests(c *gin.Context) {
	var requests []models.RequestEvent
	config.DB.Find(&requests)
	c.JSON(http.StatusOK, requests)
}
