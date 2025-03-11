package utils

import (
	"net/http"

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

// JSONResponse represents the structure of a standard API response
type JSONResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

// RespondJSON sends a JSON response with a given status code
func RespondJSON(c *gin.Context, statusCode int, message string, data interface{}) {
	response := JSONResponse{
		Status:     "success",
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	if statusCode >= http.StatusBadRequest {
		response.Status = "error"
	}

	c.JSON(statusCode, response)
}
