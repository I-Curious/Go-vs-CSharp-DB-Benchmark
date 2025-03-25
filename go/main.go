package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Database connection string for Docker container to access PostgreSQL on the host machine
const connStr = "postgres://postgres:<password>@host.docker.internal:5432/postgres?sslmode=disable"

// Global variable for the database connection
var db *sql.DB

// Handler for the root route
func handler(w http.ResponseWriter, r *http.Request) {
	// Querying the database to get all records from the users table
	rows, err := db.Query("SELECT id, name, email FROM users") // Update with the correct columns
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Writing the query result to the response
	var id int
	var name, email string
	for rows.Next() {
		if err := rows.Scan(&id, &name, &email); err != nil {
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
		// Display user details (id, name, email)
		fmt.Fprintf(w, "ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	// Check if any errors occurred during iteration
	if err := rows.Err(); err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
}

// Main function
func main() {
	// Establishing a connection to the PostgreSQL database
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	// Check if the database connection is available
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}
	fmt.Println("Connected to the PostgreSQL database successfully!")

	// Handling HTTP requests
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
