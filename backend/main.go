package main

import (
	"log"
	"net/http"
    "os"

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

    // Get the port from the PORT environment variable provided by Heroku
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    // Start server
    log.Printf("Server listening on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, middleware.JsonContentTypeMiddleware(router)))
}

