package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vpreseault/csi2132-project/backend/internal"
)

func InitHandlers(r *chi.Mux, ctx *internal.AppContext) {
	r.Get("/api/hello", helloHandler(ctx))
	r.Get("/api/test", testHandler)
	r.Post("/api/customers", createCustomerHandler(ctx))
}

type Message struct {
	Message string `json:"message"`
}

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func helloHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		rows, err := ctx.DB.Query("SELECT id, name, value FROM items")

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
}

func testHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Message{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(message)
}
