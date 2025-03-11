package controllers

import (
	"library/config"
	"library/models"
	"library/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"regexp"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

// Custom validation for name (only characters allowed)
func isValidName(name string) bool {
	return regexp.MustCompile(`^[A-Za-z\s]+$`).MatchString(name)
}

// Custom password validation (at least 8 characters, 1 uppercase, 1 digit, 1 special character)
func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasUpper := false
	hasDigit := false
	hasSpecial := false
	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case '0' <= char && char <= '9':
			hasDigit = true
		case strings.ContainsRune("@$!%*?&", char):
			hasSpecial = true
		}
	}
	return hasUpper && hasDigit && hasSpecial
}

// // isValidEmail checks if an email follows the correct format
// func isValidEmail(email string) bool {
// 	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
// 	return regexp.MustCompile(pattern).MatchString(email)
// }

// Signup user
func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", gin.H{"error": err.Error()})
		return
	}

	if user.Role == "admin" {
		var existingAdmin models.User
		if err := config.DB.First(&existingAdmin, "role = ?", "admin").Error; err == nil {
			utils.RespondJSON(c, http.StatusBadRequest, "Cannot create more than one admin", nil)
			return
		}
	}

	// Validate Name
	if !isValidName(user.Name) {
		utils.RespondJSON(c, http.StatusBadRequest, "User can not be created", gin.H{"error": "Name must contain only letters and spaces"})
		return
	}

	// Validate Email
	// if !isValidEmail(user.Email) {
	// 	utils.RespondJSON(c, http.StatusBadRequest, "User can not be created", gin.H{"error": "Invalid email format"})
	// 	return
	// }

	// Validate Password Strength
	if !isValidPassword(user.Password) {
		utils.RespondJSON(c, http.StatusBadRequest, "User can not be created", gin.H{"error": "Password must be at least 8 characters long, contain one uppercase letter, one number, and one special character"})
		return
	}

	// Hash password
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	config.DB.Create(&user)
	// // Save to DB
	// if err := config.DB.Create(&user).Error; err != nil {
	// 	utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create user", gin.H{"error": err.Error()})
	// 	return
	// }

	utils.RespondJSON(c, http.StatusCreated, "User created successfully", gin.H{"user": user})
}

// SignIn updates to store session
type SignInCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignIn(c *gin.Context) {
	var credentials SignInCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	// Check if user exists in DB by email
	var user models.User
	if err := config.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		utils.RespondJSON(c, http.StatusUnauthorized, "Invalid email", nil)
		return
	}

	// Validate password
	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		utils.RespondJSON(c, http.StatusUnauthorized, "Invalid password", nil)
		return
	}

	// Generate JWT Token
	token, _ := utils.GenerateJWT(user.ID, user.Email, user.Role)
	// if err != nil {
	// 	utils.RespondJSON(c, http.StatusInternalServerError, "Failed to generate token", nil)
	// 	return
	// }

	// Check if there's an existing session, and update or create a new one
	var session models.Session
	// Check if there is an existing active session
	config.DB.Where("user_id = ? AND is_active = ?", user.ID, true)
	// if err := config.DB.Where("user_id = ? AND is_active = ?", user.ID, true).First(&session).Error; err == nil {
	// 	// If a session exists, invalidate it first (or update)
	// 	session.IsActive = false
	// 	config.DB.Save(&session)
	// }

	// Create new session for the user
	session = models.Session{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24), // Session expiry time
		IsActive:  true,
	}

	// Store the new session in the database
	config.DB.Create(&session)
	// if err := config.DB.Create(&session).Error; err != nil {
	// 	utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create session", nil)
	// 	return
	// }

	utils.RespondJSON(c, http.StatusOK, "User logged-in successfully !!", gin.H{"token": token, "expiry_time": session.ExpiresAt})
}

