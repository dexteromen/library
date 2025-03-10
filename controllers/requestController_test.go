package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"library/config"
	"library/models"
	"library/utils"
	"net/http"
	"net/http/httptest"
	"time"

	// "strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetRequests(t *testing.T) {
	config.ConnectDBTest()
	gin.SetMode(gin.TestMode)

	config.DB.Exec("DELETE FROM request_events") // Reset table

	// Create sample requests with valid request dates
	testRequests := []models.RequestEvent{
		{ISBN: "1234567890", ReaderID: 1, RequestDate: time.Now().Format("2006-01-02 15:04:05"), RequestType: "Borrow", IssueStatus: "Pending"},
		{ISBN: "0987654321", ReaderID: 2, RequestDate: time.Now().Format("2006-01-02 15:04:05"), RequestType: "Borrow", IssueStatus: "Approved"},
	}

	for _, req := range testRequests {
		config.DB.Create(&req)
	}

	time.Sleep(100 * time.Millisecond) // Ensure DB transactions are committed

	router := gin.Default()
	router.GET("/requests", GetRequests)

	// Test case: Successfully fetch requests
	t.Run("Fetch Requests Successfully", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/requests", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assertions
		assert.Equal(t, http.StatusOK, recorder.Code)

		var response utils.JSONResponse
		json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, "All Requests", response.Message)
		assert.NotNil(t, response.Data)

		requestList, ok := response.Data.([]interface{})
		assert.True(t, ok, "Response data should be a list")
		assert.Len(t, requestList, 2) // Expecting 2 requests
	})
}

