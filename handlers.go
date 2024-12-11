package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type Todo struct {
    Id       int    `json:"id"`
    Title    string `json:"title"`
    Content  string `json:"content"`
    Finished bool   `json:"finished"`
}

var mu sync.Mutex


// Route:	/create
// Method:	POST
// Body: 	Todo in JSON (excluding id because it is auto-generated by database)
// Return:	Todo in JSON
func CreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo Todo

    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
	// Bind the body into the todo variable
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    stmt, err := db.Prepare("INSERT INTO todos(title, content, finished) VALUES (?, ?, ?)")
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    res, err := stmt.Exec(todo.Title, todo.Content, todo.Finished)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

	id, err := res.LastInsertId()
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    todo.Id = int(id)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo) // Return this new todo after successfully created in the database
}

// Route:	/getall
// Method:	GET
// Return:	Todo[] in JSON
func GetTodos(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    rows, err := db.Query("SELECT id, title, content, finished FROM todos")
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    defer rows.Close()

    var todoList []Todo
    for rows.Next() {
        var todo Todo
        if err := rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Finished); err != nil {
            http.Error(w, "Database error", http.StatusInternalServerError)
            return
        }
        todoList = append(todoList, todo)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todoList) // Return the todo list after successfully query from the database
}

// Route:	/update
// Method:	PUT
// Body: 	Todo in JSON
// Return:	Todo in JSON
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
    var updatedTodo Todo
    if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    stmt, err := db.Prepare("UPDATE todos SET title = ?, content = ?, finished = ? WHERE id = ?")
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    res, err := stmt.Exec(updatedTodo.Title, updatedTodo.Content, updatedTodo.Finished, updatedTodo.Id)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Todo id requested not found", http.StatusNotFound)
		return
	}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedTodo) // Return the updated todo after successfully updated in the database
}

// Route:	/delete?id=43
// Method:	DELETE
// Return:	Status code
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid todo ID", http.StatusBadRequest)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    stmt, err := db.Prepare("DELETE FROM todos WHERE id = ?")
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    res, err := stmt.Exec(id)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Todo id requested not found", http.StatusNotFound)
		return
	}

    w.WriteHeader(http.StatusNoContent)
}

// Route:	/getbyid?id=43
// Method:	GET
// Return:	Todo in JSON
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid todo ID", http.StatusBadRequest)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    var todo Todo
    err = db.QueryRow("SELECT id, title, content, finished FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Finished)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Todo not found", http.StatusNotFound)
        } else {
            http.Error(w, "Database error", http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

// Route:	/getbycompleted
// Method:	GET
// Return:	Todo[] in JSON
func GetTodosByCompleted(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    rows, err := db.Query("SELECT id, title, content, finished FROM todos WHERE finished = true")
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    defer rows.Close()

    var todoList []Todo
    for rows.Next() {
        var todo Todo
        if err := rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Finished); err != nil {
            http.Error(w, "Database error", http.StatusInternalServerError)
            return
        }
        todoList = append(todoList, todo)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todoList)
}