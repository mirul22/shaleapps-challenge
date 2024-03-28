package controllers

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type Todo struct {
    ID        int    `json:"id"`
    Task      string `json:"task"`
    Completed bool   `json:"completed"`
}

// GetTodos returns all todos
func GetTodos(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM todos ORDER BY id DESC")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		todos := []Todo{}
		for rows.Next() {
			var todo Todo
			if err := rows.Scan(&todo.ID, &todo.Task, &todo.Completed); err != nil {
				log.Fatal(err)
				return
			}
			todos = append(todos, todo)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
			return
		}

		json.NewEncoder(w).Encode(todos)
	}
}

// CreateTodo creates a new todo
func CreateTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			log.Println("Error decoding JSON:", err)
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		log.Println("Received todo:", todo)

		err = db.QueryRow("INSERT INTO todos (task, completed) VALUES ($1, $2) RETURNING id", todo.Task, todo.Completed).Scan(&todo.ID)
		if err != nil {
			log.Println("Error executing SQL query:", err)
			http.Error(w, "Error executing SQL query", http.StatusInternalServerError)
			return
		}

		log.Println("Todo inserted successfully")

		json.NewEncoder(w).Encode(todo)
	}
}

// UpdateTodo updates a todo
func UpdateTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		json.NewDecoder(r.Body).Decode(&todo)

		vars := mux.Vars(r)
		id := vars["id"]

		err := db.QueryRow("UPDATE todos SET task=$1, completed=$2 WHERE id=$3 RETURNING id", todo.Task, todo.Completed, id).Scan(&todo.ID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(todo)
	}
}

// DeleteTodo deletes a todo
func DeleteTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		result, err := db.Exec("DELETE FROM todos WHERE id=$1", id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode("Todo deleted successfully")
	}
}
