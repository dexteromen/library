package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserIndex - Restricted route for regular users
func UserIndex(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome, User!"})
}
