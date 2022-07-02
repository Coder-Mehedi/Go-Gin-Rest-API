package main

import (
	"example/todo-go/controllers"
	"example/todo-go/models"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID	string `json:"id"`
	Title	string `json:"title"`
	Completed	bool `json:"completed"`
}

var todos = []todo{
	{ID: "1", Title: "Task 1", Completed: false},
	{ID: "2", Title: "Task 2", Completed: true},
	{ID: "3", Title: "Task 3", Completed: false},
}



func main(){
	router := gin.Default()
	models.ConnectDatabase()
	router.GET("/todos", controllers.GetTodos)
	router.GET(("/todos/:id"), controllers.GetTodo)
	router.POST("/todos", controllers.CreateTodo)
	router.PATCH("/todos/:id", controllers.UpdateTodo)
	router.PATCH("/todos/:id/toggle", controllers.ToggleTodoStatus)
	router.DELETE("/todos/:id", controllers.DeleteTodo)
	router.Run("localhost:9090")
}
