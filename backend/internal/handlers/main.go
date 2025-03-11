package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/vpreseault/csi2132-project/backend/internal"
)

func InitHandlers(r *chi.Mux, ctx *internal.AppContext) {
	r.Post("/api/customers", createCustomerHandler(ctx))
	r.Get("/api/customers/{name}", getCustomerHandler(ctx))
	r.Post("/api/employees", createEmployeeHandler(ctx))
	r.Get("/api/employees/{name}", getEmployeeHandler(ctx))
}
