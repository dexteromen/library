package middleware

import (
	"net/http"
	"strings"

	"library/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks the JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// Extract token after "Bearer "
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		// Parse the JWT token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set claims in context to use in other handlers
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next() // Continue to the next handler
	}
}
