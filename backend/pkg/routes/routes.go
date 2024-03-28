package routes

import (
    "database/sql"
    "net/http"

    "github.com/gorilla/mux"
    "todo-api/pkg/controllers"
)

func SetupRoutes(db *sql.DB) http.Handler {
    // Initialize mux router
    router := mux.NewRouter()

    // Define routes with /api prefix
    api := router.PathPrefix("/api").Subrouter()
    api.HandleFunc("/todos", controllers.GetTodos(db)).Methods("GET")
    api.HandleFunc("/todos", controllers.CreateTodo(db)).Methods("POST")
    api.HandleFunc("/todos/{id}", controllers.UpdateTodo(db)).Methods("PUT")
    api.HandleFunc("/todos/{id}", controllers.DeleteTodo(db)).Methods("DELETE")

    return router
}
