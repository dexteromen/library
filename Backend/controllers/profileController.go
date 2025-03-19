package controllers

import (
	// "library/config"
	"library/models"

	"library/utils"
	"net/http"

	// "time"

	"github.com/gin-gonic/gin"
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

// func GetProfileByToken(c *gin.Context){
// }
