package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Register(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: Cors.AllowedOrigins,
	}))
}
