package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vpreseault/csi2132-project/backend/internal"
	"github.com/vpreseault/csi2132-project/backend/internal/queries"
)

func createRoomHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var payload struct {
			EmployeeID int `json:"employee_ID"`
			internal.Room
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload.HotelID, err = getHotelByEmployeeID(ctx, payload.EmployeeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tx, err := ctx.DB.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()

		err = tx.QueryRow(
			queries.CreateRoom,
			payload.HotelID,
			payload.RoomNumber,
			payload.Capacity,
			payload.Price,
			payload.ViewType,
			payload.Extendable,
			payload.Damaged,
		).Scan(&payload.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Add new amenities
		for _, v := range payload.Amenities {
			_, err = tx.Exec(queries.InsertRoomAmenity, payload.ID, v.ID)
			if err != nil {
				http.Error(w, fmt.Errorf("could not insert new amenity (%v): %v", v.Name, err.Error()).Error(), http.StatusInternalServerError)
				return
			}
		}

		if err := tx.Commit(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(struct {
			Message string
		}{
			Message: fmt.Sprintf("Successfully created room with ID: %v", payload.ID),
		})
	}
}

func getHotelRoomsHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		param := r.URL.Query().Get("employee_ID")
		if param == "" {
			http.Error(w, "employee_ID must be provided", http.StatusInternalServerError)
			return
		}

		employeeID, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided employee_ID '%v' is not a number", employeeID).Error(), http.StatusInternalServerError)
			return
		}

		hotelID, err := getHotelByEmployeeID(ctx, employeeID)
		if err != nil {
			http.Error(w, fmt.Errorf("could not get employee's hotel: %v", err.Error()).Error(), http.StatusInternalServerError)
			return
		}

		rows, err := ctx.DB.Query(queries.GetRooms, hotelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var rooms []internal.Room
		for rows.Next() {
			var room internal.Room
			var amenitiesJSON []byte
			if err := rows.Scan(
				&room.ID,
				&room.HotelID,
				&room.RoomNumber,
				&room.Capacity,
				&room.Price,
				&room.ViewType,
				&room.Extendable,
				&room.Damaged,
				&amenitiesJSON,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := json.Unmarshal(amenitiesJSON, &room.Amenities); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			rooms = append(rooms, room)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(rooms)
	}
}

func updateRoom(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		param := chi.URLParam(r, "room_ID")
		roomID, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided room_ID '%v' is not a number", param).Error(), http.StatusBadRequest)
			return
		}

		var room internal.Room
		if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if room.RoomNumber == "" || room.ViewType == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		if room.Capacity < 1 || room.Capacity > 10 {
			http.Error(w, fmt.Sprintf("Invalid Capacity value %v", room.Capacity), http.StatusBadRequest)
			return
		}

		// Start transaction
		tx, err := ctx.DB.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()

		res, err := tx.Exec(
			queries.UpdateRoom,
			roomID,
			room.RoomNumber,
			room.Capacity,
			room.Price,
			room.ViewType,
			room.Extendable,
			room.Damaged,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rows, err := res.RowsAffected(); err != nil {
			http.Error(w, "Error checking update status", http.StatusInternalServerError)
			return
		} else if rows == 0 {
			http.Error(w, fmt.Sprintf("Room with ID %v not found", roomID), http.StatusNotFound)
			return
		}

		// Delete existing amenities
		_, err = tx.Exec(queries.DeleteRoomAmenities, roomID)
		if err != nil {
			http.Error(w, fmt.Errorf("could not delete old amenities: %v", err.Error()).Error(), http.StatusInternalServerError)
			return
		}

		// Add new amenities
		for _, v := range room.Amenities {
			_, err = tx.Exec(queries.InsertRoomAmenity, roomID, v.ID)
			if err != nil {
				http.Error(w, fmt.Errorf("could not insert new amenity (%v): %v", v.Name, err.Error()).Error(), http.StatusInternalServerError)
				return
			}
		}

		if err := tx.Commit(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(struct{ Message string }{Message: fmt.Sprintf("Successfully updated room with ID: %v", roomID)})
	}
}

func deleteRoomByID(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		param := chi.URLParam(r, "room_ID")
		roomID, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided room_ID '%v' is not a number", param).Error(), http.StatusBadRequest)
			return
		}

		res, err := ctx.DB.Exec(queries.DeleteRoomByID, roomID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if rows, err := res.RowsAffected(); err != nil {
			http.Error(w, "Error checking deletion status", http.StatusInternalServerError)
			return
		} else if rows == 0 {
			http.Error(w, fmt.Sprintf("Room with ID %v not found", roomID), http.StatusNotFound)
			return
		} else if rows > 1 {
			log.Printf("Multiple rooms deleted from id: %v", roomID)
		}

		json.NewEncoder(w).Encode(struct{ Message string }{Message: fmt.Sprintf("Successfully deleted room with ID: %v", roomID)})
	}
}

func getAmenitiesHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		rows, err := ctx.DB.Query(queries.GetAmenities)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var amenities []internal.Amenity
		for rows.Next() {
			var amenity internal.Amenity
			if err := rows.Scan(
				&amenity.ID,
				&amenity.Name,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			amenities = append(amenities, amenity)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(amenities)
	}
}
