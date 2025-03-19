// middleware/authMiddleware.go
package middlewares

import (
	"library/config"
	"library/models"
	"library/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")[7:]
// 		if token == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
// 			c.Abort()
// 			return
// 		}

// 		claims, err := utils.ParseToken(token)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		var session models.Session
// 		if err := config.DB.Where("token = ? AND is_active = ?", token, true).First(&session).Error; err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired or invalid"})
// 			c.Abort()
// 			return
// 		}

// 		// Set the user ID in the context
// 		c.Set("user_id", claims.UserID)
// 		c.Set("email", claims.Email)
// 		c.Set("role", claims.Role)
// 		c.Next()
// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is empty or too short
		if len(authorizationHeader) < 8 || authorizationHeader[:7] != "Bearer " {
			utils.RespondJSON(c, http.StatusUnauthorized, "Authorization token is required", nil)
			c.Abort()
			return
		}

		// Extract token
		token := authorizationHeader[7:]

		// Check if token is empty
		// if token == "" {
		// 	utils.RespondJSON(c, http.StatusUnauthorized, "Invalid token format", nil)
		// 	c.Abort()
		// 	return
		// }

		// Validate JWT token
		claims, err := utils.ParseToken(token)
		if err != nil {
			utils.RespondJSON(c, http.StatusUnauthorized, "Invalid token", nil)
			c.Abort()
			return
		}

		// Check session in DB
		// var session models.Session
		// if err := config.DB.Where("token = ? AND is_active = ?", token, true).First(&session).Error; err != nil {
		// 	utils.RespondJSON(c, http.StatusUnauthorized, "Session expired or invalid", nil)
		// 	c.Abort()
		// 	return
		// }

		// Set user details in context for future use
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		var currentUser models.User
		config.DB.Model(&models.User{}).Where("email = ?", claims.Email).First(&currentUser)
		c.Set("currentUser", currentUser)

		c.Next()
	}
}
