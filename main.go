package main

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read book", Completed: false},
	{ID: "3", Item: "Learn golang", Completed: false},
}

func getTodos(context *gin.Context) {
	context.JSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.JSON(http.StatusCreated, newTodo)
}


func getTodo(context *gin.Context) {
	id := context.Param("id")
	
	todo, err := getTodoById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{ "message": "Todo not found"})
	}

	context.JSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{ "message": "Todo not found"})
	}

	todo.Completed = !todo.Completed

	context.JSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}