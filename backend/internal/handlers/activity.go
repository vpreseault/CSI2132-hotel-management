package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/vpreseault/csi2132-project/backend/internal"
	"github.com/vpreseault/csi2132-project/backend/internal/queries"
)

func createBookingHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqBody struct {
			CustomerID *int    `json:"customer_ID"`
			RoomID     *int    `json:"room_ID"`
			StartDate  *string `json:"start_date"`
			EndDate    *string `json:"end_date"`
			TotalPrice *int    `json:"total_price"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if reqBody.CustomerID == nil ||
			reqBody.RoomID == nil ||
			reqBody.StartDate == nil ||
			*reqBody.StartDate == "" ||
			reqBody.EndDate == nil ||
			*reqBody.EndDate == "" ||
			reqBody.TotalPrice == nil {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		var newBooking = internal.Booking{
			CustomerID:  *reqBody.CustomerID,
			RoomID:      *reqBody.RoomID,
			BookingDate: time.Now().UTC().Format("2006-01-02"),
			StartDate:   *reqBody.StartDate,
			EndDate:     *reqBody.EndDate,
			TotalPrice:  float32(*reqBody.TotalPrice),
		}

		err := ctx.DB.QueryRow(
			queries.CreateBooking,
			newBooking.CustomerID,
			newBooking.RoomID,
			newBooking.BookingDate,
			newBooking.StartDate,
			newBooking.EndDate,
			newBooking.TotalPrice,
		).Scan(&newBooking.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(newBooking)
	}
}

func getBookingsHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		customerID := r.URL.Query().Get("customer_ID")
		hotelID := r.URL.Query().Get("hotel_ID")
		employeeID := r.URL.Query().Get("employee_ID")

		if customerID == "" && hotelID == "" && employeeID == "" {
			http.Error(w, "Either customer_ID, hotel_ID, employee_ID must be provided", http.StatusInternalServerError)
			return
		}

		var query string
		var param string

		if employeeID != "" {
			param = employeeID
		} else if customerID != "" {
			query = queries.GetBookingsByCustomerID
			param = customerID
		} else {
			query = queries.GetBookingsByHotelID
			param = hotelID
		}

		arg, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided ID '%v' is not a number", param).Error(), http.StatusInternalServerError)
			return
		}

		if employeeID != "" {
			hotelID, err := getHotelByEmployeeID(ctx, arg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			arg = hotelID
			query = queries.GetBookingsByHotelID
		}

		bookings, err := getBookings(ctx, query, arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(bookings)
	}
}

func getBookings(ctx *internal.AppContext, query string, id int) ([]internal.BookingDisplay, error) {
	rows, err := ctx.DB.Query(query, id)
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
			&booking.HotelName,
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

func getRentingsHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		customerID := r.URL.Query().Get("customer_ID")
		hotelID := r.URL.Query().Get("hotel_ID")
		employeeID := r.URL.Query().Get("employee_ID")

		if customerID == "" && hotelID == "" && employeeID == "" {
			http.Error(w, "Either customer_ID, hotel_ID, employee_ID must be provided", http.StatusInternalServerError)
			return
		}

		var query string
		var param string

		if employeeID != "" {
			param = employeeID
		} else if customerID != "" {
			query = queries.GetRentingsByCustomerID
			param = customerID
		} else {
			query = queries.GetRentingsByHotelID
			param = hotelID
		}

		arg, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided ID '%v' is not a number", param).Error(), http.StatusInternalServerError)
			return
		}

		if employeeID != "" {
			hotelID, err := getHotelByEmployeeID(ctx, arg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			arg = hotelID
			query = queries.GetRentingsByHotelID
		}

		rentings, err := getRentings(ctx, query, arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(rentings)
	}
}

func getRentings(ctx *internal.AppContext, query string, id int) ([]internal.RentingDisplay, error) {
	rows, err := ctx.DB.Query(query, id)
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
			&renting.HotelName,
			&renting.RoomNumber,
			&renting.CheckInDate,
			&renting.CheckOutDate,
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

func createRentingFromBookingHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqBody struct {
			BookingID  *int `json:"booking_ID"`
			EmployeeID *int `json:"employee_ID"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if reqBody.BookingID == nil || reqBody.EmployeeID == nil {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		var booking internal.Booking
		err := ctx.DB.QueryRow(queries.GetBookingByID, *reqBody.BookingID).Scan(
			&booking.ID,
			&booking.CustomerID,
			&booking.RoomID,
			&booking.StartDate,
			&booking.EndDate,
			&booking.TotalPrice,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var renting = internal.Renting{
			EmployeeID:   *reqBody.EmployeeID,
			CustomerID:   booking.CustomerID,
			RoomID:       booking.RoomID,
			BookingID:    &booking.ID,
			CheckInDate:  booking.StartDate,
			CheckOutDate: booking.EndDate,
			Payment:      true,
			TotalPrice:   booking.TotalPrice,
		}

		err = ctx.DB.QueryRow(
			queries.CreateRenting,
			renting.EmployeeID,
			renting.CustomerID,
			renting.RoomID,
			renting.BookingID,
			renting.CheckInDate,
			renting.CheckOutDate,
			renting.Payment,
			renting.TotalPrice,
		).Scan(&renting.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(renting)
	}
}

func createRentingHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqBody struct {
			CustomerID   *int     `json:"customer_ID"`
			EmployeeID   *int     `json:"employee_ID"`
			RoomID       *int     `json:"room_ID"`
			CheckInDate  *string  `json:"check_in_date"`
			CheckOutDate *string  `json:"check_out_date"`
			TotalPrice   *float32 `json:"total_price"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if reqBody.CustomerID == nil ||
			reqBody.EmployeeID == nil ||
			reqBody.RoomID == nil ||
			reqBody.CheckInDate == nil ||
			*reqBody.CheckInDate == "" ||
			reqBody.CheckOutDate == nil ||
			*reqBody.CheckOutDate == "" ||
			reqBody.TotalPrice == nil {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		var renting = internal.Renting{
			EmployeeID:   *reqBody.EmployeeID,
			CustomerID:   *reqBody.CustomerID,
			RoomID:       *reqBody.RoomID,
			BookingID:    nil,
			CheckInDate:  *reqBody.CheckInDate,
			CheckOutDate: *reqBody.CheckOutDate,
			Payment:      true,
			TotalPrice:   *reqBody.TotalPrice,
		}

		err := ctx.DB.QueryRow(
			queries.CreateRenting,
			renting.EmployeeID,
			renting.CustomerID,
			renting.RoomID,
			renting.BookingID,
			renting.CheckInDate,
			renting.CheckOutDate,
			renting.Payment,
			renting.TotalPrice,
		).Scan(&renting.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(renting)
	}
}

func getArchivesHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		customerID := r.URL.Query().Get("customer_ID")

		if customerID == "" {
			http.Error(w, "Param customer_ID must be provided", http.StatusInternalServerError)
			return
		}

		arg, err := strconv.Atoi(customerID)
		if err != nil {
			http.Error(w, fmt.Errorf("provided customer_ID '%v' is not a number", customerID).Error(), http.StatusInternalServerError)
			return
		}

		archives, err := getArchives(ctx, arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(archives)
	}
}

func getArchives(ctx *internal.AppContext, customerID int) ([]internal.ArchiveDisplay, error) {
	rows, err := ctx.DB.Query(queries.GetRentingArchivesByCustomerID, customerID)
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

func getActivityHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		customerID := r.URL.Query().Get("customer_ID")
		hotelID := r.URL.Query().Get("hotel_ID")

		if customerID == "" && hotelID == "" {
			http.Error(w, "Either customer_ID or hotel_ID must be provided", http.StatusInternalServerError)
			return
		}

		var queryList []string
		var param string

		if customerID != "" {
			queryList = []string{queries.GetBookingsByCustomerID, queries.GetRentingsByCustomerID, queries.GetRentingArchivesByCustomerID}
			param = customerID
		} else {
			queryList = []string{queries.GetBookingsByHotelID, queries.GetRentingsByHotelID}
			param = hotelID
		}

		arg, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided ID '%v' is not a number", param).Error(), http.StatusInternalServerError)
			return
		}

		bookings, err := getBookings(ctx, queryList[0], arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rentings, err := getRentings(ctx, queryList[1], arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var archives []internal.ArchiveDisplay
		if len(queryList) > 2 {
			archives, err = getArchives(ctx, arg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		payload := struct {
			Bookings []internal.BookingDisplay `json:"bookings"`
			Rentings []internal.RentingDisplay `json:"rentings"`
			Archives []internal.ArchiveDisplay `json:"archives"`
		}{
			Bookings: bookings,
			Rentings: rentings,
			Archives: archives,
		}

		json.NewEncoder(w).Encode(payload)
	}
}
