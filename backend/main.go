package main

import (
	"cmp"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

func createCorsMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		corsMiddleware := cors.Handler(cors.Options{
			AllowedOrigins: allowedOrigins,
		})

		// Print the allowed origins
		log.Printf("CORS Allowed Origins: %v", allowedOrigins)

		return corsMiddleware(next)
	}
}

func main() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	log.Printf("dbURL: %v\n", dbURL)
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Close the database connection when the program exits
	log.Println("Successfully established a connection to the database")

	err = db.Ping() // Check if the connection is working
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	log.Println("Successfully tested connection to the database")

	r := chi.NewRouter()
	allowedOrigins := []string{os.Getenv("ALLOWED_ORIGINS")}
	log.Println("cors.AllowedOrigins: ", allowedOrigins)
	r.Use(createCorsMiddleware(allowedOrigins))

	r.Get("/api/hello", helloHandler)
	r.Get("/api/test", testHandler)

	port := cmp.Or(os.Getenv("PORT"), "3000")

	var s = &http.Server{
		Addr:              ":" + port,
		Handler:           r,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		ReadHeaderTimeout: 1 * time.Second,
	}

	log.Printf("starting server %v\n", s.Addr)

	if err := s.ListenAndServe(); err != nil {
		log.Printf("server exited with error: %v\n", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	log.Println("headers: ", headers)

	startTime := time.Now()
	log.Printf("helloHandler started at: %v\n", startTime)
	w.Header().Set("Content-Type", "application/json")

	queryStartTime := time.Now()
	rows, err := db.Query("SELECT id, name, value FROM items")
	queryEndTime := time.Now()
	log.Printf("Query execution time: %v\n", queryEndTime.Sub(queryStartTime))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Handle errors
		return
	}
	defer rows.Close()

	var items []Item
	rowCount := 0
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
		rowCount++
	}

	log.Printf("Number of rows returned from database: %v", rowCount)

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonStartTime := time.Now()
	json.NewEncoder(w).Encode(items)
	jsonEndTime := time.Now()
	log.Printf("JSON encoding time: %v\n", jsonEndTime.Sub(jsonStartTime))

	endTime := time.Now()
	log.Printf("helloHandler finished at: %v, total time: %v\n", endTime, endTime.Sub(startTime))
}

func testHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Message{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(message)
}
