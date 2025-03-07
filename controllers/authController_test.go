package controllers

import (
	"bytes"
	"encoding/json"
	"library/config"
	"library/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	// Setup
	config.LoadEnvVariablesTest()
	config.ConnectDBTest()
	router := gin.Default()
	router.POST("/signup", SignUp)

	// Test data
	user := models.User{
		Name:          "Test User",
		Email:         "testuser@example.com",
		Password:      "Test@1234",
		ContactNumber: "1234567890",
	}
	jsonValue, _ := json.Marshal(user)

	// Create a request
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, "User created successfully", response["message"])
}

func TestSignIn(t *testing.T) {
	// Setup
	config.LoadEnvVariablesTest()
	config.ConnectDBTest()
	router := gin.Default()
	router.POST("/signin", SignIn)

	// Test data
	credentials := SignInCredentials{
		Email:    "testuser@example.com",
		Password: "Test@1234",
	}
	jsonValue, _ := json.Marshal(credentials)

	// Create a request
	req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(resp, req)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, "User logged-in successfully !!", response["message"])
}
