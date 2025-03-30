package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/vpreseault/csi2132-project/backend/internal"
	"github.com/vpreseault/csi2132-project/backend/internal/queries"
)

func InitHandlers(r *chi.Mux, ctx *internal.AppContext) {
	// Users
	r.Post("/api/customers", createCustomerHandler(ctx))
	r.Get("/api/customers/{name}", getCustomerHandler(ctx))
	r.Patch("/api/customers/{customer_ID}", updateCustomerByID(ctx))
	r.Delete("/api/customers", deleteCustomerByID(ctx))

	r.Post("/api/employees", createEmployeeHandler(ctx))
	r.Get("/api/employees", getEmployeesHandler(ctx))
	r.Patch("/api/employees/{employee_ID}", updateEmployeeByID(ctx))
	r.Delete("/api/employees", deleteEmployeeByID(ctx))

	// Search
	r.Post("/api/search", RoomSearchHandler(ctx))

	// Activity
	r.Get("/api/activity", getActivityHandler(ctx))
	r.Get("/api/bookings", getBookingsHandler(ctx))
	r.Get("/api/rentings", getRentingsHandler(ctx))
	r.Post("/api/renting-from-booking", createRentingFromBookingHandler(ctx))
	r.Get("/api/archives", getArchivesHandler(ctx))

	// Admin
	r.Get("/api/chains", getChainsHandler(ctx))
	r.Delete("/api/chains/{chain_ID}", deleteChainByID(ctx))

	r.Get("/api/hotels", getHotelsHandler(ctx))
	r.Delete("/api/hotels/{hotel_ID}", deleteHotelByID(ctx))
	r.Patch("/api/hotels/{hotel_ID}", updateHotel(ctx))
	// r.Get("/api/rooms", getRoomsHandler(ctx))
	r.Delete("/api/rooms/{room_ID}", deleteRoomByID(ctx))
}

func getHotelByEmployeeID(ctx *internal.AppContext, employeeID int) (int, error) {
	var hotelID int
	err := ctx.DB.QueryRow(queries.GetEmployeeHotelID, employeeID).Scan(&hotelID)
	if err != nil {
		return 0, err
	}

	return hotelID, nil
}
