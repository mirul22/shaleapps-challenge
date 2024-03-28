package database

import (
    "database/sql"
    "os"

    _ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
    // Connect to database
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        return nil, err
    }

    // Check if the connection is established
    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}

func CreateTodoTable(db *sql.DB) error {
    // Create the table if it doesn't exist
    _, err := db.Exec(`CREATE TABLE IF NOT EXISTS todos (
        id SERIAL PRIMARY KEY,
        task TEXT,
        completed BOOLEAN
    )`)
    if err != nil {
        return err
    }
    return nil
}
