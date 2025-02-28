package controllers

import (
	"github.com/dexteromen/library/initializers"
	"github.com/dexteromen/library/models"
	"github.com/gin-gonic/gin"
)

func TodosCreate(c *gin.Context) {
	// Get data from req body
	var body struct {
		Content string
		Status  bool
	}
	c.Bind(&body)

	// Create a todo
	todo := models.Todo{
		Content: body.Content,
		Status:  body.Status,
	}
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosIndex(c *gin.Context) {
	// Get all the todos
	var todos []models.Todo
	initializers.DB.Find(&todos)

	// Return todos in response
	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func TodosShow(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// Get a sing todo
	var todo models.Todo
	initializers.DB.First(&todo, id)

	// Return todo in response
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosUpdate(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// get the data of req body
	var body struct {
		Content string
		Status  bool
	}
	c.Bind(&body)

	// Get a single todo that we want to update
	var todo models.Todo
	initializers.DB.First(&todo, id)

	// Update it
	initializers.DB.Model(&todo).Updates(models.Todo{
		Content: body.Content,
		Status:  body.Status,
	})

	// Return response
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func TodosDelete(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// Delete the Todo
	initializers.DB.Delete(&models.Todo{}, id)

	// Return response
	c.JSON(200, gin.H{
		"message": "Todo removed Successfully",
	})
}
