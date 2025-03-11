package controllers

import (
	"bytes"
	"encoding/json"
	"library/config"
	"library/models"
	"library/utils"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Helper function to clear the test database before each test
func clearTestDatabase() {
	config.DB.Exec("DELETE FROM users") // Adjust table name if necessary
}

func TestSignUp(t *testing.T) {
	tests := []struct {
		name         string
		input        models.User
		expectedCode int
		expectedMsg  string
	}{
		{
			name: "Successful Signup",
			input: models.User{
				Name:          "Test User",
				Email:         "testuser@example.com",
				Password:      "Test@1234",
				ContactNumber: "9876543210",
			},
			expectedCode: http.StatusCreated,
			expectedMsg:  "User created successfully",
		},
		{
			name: "Invalid Name",
			input: models.User{
				Name:          "Test User1232",
				Email:         "testuser@example.com",
				Password:      "Test@1234",
				ContactNumber: "9876543210",
			},
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "User can not be created",
		},
		{
			name: "Invalid Email Format",
			input: models.User{
				Name:          "Test User",
				Email:         "invalid-email",
				Password:      "Test@1234",
				ContactNumber: "1234567890",
			},
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "Invalid input",
		},
		{
			name: "Weak Password",
			input: models.User{
				Name:          "Test User",
				Email:         "testuser2@example.com",
				Password:      "12345",
				ContactNumber: "9876543211",
			},
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "User can not be created",
		},
		{
			name: "Missing Name",
			input: models.User{
				Email:         "testuser3@example.com",
				Password:      "Test@1234",
				ContactNumber: "9876543212",
			},
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "Invalid input",
		},
		{
			name: "Admin Role Already Exists",
			input: models.User{
				Name:          "Admin User",
				Email:         "admin@example.com",
				Password:      "Admin@1234",
				Role:          "admin",
				ContactNumber: "9876543213",
			},
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "Cannot create more than one admin",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Setup
			config.ConnectDBTest()
			clearTestDatabase() // Clear DB to prevent duplicate errors

			// Ensure only one admin exists before the admin test
			if test.input.Role == "admin" {
				config.DB.Create(&models.User{
					Name:          "Existing Admin",
					Email:         "existingadmin@example.com",
					Password:      "Admin@1234",
					Role:          "admin",
					ContactNumber: "9876543214",
				})
			}

			router := gin.Default()
			router.POST("/signup", SignUp)

			// Prepare request
			jsonValue, _ := json.Marshal(test.input)
			req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(resp, req)

			// Assertions
			assert.Equal(t, test.expectedCode, resp.Code)
			var response map[string]interface{}
			json.Unmarshal(resp.Body.Bytes(), &response)
			assert.Contains(t, response["message"], test.expectedMsg)
		})
	}
}

func setupTestUser() {
	config.ConnectDBTest()
	db := config.DB
	db.Exec("DELETE FROM users WHERE email = ?", "testuser@example.com")

	hashedPassword, _ := utils.HashPassword("Test@1234")
	testUser := models.User{
		Name:          "Test User",
		Email:         "testuser@example.com",
		Password:      hashedPassword,
		ContactNumber: "1234567890",
	}
	db.Create(&testUser)
}

func TestSignIn(t *testing.T) {
	setupTestUser()

	tests := []struct {
		name         string
		input        map[string]string
		expectedCode int
		expectedMsg  string
	}{
		{
			name: "Successful Login",
			input: map[string]string{
				"email":    "testuser@example.com",
				"password": "Test@1234",
			},
			expectedCode: http.StatusOK,
			expectedMsg:  "User logged-in successfully",
		},
		{
			name: "Invalid Email",
			input: map[string]string{
				"email":    "invalid@example.com",
				"password": "Test@1234",
			},
			expectedCode: http.StatusUnauthorized,
			expectedMsg:  "Invalid email",
		},
		{
			name: "Incorrect Password",
			input: map[string]string{
				"email":    "testuser@example.com",
				"password": "WrongPassword",
			},
			expectedCode: http.StatusUnauthorized,
			expectedMsg:  "Invalid password",
		},
		{
			name: "Invalid Input",
			input: map[string]string{
				"email2":    "testuser@example.com",
				"password1": "WrongPassword",
			},
			expectedCode: http.StatusBadRequest,
			// expectedMsg:  "Invalid Input",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := gin.Default()
			router.POST("/signin", SignIn)

			jsonValue, _ := json.Marshal(test.input)
			req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer(jsonValue))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, test.expectedCode, resp.Code)
			// var response map[string]interface{}
			// json.Unmarshal(resp.Body.Bytes(), &response)
			// assert.Contains(t, response["message"], test.expectedMsg)
		})
	}
}

func TestSignOut(t *testing.T) {
	setupTestUser()

	// First, sign in to get a valid token
	router := gin.Default()
	router.POST("/signin", SignIn)
	signInData := map[string]string{
		"email":    "testuser@example.com",
		"password": "Test@1234",
	}
	jsonValue, _ := json.Marshal(signInData)
	req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var signInResponse map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &signInResponse)

	data, ok := signInResponse["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("Failed to extract data field from sign-in response: %v", signInResponse)
	}
	token, ok := data["token"].(string)
	if !ok || token == "" {
		t.Fatalf("Failed to extract token from sign-in response: %v", signInResponse)
	}

	// Test Logout
	router.POST("/signout", SignOut)
	req, _ = http.NewRequest("POST", "/signout", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, "User logged out successfully !!", response["message"])
}

func setupTestUsers() {
	config.ConnectDBTest()
	config.DB.Exec("DROP TABLE IF EXISTS users")
	config.DB.AutoMigrate(&models.User{}) // Ensure table exists

	db := config.DB
	// db.Exec("DELETE FROM users")
	// db.Exec("DROP TABLE users")

	hashedPassword, _ := utils.HashPassword("Test@1234")
	testUser := models.User{
		Name:          "Test User",
		Email:         "testuser@example.com",
		Password:      hashedPassword,
		ContactNumber: "1234567890",
	}
	db.Create(&testUser)
}

func TestGetUsers(t *testing.T) {
	// clearTestDatabase()
	setupTestUsers()

	router := gin.Default()
	router.GET("/users", GetUsers)

	req, _ := http.NewRequest("GET", "/users", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Contains(t, response["message"], "All users retrived !!")
}

func TestGetUserById(t *testing.T) {
	setupTestUsers()

	router := gin.Default()
	router.GET("/users/:id", GetUserById)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Contains(t, response["message"], "User retrived !!")
}

// func TestUpdateUserById(t *testing.T) {
// 	setupTestUsers()

// 	router := gin.Default()
// 	router.PUT("/users/:id", UpdateUserById)

// 	updatedUser := models.User{
// 		Name:          "Updated Name",
// 		Email:         "updated@example.com",
// 		Password:      "Updated@1234",
// 		ContactNumber: "1234567890",
// 	}

// 	jsonValue, _ := json.Marshal(updatedUser)
// 	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp := httptest.NewRecorder()
// 	router.ServeHTTP(resp, req)

// 	assert.Equal(t, http.StatusOK, resp.Code)
// 	var response map[string]interface{}
// 	json.Unmarshal(resp.Body.Bytes(), &response)
// 	assert.Contains(t, response["message"], "User updated successfully")
// }

func TestDeleteUserById(t *testing.T) {
	setupTestUsers()

	router := gin.Default()
	router.DELETE("/users/:id", DeleteUserById)

	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Contains(t, response["message"], "User deleted successfully")
}