// Logout handler
func SignOut(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")

	// Check if the Authorization header is empty or too short
	// if len(authorizationHeader) < 8 || authorizationHeader[:7] != "Bearer " {
	// 	utils.RespondJSON(c, http.StatusUnauthorized, "No token provided !!", nil)
	// 	return
	// }

	token := authorizationHeader[7:] // Extract the token

	// Check if the token is actually empty
	// if token == "" {
	// 	utils.RespondJSON(c, http.StatusUnauthorized, "No token provided !!", nil)
	// 	return
	// }

	// Update session status in the database
	config.DB.Model(&models.Session{}).Where("token = ?", token).Update("is_active", false)

	// result := config.DB.Model(&models.Session{}).Where("token = ?", token).Update("is_active", false)
	// if result.Error != nil {
	// 	utils.RespondJSON(c, http.StatusInternalServerError, "Database error while logging out", nil)
	// 	return
	// }

	utils.RespondJSON(c, http.StatusOK, "User logged out successfully !!", nil)
}

// GetUsers handles GET requests to fetch all users
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	// if err := config.DB.Find(&users).Error; err != nil {
	// 	utils.RespondJSON(c, http.StatusInternalServerError, "Failed to fetch users", nil)
	// 	return
	// }
	// if len(users) <= 0 {
	// 	utils.RespondJSON(c, http.StatusOK, "No users found in datatbase !!", nil)
	// 	return
	// }
	utils.RespondJSON(c, http.StatusOK, "All users retrived !!", gin.H{"User": users})
}

// GetUserById handles GET requests to fetch a user by ID
func GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	utils.RespondJSON(c, http.StatusBadRequest, "Error: Invalid user ID", nil)
	// 	return
	// }

	var user models.User
	config.DB.First(&user, id)
	// if err := config.DB.First(&user, id).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		utils.RespondJSON(c, http.StatusNotFound, "Error: User not found", nil)
	// 	} else {
	// 		utils.RespondJSON(c, http.StatusInternalServerError, "Error: Failed to fetch user", nil)
	// 	}
	// 	return
	// }

	utils.RespondJSON(c, http.StatusOK, "User retrived !!", gin.H{"User": user})
}

// // UpdateUserById handles PUT requests to update a user by ID
// func UpdateUserById(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	// if err != nil {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
// 	// 	return
// 	// }

// 	var user models.User
// 	config.DB.First(&user, id)
// 	// if err := config.DB.First(&user, id).Error; err != nil {
// 	// 	if err == gorm.ErrRecordNotFound {
// 	// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 	// 	} else {
// 	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
// 	// 	}
// 	// 	return
// 	// }

// 	var updatedUser models.User
// 	// if err := c.ShouldBindJSON(&updatedUser); err != nil {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 	// 	return
// 	// }

// 	// Validate Name
// 	if !isValidName(updatedUser.Name) {
// 		// c.JSON(http.StatusBadRequest, gin.H{"error": "Name must contain only letters and spaces"})
// 		return
// 	}

// 	// Validate Email
// 	if !isValidEmail(updatedUser.Email) {
// 		// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
// 		return
// 	}

// 	// Validate Password Strength
// 	if !isValidPassword(updatedUser.Password) {
// 		// c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters long, contain one uppercase letter, one number, and one special character"})
// 		return
// 	}

// 	// Update only provided fields
// 	user.Name = updatedUser.Name
// 	user.Email = updatedUser.Email
// 	user.ContactNumber = updatedUser.ContactNumber
// 	user.Role = updatedUser.Role

// 	// Update password only if provided
// 	if updatedUser.Password != "" {
// 		hashedPassword, _ := utils.HashPassword(updatedUser.Password)
// 		// if err != nil {
// 		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
// 		// 	return
// 		// }
// 		user.Password = hashedPassword
// 	}
// 	config.DB.Save(&user)
// 	// if err := config.DB.Save(&user).Error; err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
// 	// 	return
// 	// }

// 	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
// }

// DeleteUserById handles DELETE requests to delete a user by ID
func DeleteUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	// 	return
	// }
	config.DB.Delete(&models.User{}, id)
	// if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
