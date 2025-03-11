package controllers

import (
	"library/config"
	"library/models"
	"library/utils"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book in the library
// @Tags books
// @Accept  json
// @Produce  json
// @Param   book  body      models.BookInventory  true  "Book data"
// @Success 201   {object}  map[string]interface{}  "Book created successfully"
// @Failure 400   {object}  map[string]interface{}  "Invalid input"
// @Failure 404   {object}  map[string]interface{}  "User not found"
// @Failure 500   {object}  map[string]interface{}  "Failed to create book"
// @Router /book [post]
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

// GetBooks godoc
// @Summary Get all books
// @Description Retrieve a list of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200   {object}  map[string]interface{}  "Books retrieved successfully"
// @Failure 500   {object}  map[string]interface{}  "Failed to fetch books"
// @Router /books [get]
func GetBooks(c *gin.Context) {
	var books []models.BookInventory
	config.DB.Find(&books)
	utils.RespondJSON(c, http.StatusOK, "Books retrieved successfully", books)
}

// GetBookByID godoc
// @Summary Get a book by ID
// @Description Retrieve a book by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param   id     path      string  true  "Book ID"
// @Success 200   {object}  map[string]interface{}  "Book retrieved successfully"
// @Failure 404   {object}  map[string]interface{}  "Book not found"
// @Router /book/{id} [get]
func GetBookByID(c *gin.Context) {
	var book models.BookInventory
	bookID := c.Param("id")

	if err := config.DB.First(&book, "isbn = ?", bookID).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Book retrieved successfully", book)
}

// UpdateBookByID godoc
// @Summary Update a book by ID
// @Description Update a book's details by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param   id     path      string  true  "Book ID"
// @Param   book   body      map[string]interface{}  true  "Updated book data"
// @Success 200   {object}  map[string]interface{}  "Book updated successfully"
// @Failure 400   {object}  map[string]interface{}  "Invalid input"
// @Failure 404   {object}  map[string]interface{}  "Book not found"
// @Failure 500   {object}  map[string]interface{}  "Failed to update book"
// @Router /book/{id} [put]
func UpdateBookByID(c *gin.Context) {
	var book models.BookInventory
	bookID := c.Param("id")

	// Check if the book exists
	if err := config.DB.First(&book, "isbn = ?", bookID).Error; err != nil {
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

// DeleteBookByID godoc
// @Summary Delete a book by ID
// @Description Delete a book by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param   id     path      string  true  "Book ID"
// @Success 200   {object}  map[string]interface{}  "Book deleted successfully"
// @Failure 404   {object}  map[string]interface{}  "Book not found"
// @Failure 500   {object}  map[string]interface{}  "Failed to delete book"
// @Router /book/{id} [delete]
func DeleteBookByID(c *gin.Context) {
	var book models.BookInventory
	bookID := c.Param("id")

	// Check if the book exists
	if err := config.DB.First(&book, "isbn = ?", bookID).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Book not found", nil)
		return
	}

	// Delete book
	config.DB.Delete(&book)

	utils.RespondJSON(c, http.StatusOK, "Book deleted successfully", nil)
}

// SearchBooks godoc
// @Summary Search books
// @Description Search books by title, author, or publisher
// @Tags books
// @Accept  json
// @Produce  json
// @Param   title     query    string  false  "Book title"
// @Param   authors   query    string  false  "Book authors"
// @Param   publisher query    string  false  "Book publisher"
// @Success 200   {object}  map[string]interface{}  "Books retrieved successfully"
// @Failure 500   {object}  map[string]interface{}  "Failed to fetch books"
// @Router /search [get]
func SearchBooks(c *gin.Context) {
	var books []models.BookInventory
	title := c.Query("title")
	author := c.Query("authors")
	publisher := c.Query("publisher")

	query := config.DB

	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("authors ILIKE ?", "%"+author+"%")
	}
	if publisher != "" {
		query = query.Where("publisher ILIKE ?", "%"+publisher+"%")
	}

	query.Find(&books)

	utils.RespondJSON(c, http.StatusOK, "Books retrieved successfully", books)
}
