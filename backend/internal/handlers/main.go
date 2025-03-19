package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/vpreseault/csi2132-project/backend/internal"
)

func InitHandlers(r *chi.Mux, ctx *internal.AppContext) {
	// Auth
	r.Post("/api/customers", createCustomerHandler(ctx))
	r.Get("/api/customers/{name}", getCustomerHandler(ctx))
	r.Post("/api/employees", createEmployeeHandler(ctx))
	r.Get("/api/employees/{name}", getEmployeeHandler(ctx))

	// Activity
	r.Get("/api/activity/{customer_ID}", getCustomerActivityHandler(ctx))
	r.Get("/api/bookings/{customer_ID}", getCustomerBookingsHandler(ctx))
	r.Get("/api/rentings/{customer_ID}", getCustomerRentingsHandler(ctx))
	r.Get("/api/archives/{customer_ID}", getCustomerArchivesHandler(ctx))
}
