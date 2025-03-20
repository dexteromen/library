package controllers

import (
	// "library/config"
	"library/config"
	"library/models"
	"time"

	"library/utils"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Assuming middleware sets currentUser in context
func GetProfile(c *gin.Context) {
	var currentUser models.User

	userDetails, exists := c.Get("currentUser")
	if !exists {
		utils.RespondJSON(c, http.StatusUnauthorized, "User not found in context", nil)
		return
	}
	user, ok := userDetails.(models.User)
	if !ok {
		utils.RespondJSON(c, http.StatusInternalServerError, "Error asserting user details", nil)
		return
	}

	currentUser = models.User{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		ContactNumber: user.ContactNumber,
		// Password:      user.Password,
		Role:  user.Role,
		LibID: user.LibID,
	}

	utils.RespondJSON(c, http.StatusOK, "Profile Retrieved.", currentUser)
}

func RefreshToken(c *gin.Context) {
    // Extract the token from the request header
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        utils.RespondJSON(c, http.StatusUnauthorized, "Authorization header is missing", nil)
        return
    }

    // Remove "Bearer " prefix if present
    if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
        tokenString = tokenString[7:]
    }

    // Parse the token
    claims, err := utils.ParseToken(tokenString)
    if err != nil {
        utils.RespondJSON(c, http.StatusUnauthorized, "Invalid token", nil)
        return
    }

    // Create a new token with the same claims but a new expiration time
    expirationTime := time.Now().Add(15 * time.Minute)
    claims.ExpiresAt = jwt.NewNumericDate(expirationTime)

    newTokenString, err := utils.GenerateJWT(claims.UserID, claims.Email, claims.Role)
    if err != nil {
        utils.RespondJSON(c, http.StatusInternalServerError, "Could not generate new token", nil)
        return
    }

    // Send the new token back to the client
    utils.RespondJSON(c, http.StatusOK, "Token refreshed successfully", gin.H{"token": newTokenString})
}

func GetProfileByToken(c *gin.Context) {
    // Extract the token from the request header
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        utils.RespondJSON(c, http.StatusUnauthorized, "Authorization header is missing", nil)
        return
    }

    // Remove "Bearer " prefix if present
    if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
        tokenString = tokenString[7:]
    }

    // Parse the token
    claims, err := utils.ParseToken(tokenString)
    if err != nil {
        utils.RespondJSON(c, http.StatusUnauthorized, "Invalid token", nil)
        return
    }

    // Retrieve user details from the database using the user ID from the token claims
    var user models.User
    if err := config.DB.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
        utils.RespondJSON(c, http.StatusInternalServerError, "User not found", nil)
        return
    }

    // Respond with the user details
    utils.RespondJSON(c, http.StatusOK, "User details retrieved successfully", user)
}