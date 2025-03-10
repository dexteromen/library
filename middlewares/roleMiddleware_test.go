package middlewares

import (
	// "encoding/json"
	"library/config"
	"library/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// func TestRoleMiddleware(t *testing.T) {
// 	// Setup
// 	router := gin.Default()
// 	router.Use(func(c *gin.Context) {
// 		c.Set("role", "admin")
// 	})
// 	router.Use(RoleMiddleware("admin"))
// 	router.GET("/admin", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "Admin access granted"})
// 	})

// 	// Create a request
// 	req, _ := http.NewRequest("GET", "/admin", nil)
// 	resp := httptest.NewRecorder()

// 	// Perform the request
// 	router.ServeHTTP(resp, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusOK, resp.Code)
// 	var response map[string]interface{}
// 	json.Unmarshal(resp.Body.Bytes(), &response)
// 	assert.Equal(t, "Admin access granted", response["message"])
// }

func setupTestDB_Role() {
	config.ConnectDBTest()
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

func TestRoleMiddleware(t *testing.T) {
	setupTestDB_Role()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		// Simulating user role retrieval from request (this would normally be extracted from JWT, session, etc.)
		userRole := c.Query("role")
		if userRole == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Missing role"})
			c.Abort()
			return
		}
		c.Set("role", userRole)
		c.Next()
	})

	r.GET("/admin", RoleMiddleware("admin"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome, Admin!"})
	})

	r.GET("/reader", RoleMiddleware("reader"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome, Reader!"})
	})

	r.GET("/owner", RoleMiddleware("owner"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome, Owner!"})
	})

	// Test case: Admin access
	t.Run("Admin access allowed", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/admin?role=admin", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Welcome, Admin!")
	})

	// Test case: Reader denied admin access
	t.Run("Reader access denied to admin", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/admin?role=reader", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.Contains(t, w.Body.String(), "Forbidden access") // Updated message
	})

	// Test case: Owner access allowed
	t.Run("Owner access allowed", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/owner?role=owner", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Welcome, Owner!")
	})

	// Test case: Unauthorized access (no role)
	t.Run("Unauthorized access", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/admin", nil) // No role provided
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.Contains(t, w.Body.String(), "Missing role")
	})
}
