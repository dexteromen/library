package controllers

import (
	"time"
	"library/config"
	"library/models"
	"library/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup user
func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	// Save to DB
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// SignIn updates to store session
func SignIn(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	session := models.Session{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 1),
		IsActive:  true,
	}
	config.DB.Create(&session)

	c.JSON(http.StatusOK, gin.H{"token": token, "message": "User logged-in"})
}

// Logout handler
func SignOut(c *gin.Context) {
	token := c.GetHeader("Authorization")[7:]

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	config.DB.Model(&models.Session{}).Where("token = ?", token).Update("is_active", false)
	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

// AdminIndex - Restricted route for admin users
func AdminIndex(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome, Admin!"})
}

// UserIndex - Restricted route for regular users
func UserIndex(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome, User!"})
}
