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
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	r := chi.NewRouter()

	allowedOrigins := []string{os.Getenv("ALLOWED_ORIGINS")}
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
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, name, value FROM items")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item
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

	json.NewEncoder(w).Encode(items)
}

func testHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Message{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(message)
}
