package controllers

import (
	"library/config"
	"library/models"
	"library/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Admin: Manage Books
func CreateBook(c *gin.Context) {
	var book models.BookInventory
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to create book", nil)
		return
	}

	utils.APIResponse(c, http.StatusCreated, "Book created successfully", book)
}

// Everyone can get a book
func GetBooks(c *gin.Context) {
	var books []models.BookInventory
	config.DB.Find(&books)
	utils.APIResponse(c, http.StatusOK, "Books retrieved successfully", books)
}

// // Admin: Update Book by ID
func UpdateBookByID(c *gin.Context) {
	var book models.BookInventory
	bookID := c.Param("id")

	// Check if the book exists
	if err := config.DB.First(&book, "book_id = ?", bookID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	// Create a map for updating only allowed fields
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	// Prevent BookID from being updated
	delete(updateData, "book_id")

	// Update only the fields provided in JSON
	if err := config.DB.Model(&book).Updates(updateData).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to update book", nil)
		return
	}

	utils.APIResponse(c, http.StatusOK, "Book updated successfully", book)
}

// Admin: Delete Book by ID
func DeleteBookByID(c *gin.Context) {
	var book models.BookInventory
	bookID := c.Param("id")

	// Check if the book exists
	if err := config.DB.First(&book,"book_id = ?", bookID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	// Delete book
	if err := config.DB.Delete(&book).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to delete book", nil)
		return
	}

	utils.APIResponse(c, http.StatusOK, "Book deleted successfully", nil)
}

// User: Request a Book
func RequestBook(c *gin.Context) {
	var request models.RequestEvents
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	request.RequestDate = time.Now()

	if err := config.DB.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to request book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book request submitted"})
}

// Approver: Approve or Reject Requests
func ApproveRequestById(c *gin.Context) {
	var request models.RequestEvents
	if err := config.DB.First(&request, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	request.ApprovalDate = time.Now()
	request.ApproverID = uint(c.GetInt("userID"))

	if err := config.DB.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request approved"})
}


// func UpdateBookByID(c *gin.Context) {
// 	var book models.BookInventory
// 	bookID := c.Param("id")

// 	// Check if the book exists
// 	if err := config.DB.First(&book, "book_id = ?", bookID).Error; err != nil {
// 		utils.APIResponse(c, http.StatusNotFound, "Book not found", nil)
// 		return
// 	}

// 	// Bind JSON input to book struct
// 	if err := c.ShouldBindJSON(&book); err != nil {
// 		utils.APIResponse(c, http.StatusBadRequest, "Invalid input", nil)
// 		return
// 	}

// 	// Update book
// 	if err := config.DB.Save(&book).Error; err != nil {
// 		utils.APIResponse(c, http.StatusInternalServerError, "Failed to update book", nil)
// 		return
// 	}

// 	utils.APIResponse(c, http.StatusOK, "Book updated successfully", book)
// }