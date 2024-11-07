package utils

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDatabase() error {
    // Define connection parameters with logging
	connStr := "host=localhost port=5432 user=postgres dbname=osds password=1234 sslmode=disable"
    fmt.Printf("Attempting to connect with: %s\n", connStr)

    var err error
    // Open database connection
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        fmt.Printf("Error opening database: %v\n", err)
        return fmt.Errorf("error opening database: %v", err)
    }
    fmt.Println("Successfully opened database")

    // Test the connection
    err = DB.Ping()
    if err != nil {
        fmt.Printf("Error pinging database: %v\n", err)
        return fmt.Errorf("error connecting to the database: %v", err)
    }
    fmt.Println("Successfully pinged database")

    // Set connection pool settings
    DB.SetMaxOpenConns(25)
    DB.SetMaxIdleConns(5)
    DB.SetConnMaxLifetime(time.Minute * 5)
    fmt.Println("Set database connection pool parameters")

    fmt.Println("Database connection fully configured and ready")
    return nil
}

func CloseDatabase() error {
    if DB != nil {
        fmt.Println("Closing database connection")
        return DB.Close()
    }
    fmt.Println("No database connection to close")
    return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
    return DB
}

// TestConnection tests if the database connection is working
func TestConnection() error {
    if DB == nil {
        return fmt.Errorf("database connection is nil")
    }
    return DB.Ping()
}