package middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoleMiddleware(t *testing.T) {
	// Setup
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("role", "admin")
	})
	router.Use(RoleMiddleware("admin"))
	router.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Admin access granted"})
	})

	// Create a request
	req, _ := http.NewRequest("GET", "/admin", nil)
	resp := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, "Admin access granted", response["message"])
}
