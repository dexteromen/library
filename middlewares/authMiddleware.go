// middleware/authMiddleware.go
package middleware

import (
	"net/http"
	"library/config"
	"library/models"
	"library/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")[7:]
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		var session models.Session
		if err := config.DB.Where("token = ? AND is_active = ?", token, true).First(&session).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired or invalid"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Next()
	}
}