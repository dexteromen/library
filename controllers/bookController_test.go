package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"library/config"
	"library/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	// "library/utils"
)

func TestCreateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.ConnectDBTest()

	// Clear test DB
	config.DB.Exec("DELETE FROM users")
	config.DB.Exec("DELETE FROM book_inventories")

	r := gin.Default()

	// Middleware to mock user ID
	r.Use(func(c *gin.Context) {
		c.Set("user_id", uint(1)) // Set user ID manually
		c.Next()
	})

	r.POST("/books", CreateBook)

	// Test case 1: Invalid JSON input
	invalidJSON := []byte(`{"title": 123}`)
	req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test case 2: User not found
	validBook := models.BookInventory{
		ISBN:            "1234567890",
		Title:           "Test Book",
		Authors:         "Author Name",
		Publisher:       "Test Publisher",
		Version:         "1st Edition",
		TotalCopies:     10,
		AvailableCopies: 10,
	}
	jsonBook, _ := json.Marshal(validBook)
	req, _ = http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(jsonBook))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Test case 3: Successful book creation
	user := models.User{ID: 1, LibID: 101}
	config.DB.Create(&user)

	req, _ = http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(jsonBook))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func setupTestDB_Book() {
	config.ConnectDBTest()
	config.DB.Exec("DELETE FROM book_inventories")
	config.DB.Exec("DELETE FROM users")

	// Reset sequences
	config.DB.Exec("ALTER SEQUENCE book_inventories_id_seq RESTART WITH 1")
	config.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")

	// Insert dummy user
	user := models.User{ID: 1, LibID: 101, Name: "Test User", Email: "test@example.com", Password: "password"}
	config.DB.Create(&user)

	// Insert dummy books
	books := []models.BookInventory{
		{ISBN: "123456", Title: "Book One", Authors: "Author A", Publisher: "Pub1", Version: "1", TotalCopies: 10, AvailableCopies: 5, LibID: 101},
		{ISBN: "789101", Title: "Book Two", Authors: "Author B", Publisher: "Pub2", Version: "1", TotalCopies: 8, AvailableCopies: 2, LibID: 101},
	}
	config.DB.Create(&books)
}

func TestGetBooks(t *testing.T) {
	setupTestDB_Book()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/books", GetBooks)

	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBookByID(t *testing.T) {
	setupTestDB_Book()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/books/:id", GetBookByID)

	req, _ := http.NewRequest(http.MethodGet, "/books/123456", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateBookByID(t *testing.T) {
	setupTestDB_Book()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/books/:id", UpdateBookByID)

	updateData := map[string]interface{}{"title": "Updated Title"}
	jsonData, _ := json.Marshal(updateData)

	req, _ := http.NewRequest(http.MethodPut, "/books/123456", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteBookByID(t *testing.T) {
	setupTestDB_Book()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/books/:id", DeleteBookByID)

	req, _ := http.NewRequest(http.MethodDelete, "/books/123456", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSearchBooks(t *testing.T) {
	setupTestDB_Book()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/books/search", SearchBooks)

	req, _ := http.NewRequest(http.MethodGet, "/books/search?title=Book", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// func TestSearchBooks(t *testing.T) {
// 	setupTestDB_Book()
// 	gin.SetMode(gin.TestMode)
// 	r := gin.Default()
// 	r.GET("/books/search", SearchBooks)

// 	// Test case: Search by title
// 	t.Run("Search by Title", func(t *testing.T) {
// 		req, _ := http.NewRequest(http.MethodGet, "/books/search?title=Book", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		fmt.Println("Response:", w.Body.String()) // Debugging output
// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "Book One")
// 		assert.Contains(t, w.Body.String(), "Book Two")
// 	})

// 	// Test case: Search by Author
// 	t.Run("Search by Author", func(t *testing.T) {
// 		req, _ := http.NewRequest(http.MethodGet, "/books/search?author=Author A", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		fmt.Println("Response:", w.Body.String()) // Debugging output
// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "Book One")
// 		assert.NotContains(t, w.Body.String(), "Book Two")
// 	})

// 	// Test case: Search by Publisher
// 	t.Run("Search by Publisher", func(t *testing.T) {
// 		req, _ := http.NewRequest(http.MethodGet, "/books/search?publisher=Pub2", nil)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)

// 		fmt.Println("Response:", w.Body.String()) // Debugging output
// 		assert.Equal(t, http.StatusOK, w.Code)
// 		assert.Contains(t, w.Body.String(), "Book Two")
// 		assert.NotContains(t, w.Body.String(), "Book One")
// 	})
// }
