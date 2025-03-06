package routes

// "library/config"
// "library/middlewares"
// "library/models"
// "library/utils"
// "net/http"

// "regexp"

// "github.com/gin-gonic/gin"
// "gorm.io/gorm"

// func flowRoutes(router *gin.Engine) {

// 	potectedroutes := router.Group("/")
// 	potectedroutes.Use(middlewares.AuthMiddleware())
// 	{
// 		// potectedroutes.GET("/users", GetUsers)

// 		// potectedroutes.GET("/libraries", Get_Libraries_And_Assigned_Libraries)
// 		// potectedroutes.POST("/createLibraryByOwner", createLibraryByOwner)
// 		// potectedroutes.POST("/assignLibraryByOwnerToAdmin", assignLibraryByOwnerToAdmin)
// 	}
// }

// func createLibraryByOwner(c *gin.Context) {
// 	var library models.Library
// 	if err := c.ShouldBindJSON(&library); err != nil {
// 		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil)
// 		return
// 	}

// 	if err := config.DB.Create(&library).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create library", nil)
// 		return
// 	}

// 	var AssignedLibrary models.AssignedLibrary
// 	AssignedLibrary.LibID = library.ID
// 	AssignedLibrary.UserID = c.GetUint("user_id")
// 	if err := config.DB.Create(&AssignedLibrary).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to assign library to owner", nil)
// 		return
// 	}

// 	//find user by id
// 	var user models.User
// 	if err := config.DB.First(&user, c.GetUint("user_id")).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusNotFound, "User not found", nil)
// 		return
// 	}

// 	//update user lib_id and update Role to owner
// 	user.LibID = library.ID
// 	user.Role = "owner"
// 	if err := config.DB.Save(&user).Error; err != nil {
// 		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to update user", nil)
// 		return
// 	}

// 	utils.RespondJSON(c, http.StatusCreated, "Library created successfully", gin.H{"library": library, "assignedLibrary": AssignedLibrary})
// }

// func Get_Libraries_And_Assigned_Libraries(c *gin.Context) {
// 	var libraries []models.Library
// 	var assignedLibrary []models.AssignedLibrary
// 	config.DB.Find(&libraries)
// 	config.DB.Find(&assignedLibrary)
// 	utils.RespondJSON(c, http.StatusOK, "Libraries fetched successfully", gin.H{"library": libraries, "assignedLibrary": assignedLibrary})
// }
