package controllers

import (
	"library/config"
	"library/models"
	"library/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateLibrary godoc
// @Summary Create a new library
// @Description Create a new library and update the user's role to owner
// @Tags library
// @Accept  json
// @Produce  json
// @Param   library  body      models.Library  true  "Library data"
// @Success 201   {object}  map[string]interface{}  "Library created successfully"
// @Failure 400   {object}  map[string]interface{}  "Invalid input"
// @Failure 404   {object}  map[string]interface{}  "User not found"
// @Failure 500   {object}  map[string]interface{}  "Failed to create library"
// @Router /library [post]
func CreateLibrary(c *gin.Context) {
	var library models.Library
	if err := c.ShouldBindJSON(&library); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	if err := config.DB.Create(&library).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create library", err.Error())
		return
	}

	//find user by id for Updating user lib_id and Role  // reader -> owner
	var user models.User
	if err := config.DB.First(&user, c.GetUint("user_id")).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "User not found", nil)
		return
	}

	user.LibID = library.ID
	user.Role = "owner"

	config.DB.Save(&user)
	// if err := config.DB.Save(&user).Error; err != nil {
	// 	utils.RespondJSON(c, http.StatusInternalServerError, "Failed to update user", nil)
	// 	return
	// }

	//Sending response
	data := gin.H{"library": library, "Owner Of Library": user.Name, "Role": user.Role}

	utils.RespondJSON(c, http.StatusCreated, "Library created successfully", data)
}

// GetLibraries godoc
// @Summary Get all libraries
// @Description Retrieve a list of all libraries
// @Tags library
// @Accept  json
// @Produce  json
// @Success 200   {object}  map[string]interface{}  "Libraries fetched successfully"
// @Failure 500   {object}  map[string]interface{}  "Failed to fetch libraries"
// @Router /library [get]
func GetLibraries(c *gin.Context) {
	var libraries []models.Library
	config.DB.Find(&libraries)
	utils.RespondJSON(c, http.StatusOK, "Libraries fetched successfully", libraries)
}

// func UpdateLibrary(c *gin.Context) {
// 	var library models.Library
// 	libraryID := c.Param("id")

// 	if err := config.DB.First(&library, libraryID).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusNotFound, "Library not found", nil)
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&library); err != nil {
// 		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil)
// 		return
// 	}

// 	if err := config.DB.Save(&library).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to update library", nil)
// 		return
// 	}

// 	utils.RespondJSON(c, http.StatusOK, "Library updated successfully", library)
// }

// func DeleteLibrary(c *gin.Context) {
// 	var library models.Library
// 	libraryID := c.Param("id")

// 	if err := config.DB.First(&library, libraryID).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusNotFound, "Library not found", nil)
// 		return
// 	}

// 	if err := config.DB.Delete(&library).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to delete library", nil)
// 		return
// 	}

// 	utils.RespondJSON(c, http.StatusOK, "Library deleted successfully", nil)
// }
