package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"library/config"
	"library/models"
	"library/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestDB_Auth() {
	config.ConnectDBTest()
	config.DB.Exec("DELETE FROM sessions")
	config.DB.Exec("ALTER SEQUENCE sessions_id_seq RESTART WITH 1")
	config.DB.Exec("DELETE FROM users")
	config.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")

	// Insert dummy users
	users := []models.User{
		{ID: 1, LibID: 101, Name: "Admin User", Email: "admin@example.com", ContactNumber: "1234567890", Password: "password", Role: "admin"},
		{ID: 2, LibID: 102, Name: "Reader User", Email: "reader@example.com", ContactNumber: "0987654321", Password: "password", Role: "reader"},
		{ID: 3, LibID: 103, Name: "Owner User", Email: "owner@example.com", ContactNumber: "1122334455", Password: "password", Role: "owner"},
	}
	config.DB.Create(&users)
}

func TestAuthMiddleware(t *testing.T) {
	setupTestDB_Auth()

	// Create a valid token
	token, _ := utils.GenerateJWT(1, "admin@example.com", "admin")

	// Create a dummy session with the same token
	session := models.Session{
		Token:    token,
		IsActive: true,
		UserID:   1,
	}
	config.DB.Create(&session)

	// Set up Gin router with AuthMiddleware
	router := gin.Default()
	router.Use(AuthMiddleware())
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})

	// Test case: Valid token
	t.Run("Valid Token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Access granted")
	})

	// Test case: Invalid token
	t.Run("Invalid Token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer invalid_token")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Invalid token")
	})

	// Test case: Missing token
	t.Run("Missing Token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/protected", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "Authorization token is required")
	})

	config.DB.Exec("DELETE FROM sessions")
	config.DB.Exec("ALTER SEQUENCE sessions_id_seq RESTART WITH 1")
	config.DB.Exec("DELETE FROM users")
	config.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}
