package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/vpreseault/csi2132-project/backend/internal"
	"github.com/vpreseault/csi2132-project/backend/internal/queries"
)

func InitHandlers(r *chi.Mux, ctx *internal.AppContext) {
	// Auth
	r.Post("/api/customers", createCustomerHandler(ctx))
	r.Get("/api/customers/{name}", getCustomerHandler(ctx))
	r.Post("/api/employees", createEmployeeHandler(ctx))
	r.Get("/api/employees", getEmployeesHandler(ctx))

	// Search
	r.Post("/api/search", RoomSearchHandler(ctx))

	// Activity
	r.Get("/api/activity", getActivityHandler(ctx))
	r.Get("/api/bookings", getBookingsHandler(ctx))
	r.Get("/api/rentings", getRentingsHandler(ctx))
	r.Get("/api/archives", getArchivesHandler(ctx))

	// Admin
	r.Delete("/api/chain/{chain_ID}", deleteChainByID(ctx))
	r.Delete("/api/hotel/{hotel_ID}", deleteHotelByID(ctx))
	r.Delete("/api/room/{room_ID}", deleteRoomByID(ctx))
}

func getHotelByEmployeeID(ctx *internal.AppContext, employeeID int) (int, error) {
	var hotelID int
	err := ctx.DB.QueryRow(queries.GetEmployeeHotelID, employeeID).Scan(&hotelID)
	if err != nil {
		return 0, err
	}

	return hotelID, nil
}
