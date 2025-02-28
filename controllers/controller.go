package controllers

import (
	"github.com/dexteromen/library/initializers"
	"github.com/dexteromen/library/models"
	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	// "log"
	// "time"
	"net/http"
)


// Helper function to return JSON responses
func JSONResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{
		"status": status,
		"data":   data,
	})
}

// =====================================================
// Library Controller
// =====================================================

func  CreateLibrary(c *gin.Context) {
	var library models.Library
	if err := c.ShouldBindJSON(&library); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Create(&library).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	JSONResponse(c, http.StatusCreated, library)
}

func  GetLibrary(c *gin.Context) {
	var library models.Library
	if err := initializers.DB.First(&library, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Library not found"})
		return
	}
	JSONResponse(c, http.StatusOK, library)
}

func  UpdateLibrary(c *gin.Context) {
	var library models.Library
	if err := initializers.DB.First(&library, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Library not found"})
		return
	}

	if err := c.ShouldBindJSON(&library); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Save(&library).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	JSONResponse(c, http.StatusOK, library)
}

func  DeleteLibrary(c *gin.Context) {
	if err := initializers.DB.Delete(&models.Library{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Library not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Library deleted successfully"})
}

// =====================================================
// User Controller
// =====================================================

func  CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	JSONResponse(c, http.StatusCreated, user)
}

func  GetUser(c *gin.Context) {
	var user models.User
	if err := initializers.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	JSONResponse(c, http.StatusOK, user)
}

func  UpdateUser(c *gin.Context) {
	var user models.User
	if err := initializers.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	JSONResponse(c, http.StatusOK, user)
}

func  DeleteUser(c *gin.Context) {
	if err := initializers.DB.Delete(&models.User{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// =====================================================
// BookInventory Controller
// =====================================================

func  CreateBook(c *gin.Context) {
	var book models.BookInventory
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	JSONResponse(c, http.StatusCreated, book)
}

func  GetBook(c *gin.Context) {
	var book models.BookInventory
	if err := initializers.DB.First(&book, "isbn = ?", c.Param("isbn")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	JSONResponse(c, http.StatusOK, book)
}

func  UpdateBook(c *gin.Context) {
	var book models.BookInventory
	if err := initializers.DB.First(&book, "isbn = ?", c.Param("isbn")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	JSONResponse(c, http.StatusOK, book)
}

func  DeleteBook(c *gin.Context) {
	if err := initializers.DB.Delete(&models.BookInventory{}, "isbn = ?", c.Param("isbn")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// =====================================================
// RequestEvent Controller
// =====================================================

func  CreateRequestEvent(c *gin.Context) {
	var reqEvent models.RequestEvent
	if err := c.ShouldBindJSON(&reqEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&reqEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	JSONResponse(c, http.StatusCreated, reqEvent)
}

// =====================================================
// IssueRegistry Controller
// =====================================================

func  CreateIssueRegistry(c *gin.Context) {
	var issue models.IssueRegistry
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	JSONResponse(c, http.StatusCreated, issue)
}


