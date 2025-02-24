package main

import (
	"cmp"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/vpreseault/csi2132-project/backend/internal"
	"github.com/vpreseault/csi2132-project/backend/internal/handlers"
	"github.com/vpreseault/csi2132-project/backend/internal/middleware"
	"github.com/vpreseault/csi2132-project/backend/internal/storage"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	dbConn, err := storage.Connect(dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer dbConn.Close()
	ctx := &internal.AppContext{DB: dbConn}

	r := chi.NewRouter()

	allowedOrigins := []string{os.Getenv("ALLOWED_ORIGINS")}
	r.Use(middleware.CreateCorsMiddleware(allowedOrigins))

	handlers.InitHandlers(r, ctx)

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
