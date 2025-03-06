package controllers

import (
	"library/config"
	"library/models"
	"library/utils"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
)

// Owner: Create Books
func CreateBook(c *gin.Context) {
	var book models.BookInventory
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	//find user by id
	var user models.User
	if err := config.DB.First(&user, c.GetUint("user_id")).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "User not found For Getting LibID", nil)
		return
	}

	//the user should be owner to create book
	//update user lib_id of the book
	book.LibID = user.LibID
	if err := config.DB.Create(&book).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create book", nil)
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Book created successfully", book)
}

// Everyone can get a book
func GetBooks(c *gin.Context) {
	var books []models.BookInventory
	config.DB.Find(&books)
	utils.RespondJSON(c, http.StatusOK, "Books retrieved successfully", books)
}

func GetBookByID(c *gin.Context) {
	var book models.BookInventory
	bookID := c.Param("id")

	if err := config.DB.First(&book, "book_id = ?", bookID).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Book retrieved successfully", book)
}

// // Admin: Update Book by ID
func UpdateBookByID(c *gin.Context) {
	var book models.BookInventory
	bookID := c.Param("id")

	// Check if the book exists
	if err := config.DB.First(&book, "book_id = ?", bookID).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	// Create a map for updating only allowed fields
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	// Prevent BookID from being updated
	delete(updateData, "book_id")

	// Update only the fields provided in JSON
	if err := config.DB.Model(&book).Updates(updateData).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to update book", nil)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Book updated successfully", book)
}

// Admin: Delete Book by ID
func DeleteBookByID(c *gin.Context) {
	var book models.BookInventory
	bookID := c.Param("id")

	// Check if the book exists
	if err := config.DB.First(&book, "book_id = ?", bookID).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	// Delete book
	if err := config.DB.Delete(&book).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to delete book", nil)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Book deleted successfully", nil)
}

func SearchBooks(c *gin.Context) {
	var books []models.BookInventory
	title := c.Query("title")
	author := c.Query("authors")
	publisher := c.Query("publisher")

	query := config.DB

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("authors LIKE ?", "%"+author+"%")
	}
	if publisher != "" {
		query = query.Where("publisher LIKE ?", "%"+publisher+"%")
	}

	query.Find(&books)

	// for i, book := range books {
	// 	if book.AvailableCopies > 0 {
	// 		books[i].Status = "Available"
	// 	} else {
	// 		books[i].Status = "Not available, expected by " + book.ExpectedReturnDate.Format("2006-01-02")
	// 	}
	// }

	utils.RespondJSON(c, http.StatusOK, "Books retrieved successfully", books)
}
