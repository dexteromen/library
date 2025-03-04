package utils

import (
	"github.com/gin-gonic/gin"
)

// APIResponse formats the JSON response for API endpoints
func APIResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"Status":  status,
		"Message": message,
		"Data":    data,
	})
}
