package database

import (
    "database/sql"
    "os"
    "time"

    _ "github.com/lib/pq"
)

const maxRetries = 10
const retryDelay = 5 * time.Second

func ConnectDB() (*sql.DB, error) {
    var db *sql.DB
    var err error

    // Retry connecting to the database
    for i := 0; i < maxRetries; i++ {
        // Attempt to connect to the database
        db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
        if err == nil {
            // Connection successful, return the database object
            return db, nil
        }

        // Connection failed, print error and retry
        time.Sleep(retryDelay)
    }

    // Return the last error encountered if maximum retries reached
    return nil, err
}

func CreateTodoTable(db *sql.DB) error {
    // Create the table if it doesn't exist
    _, err := db.Exec(`CREATE TABLE IF NOT EXISTS todo (
        id SERIAL PRIMARY KEY,
        task TEXT,
        completed BOOLEAN
    )`)
    if err != nil {
        return err
    }
    return nil
}