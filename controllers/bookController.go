package controllers

import (
	"library/config"
	"library/models"
	"library/utils"
	"net/http"
	// "time"

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