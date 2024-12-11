package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
    initDB()
    defer db.Close()

    http.HandleFunc("GET /getall", GetTodos)
    http.HandleFunc("POST /create", CreateTodo)
    http.HandleFunc("PUT /update", UpdateTodo)
    http.HandleFunc("DELETE /delete", DeleteTodo)
	http.HandleFunc("GET /getbyid", GetTodoByID)
	http.HandleFunc("GET /getbycompleted", GetTodosByCompleted)

    fmt.Println("Server started at port 8080")

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting the server:", err)
    }
}