package main

import (
	"cmp"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
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

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Close the database connection when the program exits

	err = db.Ping() // Check if the connection is working
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	r := chi.NewRouter()

	r.Get("/api/hello", helloHandler)

	port := cmp.Or(os.Getenv("PORT"), "3000")

	var s = &http.Server{
		Addr:              ":" + port,
		Handler:           r,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		ReadHeaderTimeout: 1 * time.Second,
	}

	fmt.Printf("starting server %v", s.Addr)

	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("server exited with error: %v", err)
	}
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
