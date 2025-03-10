package controllers

import (
	// "bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"library/config"
	"library/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	// "library/controllers"
	// "library/utils"
)

func setupTestDB() {
	config.ConnectDBTest()
	config.DB.Exec("DELETE FROM users")
	config.DB.Exec("DELETE FROM book_inventories")
	config.DB.Exec("DELETE FROM issue_registeries")
	config.DB.Exec("DELETE FROM request_events")

	// Insert dummy user
	user := models.User{ID: 1, Name: "Test User", Email: "test@example.com", ContactNumber: "1234567890", Password: "password", Role: "admin", LibID: 1}
	config.DB.Create(&user)

	// Insert dummy book
	book := models.BookInventory{ISBN: "123456", Title: "Test Book", Authors: "Author", Publisher: "Publisher", Version: "1", TotalCopies: 5, AvailableCopies: 5, LibID: 1}
	config.DB.Create(&book)
}

func TestGetIssues(t *testing.T) {
	setupTestDB()
	router := gin.Default()
	router.GET("/issues", GetIssues)

	// Insert dummy issue with valid dates
	issue := models.IssueRegistery{
		IssueID:            1,
		ISBN:               "123456",
		ReaderID:           1,
		IssueApproverID:    1,
		IssueStatus:        "Issued",
		IssueDate:          time.Now().Format("2006-01-02 15:04:05"),
		ExpectedReturnDate: time.Now().AddDate(0, 0, 14).Format("2006-01-02"),
	}
	config.DB.Create(&issue)

	req, _ := http.NewRequest("GET", "/issues", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response struct {
		Message string                  `json:"message"`
		Data    []models.IssueRegistery `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "All Issues Retrieved.", response.Message)
	assert.NotEmpty(t, response.Data)
}

func TestIssueBook(t *testing.T) {
	setupTestDB()
	router := gin.Default()

	// Middleware to set user_id in context
	router.Use(func(c *gin.Context) {
		c.Set("user_id", uint(1)) // Setting a valid user ID
		c.Next()
	})

	router.POST("/issue/:id", IssueBook)

	// Insert dummy request
	request := models.RequestEvent{
		ReqID:       1,
		ISBN:        "123456",
		ReaderID:    1,
		RequestType: "Borrow",
		RequestDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	config.DB.Create(&request)

	req, _ := http.NewRequest("POST", "/issue/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code) // Expecting 200 OK
}

func addAuthHeader(req *http.Request) {
	req.Header.Set("Authorization", "Bearer test_token") // Mock authentication token
}

func TestReturnBook(t *testing.T) {
	setupTestDB()
	router := gin.Default()

	// Mock authentication middleware
	router.Use(func(c *gin.Context) {
		c.Set("user_id", uint(1)) // Mock user ID
		c.Next()
	})

	router.POST("/return/:id", ReturnBook)

	// Ensure all tables are cleared before test
	config.DB.Exec("TRUNCATE users, book_inventories, issue_registeries, request_events RESTART IDENTITY CASCADE")

	// Insert dummy user
	user := models.User{
		Name:          "John Doe",
		Email:         "john@example.com",
		ContactNumber: "1234567890",
		Password:      "password",
		Role:          "reader",
		LibID:         1,
	}
	config.DB.Create(&user)

	// Insert dummy book
	book := models.BookInventory{
		ISBN:            "123456",
		Title:           "Test Book",
		Authors:         "Test Author",
		Publisher:       "Test Publisher",
		Version:         "1",
		TotalCopies:     5,
		AvailableCopies: 3, // Start with 3 available copies
		LibID:           1,
	}
	config.DB.Create(&book)

	// Insert an issue record
	issue := models.IssueRegistery{
		ISBN:               "123456",
		ReaderID:           1, // Must match user_id from middleware
		IssueApproverID:    1,
		IssueStatus:        "Issued",
		IssueDate:          time.Now().Format("2006-01-02 15:04:05"),
		ExpectedReturnDate: time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
	}
	config.DB.Create(&issue)

	// Insert a request event
	request := models.RequestEvent{
		ISBN:        "123456",
		ReaderID:    1,
		RequestDate: time.Now().Format("2006-01-02 15:04:05"), // Ensure valid timestamp
		RequestType: "Borrow",
		IssueStatus: "Issued",
	}
	config.DB.Create(&request)

	// Send the return request
	req, _ := http.NewRequest("POST", "/return/123456", nil)
	addAuthHeader(req)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Fetch updated book details
	var updatedBook models.BookInventory
	config.DB.Where("isbn = ?", "123456").First(&updatedBook)

	// Fetch updated request event
	var updatedRequest models.RequestEvent
	config.DB.Where("isbn = ?", "123456").First(&updatedRequest)

	// Validate the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 4, updatedBook.AvailableCopies)       // Ensure copies increased by 1
	assert.Equal(t, "Return", updatedRequest.RequestType) // Ensure request type updated
}
