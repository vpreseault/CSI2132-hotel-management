package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vpreseault/csi2132-project/backend/internal"
	"github.com/vpreseault/csi2132-project/backend/internal/queries"
)

func getCustomerID(r *http.Request) (int, error) {
	param := chi.URLParam(r, "customer_ID")
	customerID, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("provided customer_ID '%v' is not a number", param)
	}

	return customerID, nil
}

func getCustomerBookingsHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		customerID, err := getCustomerID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bookings, err := getCustomerBookings(ctx, customerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(bookings)
	}
}

func getCustomerBookings(ctx *internal.AppContext, customerID int) ([]internal.Booking, error) {
	rows, err := ctx.DB.Query(queries.GetCustomerBookings, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []internal.Booking
	for rows.Next() {
		var booking internal.Booking

		if err := rows.Scan(
			&booking.BookingID,
			&booking.CustomerID,
			&booking.RoomID,
			&booking.BookingDate,
			&booking.StartDate,
			&booking.EndDate,
			&booking.TotalPrice,
		); err != nil {
			return bookings, err
		}
		bookings = append(bookings, booking)
	}

	if err = rows.Err(); err != nil {
		return bookings, err
	}

	return bookings, nil
}

func getCustomerRentings(ctx *internal.AppContext, customerID int) ([]internal.Renting, error) {
	rows, err := ctx.DB.Query(queries.GetCustomerRentings, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rentings []internal.Renting
	for rows.Next() {
		var renting internal.Renting
		var bookingIDOrNull sql.NullInt32

		if err := rows.Scan(
			&renting.RentingID,
			&renting.EmployeeID,
			&renting.CustomerID,
			&renting.RoomID,
			&bookingIDOrNull,
			&renting.CheckInDate,
			&renting.CheckOutDate,
			&renting.Payment,
			&renting.TotalPrice,
		); err != nil {
			return rentings, err
		}

		if bookingIDOrNull.Valid {
			renting.BookingID = int(bookingIDOrNull.Int32)
		}

		rentings = append(rentings, renting)
	}

	if err = rows.Err(); err != nil {
		return rentings, err
	}

	return rentings, nil
}

func getCustomerRentingsHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		customerID, err := getCustomerID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rentings, err := getCustomerRentings(ctx, customerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(rentings)
	}
}

func getCustomerArchivesHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
