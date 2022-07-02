package controllers

import (
	"net/http"

	"example/todo-go/models"

	"github.com/gin-gonic/gin"
)

// GET /todos
// Find all Todos
func GetTodos(c *gin.Context) {
	var todos []models.Todo
  models.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	// create to with auth id
	var todo models.Todo
	if e := c.BindJSON(&todo) ; e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}
	models.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
	
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	models.DB.First(&todo, id)
	c.JSON(http.StatusOK, todo)
}
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	models.DB.First(&todo, id)
	if e := c.BindJSON(&todo) ; e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}
	models.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func ToggleTodoStatus(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	models.DB.First(&todo, id)
	todo.Completed = !todo.Completed
	models.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	models.DB.First(&todo, id)
	models.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
}