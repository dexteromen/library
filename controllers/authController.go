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

// SignUp godoc
// @Summary Sign up a new user
// @Description Create a new user account
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user  body      models.User  true  "User data"
// @Success 201   {object}  map[string]interface{}  "User created successfully"
// @Failure 400   {object}  map[string]interface{}  "Invalid input"
// @Failure 500   {object}  map[string]interface{}  "Failed to create user"
// @Router /signup [post]
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

	// Validate Password Strength
	if !isValidPassword(user.Password) {
		utils.RespondJSON(c, http.StatusBadRequest, "User can not be created", gin.H{"error": "Password must be at least 8 characters long, contain one uppercase letter, one number, and one special character"})
		return
	}

	// Hash password
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	user.Email = strings.ToLower(user.Email)
	config.DB.Create(&user)

	utils.RespondJSON(c, http.StatusCreated, "User created successfully", gin.H{"user": user})
}

// SignIn updates to store session
type SignInCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignIn godoc
// @Summary Sign in a user
// @Description Authenticate a user and return a JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   credentials  body      SignInCredentials  true  "User credentials"
// @Success 200   {object}  map[string]interface{}  "User logged-in successfully"
// @Failure 400   {object}  map[string]interface{}  "Invalid input"
// @Failure 401   {object}  map[string]interface{}  "Invalid email or password"
// @Failure 500   {object}  map[string]interface{}  "Failed to create session"
// @Router /signin [post]
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

	// Create new session for the user
	session = models.Session{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24), // Session expiry time
		IsActive:  true,
	}

	// Store the new session in the database
	config.DB.Create(&session)
	utils.RespondJSON(c, http.StatusOK, "User logged-in successfully !!", gin.H{"token": token, "expiry_time": session.ExpiresAt})
}

// SignOut godoc
// @Summary Sign out a user
// @Description Invalidate the user's session token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   Authorization  header    string  true  "Bearer token"
// @Success 200   {object}  map[string]interface{}  "User logged out successfully"
// @Failure 401   {object}  map[string]interface{}  "No token provided"
// @Failure 500   {object}  map[string]interface{}  "Database error while logging out"
// @Router /signout [post]
func SignOut(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")

	token := authorizationHeader[7:] // Extract the token

	// Update session status in the database
	config.DB.Model(&models.Session{}).Where("token = ?", token).Update("is_active", false)

	utils.RespondJSON(c, http.StatusOK, "User logged out successfully !!", nil)
}

// GetUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200   {object}  map[string]interface{}  "All users retrieved"
// @Failure 500   {object}  map[string]interface{}  "Failed to fetch users"
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	utils.RespondJSON(c, http.StatusOK, "All users retrived !!", gin.H{"User": users})
}

// GetUserById godoc
// @Summary Get a user by ID
// @Description Retrieve a user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id     path      int     true  "User ID"
// @Success 200   {object}  map[string]interface{}  "User retrieved successfully"
// @Failure 400   {object}  map[string]interface{}  "Invalid user ID"
// @Failure 404   {object}  map[string]interface{}  "User not found"
// @Failure 500   {object}  map[string]interface{}  "Failed to fetch user"
// @Router /user/{id} [get]
func GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	config.DB.First(&user, id)

	utils.RespondJSON(c, http.StatusOK, "User retrived !!", gin.H{"User": user})
}

/*

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

*/

// DeleteUserById godoc
// @Summary Delete a user by ID
// @Description Delete a user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id     path      int     true  "User ID"
// @Success 200   {object}  map[string]interface{}  "User deleted successfully"
// @Failure 400   {object}  map[string]interface{}  "Invalid user ID"
// @Failure 500   {object}  map[string]interface{}  "Failed to delete user"
// @Router /user/{id} [delete]
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
