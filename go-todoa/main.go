package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Todo 1", Completed: false},
	{ID: "2", Item: "Todo 2", Completed: false},
	{ID: "3", Item: "Todo 3", Completed: false},
	{ID: "4", Item: "Todo 4", Completed: false},
	{ID: "5", Item: "Todo 5", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func todosFromJson(context *gin.Context) {
	jsonFile, err := os.Open("main.json")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(jsonFile)
	for scanner.Scan() {
		var item todo
		json.Unmarshal(scanner.Bytes(), &item)
		todos = append(todos, item)
	}
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, todos)
}

// func getTodoByID(id string) {

// }

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todo/new", addTodo)
	router.GET("/todos/json", todosFromJson)

	router.Run("localhost:9090")
}
