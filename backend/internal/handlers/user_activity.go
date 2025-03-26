package handlers

import (
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

func getCustomerBookings(ctx *internal.AppContext, customerID int) ([]internal.BookingDisplay, error) {
	rows, err := ctx.DB.Query(queries.GetCustomerBookings, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []internal.BookingDisplay
	for rows.Next() {
		var booking internal.BookingDisplay

		if err := rows.Scan(
			&booking.BookingID,
			&booking.CustomerName,
			&booking.RoomNumber,
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

func getCustomerRentings(ctx *internal.AppContext, customerID int) ([]internal.RentingDisplay, error) {
	rows, err := ctx.DB.Query(queries.GetCustomerRentings, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rentings []internal.RentingDisplay
	for rows.Next() {
		var renting internal.RentingDisplay

		if err := rows.Scan(
			&renting.RentingID,
			&renting.EmployeeName,
			&renting.CustomerName,
			&renting.RoomNumber,
			&renting.StartDate,
			&renting.EndDate,
			&renting.Payment,
			&renting.TotalPrice,
		); err != nil {
			return rentings, err
		}

		rentings = append(rentings, renting)
	}

	if err = rows.Err(); err != nil {
		return rentings, err
	}

	return rentings, nil
}

func getCustomerArchivesHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		customerID, err := getCustomerID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		archives, err := getCustomerArchives(ctx, customerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(archives)
	}
}

func getCustomerArchives(ctx *internal.AppContext, customerID int) ([]internal.ArchiveDisplay, error) {
	rows, err := ctx.DB.Query(queries.GetCustomerRentingArchives, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var archives []internal.ArchiveDisplay
	for rows.Next() {
		var archive internal.ArchiveDisplay

		if err := rows.Scan(
			&archive.ArchiveID,
			&archive.CustomerName,
			&archive.StartDate,
			&archive.EndDate,
			&archive.TotalPrice,
		); err != nil {
			return archives, err
		}

		archives = append(archives, archive)
	}

	if err = rows.Err(); err != nil {
		return archives, err
	}

	return archives, nil
}

func getCustomerActivityHandler(ctx *internal.AppContext) http.HandlerFunc {
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

		rentings, err := getCustomerRentings(ctx, customerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		archives, err := getCustomerArchives(ctx, customerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payload := struct {
			Bookings     []internal.BookingDisplay `json:"bookings"`
			Rentings     []internal.RentingDisplay `json:"rentings"`
			PastRentings []internal.ArchiveDisplay `json:"past_rentings"`
		}{
			Bookings:     bookings,
			Rentings:     rentings,
			PastRentings: archives,
		}

		json.NewEncoder(w).Encode(payload)
	}
}
