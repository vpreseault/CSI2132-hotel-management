package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func CreateCorsMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		corsMiddleware := cors.Handler(cors.Options{
			AllowedOrigins: allowedOrigins,
		})

		return corsMiddleware(next)
	}
}
