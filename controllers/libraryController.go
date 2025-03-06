package controllers

import (
	"library/config"
	"library/models"
	"library/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLibrary(c *gin.Context) {
	var library models.Library
	if err := c.ShouldBindJSON(&library); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	if err := config.DB.Create(&library).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create library", nil)
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Library created successfully", library)
}

func GetLibraries(c *gin.Context) {
	var libraries []models.Library
	config.DB.Find(&libraries)
	utils.RespondJSON(c, http.StatusOK, "Libraries fetched successfully", libraries)
}

func UpdateLibrary(c *gin.Context) {
	var library models.Library
	libraryID := c.Param("id")

	if err := config.DB.First(&library, libraryID).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Library not found", nil)
		return
	}

	if err := c.ShouldBindJSON(&library); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	if err := config.DB.Save(&library).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to update library", nil)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Library updated successfully", library)
}

func DeleteLibrary(c *gin.Context) {
	var library models.Library
	libraryID := c.Param("id")

	if err := config.DB.First(&library, libraryID).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Library not found", nil)
		return
	}

	if err := config.DB.Delete(&library).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to delete library", nil)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Library deleted successfully", nil)
}