func TestCreateRequest(t *testing.T) {
	config.ConnectDBTest()

	gin.SetMode(gin.TestMode)

	config.DB.Exec("DELETE FROM request_events")
	config.DB.Exec("DELETE FROM book_inventories") // Reset book inventory

	// Create a dummy book
	testBook := []models.BookInventory{
		// {ISBN: "1234567890", AvailableCopies: 2},
		// {ISBN: "1234567000", AvailableCopies: 0},
		{
			ISBN:            "1234567890",
			Title:           "English 1",
			Authors:         "Reader",
			Publisher:       "Reader",
			Version:         "1st",
			TotalCopies:     2,
			AvailableCopies: 2,
		},
		{
			ISBN:            "1234567000",
			Title:           "English 2",
			Authors:         "Reader",
			Publisher:       "Reader",
			Version:         "1st",
			TotalCopies:     2,
			AvailableCopies: 0,
		},
		{
			ISBN:            "111-30-30-12",
			Title:           "English 3",
			Authors:         "Reader",
			Publisher:       "Reader",
			Version:         "1st",
			TotalCopies:     2,
			AvailableCopies: 1,
		},
	}
	config.DB.Create(&testBook)

	existingRequest := []models.RequestEvent{
		{
			ReqID:       2,
			ISBN:        "111-30-30-12",
			ReaderID:    1,
			RequestDate: "2025-03-07T13:35:35.258448+05:30",
			RequestType: "Borrow",
		},
	}

	config.DB.Create(&existingRequest)

	// Create test router
	router := gin.Default()
	router.POST("/request", func(c *gin.Context) {
		c.Set("user_id", uint(1)) // Mock user ID
		CreateRequest(c)
	})

	// Test case: Successfully create request
	t.Run("Successful Request Creation", func(t *testing.T) {
		requestData := map[string]string{"ISBN": "1234567890"}
		jsonData, _ := json.Marshal(requestData)

		req, _ := http.NewRequest("POST", "/request", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assertions
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})

	t.Run("Cannot Bind JSON Data", func(t *testing.T) {
		requestData := map[string]string{"a": "b"}
		jsonData, _ := json.Marshal(requestData)

		req, _ := http.NewRequest("POST", "/request", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Book not found", func(t *testing.T) {
		requestData := map[string]string{"ISBN": "1234567891"}
		jsonData, _ := json.Marshal(requestData)

		req, _ := http.NewRequest("POST", "/request", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assertions
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("Book is not available.", func(t *testing.T) {
		requestData := map[string]string{"isbn": "1234567000"}
		jsonData, _ := json.Marshal(requestData)

		req, _ := http.NewRequest("POST", "/request", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assertions
		assert.Equal(t, http.StatusConflict, recorder.Code)
	})
	t.Run("Request already exists", func(t *testing.T) {
		requestData := map[string]string{"ISBN": "111-30-30-12"}
		jsonData, _ := json.Marshal(requestData)

		req, _ := http.NewRequest("POST", "/request", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// Assertions
		assert.Equal(t, http.StatusConflict, recorder.Code)
	})
}

// func TestApproveRequest(t *testing.T) {
// 	config.ConnectDBTest()
// 	gin.SetMode(gin.TestMode)

// 	config.DB.Exec("DELETE FROM request_events") // Reset table

// 	// Create a sample request
// 	testRequest := models.RequestEvent{
// 		ISBN:        "1234567890",
// 		ReaderID:    1,
// 		RequestDate: time.Now().Format("2006-01-02 15:04:05"), // Ensure RequestDate is set properly
// 		RequestType: "Borrow",
// 		IssueStatus: "Pending",
// 	}

// 	if err := config.DB.Create(&testRequest).Error; err != nil {
// 		t.Fatalf("Failed to create request: %v", err)
// 	}

// 	// Wait for DB transaction to complete
// 	time.Sleep(100 * time.Millisecond)

// 	// Fetch the saved request
// 	var savedRequest models.RequestEvent
// 	if err := config.DB.Where("isbn = ?", "1234567890").First(&savedRequest).Error; err != nil {
// 		t.Fatalf("Failed to get request ID after creation: %v", err)
// 	}

// 	// Set up router and apply middleware
// 	router := gin.Default()
// 	router.Use(MockAuthMiddleware()) // Mock authentication middleware
// 	router.PUT("/requests/approve/:id", ApproveRequest)

// 	req, _ := http.NewRequest("PUT", "/requests/approve/"+fmt.Sprint(savedRequest.ReqID), nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer mock_token") // Mock token

// 	// Simulate request
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Assertions
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response utils.JSONResponse
// 	json.Unmarshal(w.Body.Bytes(), &response)

// 	assert.Equal(t, "All Requests", response.Message)
// 	assert.NotNil(t, response.Data)

// 	// Verify approval status in DB
// 	config.DB.First(&savedRequest, savedRequest.ReqID)
// 	assert.Equal(t, "Approved", savedRequest.IssueStatus)
// 	assert.NotNil(t, savedRequest.ApprovalDate)
// }

// MockAuthMiddleware simulates user authentication in tests
func MockAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user_id", uint(1)) // Mock user ID
		c.Next()
	}
}

func TestApproveAndIssueRequest(t *testing.T) {
	// Setup test database
	config.ConnectDBTest()
	db := config.DB

	// Create a new request in the database
	testRequest := models.RequestEvent{
		ISBN:        "1234567890",
		ReaderID:    1,
		RequestDate: time.Now().Format("2006-01-02 15:04:05"),
		RequestType: "Borrow",
		IssueStatus: "Pending",
	}

	if err := db.Create(&testRequest).Error; err != nil {
		t.Fatalf("Failed to create test request: %v", err)
	}

	// Fetch the created request to get the actual ID
	var savedRequest models.RequestEvent
	if err := db.Where("isbn = ?", "1234567890").First(&savedRequest).Error; err != nil {
		t.Fatalf("Failed to fetch created request: %v", err)
	}

	var existingBook models.BookInventory
	if err := db.Where("isbn = ?", "1234567890").First(&existingBook).Error; err == nil {
		// Book already exists, you can update it or skip the creation
		existingBook.AvailableCopies = 5
		if err := db.Save(&existingBook).Error; err != nil {
			t.Fatalf("Failed to update existing book: %v", err)
		}
	} else {
		// Book does not exist, create a new one
		// Generate a unique title and ISBN for each test
		uniqueISBN := fmt.Sprintf("1234567890-%d", time.Now().UnixNano())
		uniqueTitle := fmt.Sprintf("Test Book %d", time.Now().UnixNano())

		book := models.BookInventory{
			ISBN:            uniqueISBN,
			Title:           uniqueTitle,
			AvailableCopies: 5,
		}

		if err := db.Create(&book).Error; err != nil {
			t.Fatalf("Failed to create test book: %v", err)
		}

	}

	// Set up router and apply middleware
	router := gin.Default()
	router.Use(MockAuthMiddleware())                         // Mock authentication middleware
	router.PUT("/approve-issue/:id", ApproveAndIssueRequest) // Assuming ApproveAndIssueRequest is the handler function

	// Prepare API request to approve and issue
	url := fmt.Sprintf("/approve-issue/%d", savedRequest.ReqID)
	reqBody := map[string]interface{}{
		"approver_id": 1, // Add necessary fields if required
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Use a test recorder to capture the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validate response status
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 status")

	// Verify the request status has been updated to "Approved And Issued"
	var updatedRequest models.RequestEvent
	if err := db.First(&updatedRequest, savedRequest.ReqID).Error; err != nil {
		t.Fatalf("Failed to fetch updated request: %v", err)
	}
	assert.Equal(t, "Approved And Issued", updatedRequest.IssueStatus, "Request should be approved and issued")

	// Check if the book inventory was updated (decreased available copies)
	var updatedBook models.BookInventory
	if err := db.Where("isbn = ?", "1234567890").First(&updatedBook).Error; err != nil {
		t.Fatalf("Failed to fetch updated book inventory: %v", err)
	}
	assert.Equal(t, 4, int(updatedBook.AvailableCopies), "Available copies should be decreased by 1")

	// Check if a new issued record was created
	var issuedRecord models.IssueRegistery
	if err := db.Where("isbn = ? AND reader_id = ?", "1234567890", 1).First(&issuedRecord).Error; err != nil {
		t.Fatalf("Failed to find issued book record: %v", err)
	}
	assert.Equal(t, "Issued", issuedRecord.IssueStatus, "Book should be issued")
}
