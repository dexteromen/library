package controllers

import (
	"library/config"
	"library/models"
	"library/utils"

	// "library/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create Request
func CreateRequest(c *gin.Context) {
	var request models.RequestEvent
	if err := c.ShouldBindJSON(&request); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.RespondJSON(c, http.StatusBadRequest, "Cannot Bind JSON Data", gin.H{"error": err.Error()})
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
	request.RequestDate = time.Now()
	request.RequestType = "Borrow"

	if err := config.DB.Create(&request).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		utils.RespondJSON(c, http.StatusInternalServerError, "Cannot create request", gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusCreated, request)
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
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		utils.RespondJSON(c, http.StatusUnauthorized, "User ID not found in context", nil)
		return
	}
	approverIDUint, ok := approverID.(uint)
	if !ok {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast user ID"})
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to cast user ID", nil)
		return
	}

	now := time.Now()
	request.ApprovalDate = &now
	request.ApproverID = &approverIDUint

	if err := config.DB.Save(&request).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		utils.RespondJSON(c, http.StatusInternalServerError, "Error", gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, request)
	utils.RespondJSON(c, http.StatusOK, "All Requests", request)
}

// Get All Requests
func GetRequests(c *gin.Context) {
	var requests []models.RequestEvent
	config.DB.Find(&requests)
	// c.JSON(http.StatusOK, requests)
	utils.RespondJSON(c, http.StatusOK, "All Requests", requests)
}
