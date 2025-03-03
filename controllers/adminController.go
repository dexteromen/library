package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminIndex - Restricted route for admin users
func AdminIndex(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome, Admin!"})
}
