package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type Message struct {
	Message string `json:"message"`
}

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

var db *sql.DB

func main() {
	var err error
	// Database connection details (replace with your own)
	connStr := "postgres://tuser:password123@localhost:5432/testdb?sslmode=disable" // Update!
	// connStr := "postgres://user:password@host:port/database?sslmode=disable" // Update!
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Close the database connection when the program exits

	err = db.Ping() // Check if the connection is working
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	http.HandleFunc("/api/hello", helloHandler)

	fmt.Println("Backend server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var items []Item
	rows, err := db.Query("SELECT id, name, value FROM items") // Query the database
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Handle errors
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items) // Send the items as JSON
}
