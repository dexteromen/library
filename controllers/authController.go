package controllers

import (
	// "fmt"
	"library/config"
	"library/models"
	"library/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"regexp"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// isValidEmail checks if an email follows the correct format
func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(pattern).MatchString(email)
}

// Signup user
func SignUp(c *gin.Context) {
	var user models.User
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		// fmt.Println()
		// fmt.Println("Binding Error:", err)
		// fmt.Println()
		// fmt.Println("Received JSON:", c.Request.Body) // Debugging output
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Show the actual error
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", gin.H{"error": err.Error()})
		return
	}
	// fmt.Println("Password:", user.Password, "Valid:", isValidPassword(user.Password))

	// Validate Name
	if !isValidName(user.Name) {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Name must contain only letters and spaces"})
		utils.RespondJSON(c, http.StatusBadRequest, "User can not be created", gin.H{"error": "Name must contain only letters and spaces"})
		return
	}

	// Validate Email
	if !isValidEmail(user.Email) {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		utils.RespondJSON(c, http.StatusBadRequest, "User can not be created", gin.H{"error": "Invalid email format"})
		return
	}

	// Validate Password Strength
	if !isValidPassword(user.Password) {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters long, contain one uppercase letter, one number, and one special character"})
		utils.RespondJSON(c, http.StatusBadRequest, "User can not be created", gin.H{"error": "Password must be at least 8 characters long, contain one uppercase letter, one number, and one special character"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to hash password", gin.H{"error": err.Error()})
		return
	}
	user.Password = hashedPassword

	// Save to DB
	if err := config.DB.Create(&user).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user :-" + err.Error()})
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create user", gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	utils.RespondJSON(c, http.StatusCreated, "User created successfully", gin.H{"user": user})
}

// SignIn updates to store session
type SignInCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context) {
	var credentials SignInCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	// Check if user exists in DB by email
	var user models.User
	if err := config.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		utils.RespondJSON(c, http.StatusUnauthorized, "Invalid email", nil)
		return
	}

	// Validate password
	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		utils.RespondJSON(c, http.StatusUnauthorized, "Invalid password", nil)
		return
	}

	// Generate JWT Token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	// Check if there's an existing session, and update or create a new one
	var session models.Session
	// Check if there is an existing active session
	if err := config.DB.Where("user_id = ? AND is_active = ?", user.ID, true).First(&session).Error; err == nil {
		// If a session exists, invalidate it first (or update)
		session.IsActive = false
		config.DB.Save(&session)
	}

	// Create new session for the user
	session = models.Session{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 1), // Session expiry time
		IsActive:  true,
	}

	// Store the new session in the database
	if err := config.DB.Create(&session).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create session", nil)
		return
	}

	// Send response with token and session expiry
	// c.JSON(http.StatusOK, gin.H{
	// 	"token":       token,
	// 	"expiry_time": session.ExpiresAt,
	// 	"message":     "User logged-in successfully",
	// })

	utils.RespondJSON(c, http.StatusOK, "User logged-in successfully !!", gin.H{"token": token, "expiry_time": session.ExpiresAt})
}

// Logout handler
func SignOut(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")

	// Check if the Authorization header is empty or too short
	if len(authorizationHeader) < 8 || authorizationHeader[:7] != "Bearer " {
		utils.RespondJSON(c, http.StatusUnauthorized, "No token provided !!", nil)
		return
	}

	token := authorizationHeader[7:] // Extract the token

	// Check if the token is actually empty
	if token == "" {
		utils.RespondJSON(c, http.StatusUnauthorized, "No token provided !!", nil)
		return
	}

	// Update session status in the database
	result := config.DB.Model(&models.Session{}).Where("token = ?", token).Update("is_active", false)
	if result.Error != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Database error while logging out", nil)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "User logged out successfully !!", nil)
}

/*
// // AdminIndex - Restricted route for admin users
// func AdminIndex(c *gin.Context) {
// 	role, exists := c.Get("role")
// 	if !exists || role != "admin" {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Welcome, Admin!"})
// }

// // UserIndex - Restricted route for regular users
// func UserIndex(c *gin.Context) {
// 	role, exists := c.Get("role")
// 	if !exists || role != "user" {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Welcome, User!"})
// }
*/

// GetUsers handles GET requests to fetch all users
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to fetch users", nil)
		return
	}
	if len(users) <= 0 {
		utils.RespondJSON(c, http.StatusOK, "No users found in datatbase !!", nil)
	}
	// c.JSON(http.StatusOK, users)
	utils.RespondJSON(c, http.StatusOK, "All users retrived !!", gin.H{"User": users})
}

// GetUserById handles GET requests to fetch a user by ID
func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		utils.RespondJSON(c, http.StatusBadRequest, "Error: Invalid user ID", nil)
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			utils.RespondJSON(c, http.StatusNotFound, "Error: User not found", nil)
		} else {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
			utils.RespondJSON(c, http.StatusInternalServerError, "Error: Failed to fetch user", nil)
		}
		return
	}

	// c.JSON(http.StatusOK, user)
	utils.RespondJSON(c, http.StatusOK, "User retrived !!", gin.H{"User": user})
}

// UpdateUserById handles PUT requests to update a user by ID
func UpdateUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		}
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate Name
	if !isValidName(updatedUser.Name) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name must contain only letters and spaces"})
		return
	}

	// Validate Email
	if !isValidEmail(updatedUser.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Validate Password Strength
	if !isValidPassword(updatedUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters long, contain one uppercase letter, one number, and one special character"})
		return
	}

	// Update only provided fields
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.ContactNumber = updatedUser.ContactNumber
	user.Role = updatedUser.Role

	// Update password only if provided
	if updatedUser.Password != "" {
		hashedPassword, err := utils.HashPassword(updatedUser.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = hashedPassword
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

// DeleteUserById handles DELETE requests to delete a user by ID
func DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
