package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// // Middleware to check role-based access
// func RoleMiddleware(requiredRole string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		role, exists := c.Get("role")
// 		if !exists || role != requiredRole {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

// RoleMiddleware checks if the user has one of the allowed roles
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		// Convert role to string (in case it's stored differently)
		userRole, ok := role.(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		// Check if user role is in allowedRoles
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				c.Next()
				return
			}
		}

		// If no match, reject access
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
	}
}
