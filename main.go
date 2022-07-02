package main

import (
	"net/http"

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

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var todo todo
	if e := c.BindJSON(&todo) ; e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}
	todos = append(todos, todo)
	c.JSON(http.StatusCreated, todo)
}

func getTodo(c *gin.Context) {
	id := c.Param("id")
	for _, t := range todos {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	for i, t := range todos {
		if t.ID == id {
			var todo todo
			if e := c.BindJSON(&todo); e != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
				return
			}
			todos[i] = todo
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func toggleTodoStatus(c *gin.Context) {
	id := c.Param("id")
	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = !t.Completed
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func main(){
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET(("/todos/:id"), getTodo)
	router.POST("/todos", createTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.PATCH("/todos/:id/toggle", toggleTodoStatus)
	router.Run("localhost:9090")
}
