package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"todo-api/pkg/database"
	"todo-api/pkg/routes"
	"todo-api/pkg/middleware"
)

type Todo struct {
	ID        int    `json:"id"`
	Task	  string `json:"task"`
	Completed bool   `json:"completed"`
}

func main() {
	// Connect to database
    db, err := database.ConnectDB()
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }
    defer db.Close()

    // Create todo table
    err = database.CreateTodoTable(db)
    if err != nil {
        log.Fatal("Error creating todo table:", err)
    }

	// Get router from routes package
    router := routes.SetupRoutes(db)

    // Start server
    log.Fatal(http.ListenAndServe(":8000", middleware.JsonContentTypeMiddleware(router)))
}


