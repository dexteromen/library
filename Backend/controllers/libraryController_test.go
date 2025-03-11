package controllers

import (
	// "bytes"
	"bytes"
	"encoding/json"
	// "fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"library/config"
	"library/models"
	"library/utils"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	// "library/controllers"
	// "library/utils"
)

func TestCreateLibrary(t *testing.T) {
	config.ConnectDBTest()

	gin.SetMode(gin.TestMode)

	config.DB.Exec("DELETE FROM users")     // Reset users table
	config.DB.Exec("DELETE FROM libraries") // Reset libraries table

	// Create a dummy user with "reader" role
	testUser := models.User{
		Name:  "John Doe",
		Email: "john@example.com",
		Role:  "reader",
	}
	config.DB.Create(&testUser)

	// Create test router
	router := gin.Default()
	router.POST("/library", func(c *gin.Context) {
		// Mock user ID middleware
		c.Set("user_id", testUser.ID)
		CreateLibrary(c)
	})

	// Test case: Successful library creation
	t.Run("Successful Library Creation", func(t *testing.T) {
		libraryData := map[string]string{"name": "Unique Library"}
		jsonData, _ := json.Marshal(libraryData)

		req, _ := http.NewRequest("POST", "/library", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assertions
		assert.Equal(t, http.StatusCreated, recorder.Code)

		var response utils.JSONResponse
		json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, "Library created successfully", response.Message)
		assert.NotNil(t, response.Data)

		// Verify user role change
		var updatedUser models.User
		config.DB.First(&updatedUser, testUser.ID)
		assert.Equal(t, "owner", updatedUser.Role)

		// Validate response data
		ownerData, ok := response.Data.(map[string]interface{})
		assert.True(t, ok, "Response data should be a valid map")
		assert.Equal(t, testUser.Name, ownerData["Owner Of Library"])
	})

	t.Run("Invalid input", func(t *testing.T) {
		libraryData := map[string]string{"name1": "Unique Library"}
		jsonData, _ := json.Marshal(libraryData)

		req, _ := http.NewRequest("POST", "/library", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestGetLibraries(t *testing.T) {
	config.ConnectDBTest()

	gin.SetMode(gin.TestMode)

	config.DB.Exec("DELETE FROM libraries") // Reset libraries table

	// Create dummy libraries
	library1 := models.Library{Name: "Library A"}
	library2 := models.Library{Name: "Library B"}
	config.DB.Create(&library1)
	config.DB.Create(&library2)

	// Create test router
	router := gin.Default()
	router.GET("/libraries", GetLibraries)

	// Test case: Fetch all libraries
	t.Run("Fetch Libraries Successfully", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/libraries", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assertions
		assert.Equal(t, http.StatusOK, recorder.Code)

		var response utils.JSONResponse
		json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, "Libraries fetched successfully", response.Message)
		assert.NotNil(t, response.Data)

		// Check if both libraries are returned
		libraries, ok := response.Data.([]interface{})
		assert.True(t, ok, "Response data should be a list")
		assert.Len(t, libraries, 2)
	})
}

// func TestUpdateLibrary(t *testing.T) {
// 	config.ConnectDBTest()

// 	gin.SetMode(gin.TestMode)

// 	config.DB.Exec("DELETE FROM libraries") // Reset libraries table

// 	// Create a dummy library
// 	testLibrary := models.Library{Name: "Old Library"}
// 	config.DB.Create(&testLibrary)

// 	// Create test router
// 	router := gin.Default()
// 	router.PUT("/library/:id", UpdateLibrary)

// 	// Test case: Successfully update library
// 	t.Run("Successful Library Update", func(t *testing.T) {
// 		updateData := map[string]string{"name": "Updated Library"}
// 		jsonData, _ := json.Marshal(updateData)

// 		req, _ := http.NewRequest("PUT", fmt.Sprintf("/library/%d", testLibrary.ID), bytes.NewBuffer(jsonData))
// 		req.Header.Set("Content-Type", "application/json")
// 		recorder := httptest.NewRecorder()
// 		router.ServeHTTP(recorder, req)

// 		// Assertions
// 		assert.Equal(t, http.StatusOK, recorder.Code)

// 		var response utils.JSONResponse
// 		json.Unmarshal(recorder.Body.Bytes(), &response)

// 		assert.Equal(t, "Library updated successfully", response.Message)
// 		assert.NotNil(t, response.Data)

// 		// Verify updated name in database
// 		var updatedLibrary models.Library
// 		config.DB.First(&updatedLibrary, testLibrary.ID)
// 		assert.Equal(t, "Updated Library", updatedLibrary.Name)
// 	})
// }
// func TestDeleteLibrary(t *testing.T) {
// 	config.ConnectDBTest()

// 	gin.SetMode(gin.TestMode)

// 	config.DB.Exec("DELETE FROM libraries") // Reset libraries table

// 	// Create a dummy library
// 	testLibrary := models.Library{Name: "Library To Delete"}
// 	config.DB.Create(&testLibrary)

// 	// Create test router
// 	router := gin.Default()
// 	router.DELETE("/library/:id", DeleteLibrary)

// 	// Test case: Successfully delete library
// 	t.Run("Successful Library Deletion", func(t *testing.T) {
// 		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/library/%d", testLibrary.ID), nil)
// 		recorder := httptest.NewRecorder()
// 		router.ServeHTTP(recorder, req)

// 		// Assertions
// 		assert.Equal(t, http.StatusOK, recorder.Code)

// 		var response utils.JSONResponse
// 		json.Unmarshal(recorder.Body.Bytes(), &response)

// 		assert.Equal(t, "Library deleted successfully", response.Message)

// 		// Verify library is deleted
// 		var deletedLibrary models.Library
// 		err := config.DB.First(&deletedLibrary, testLibrary.ID).Error
// 		assert.Error(t, err, "Library should not exist after deletion")
// 	})
// }
