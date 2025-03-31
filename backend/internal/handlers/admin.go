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

type hotelChain struct {
	ID   int    `json:"chain_ID"`
	Name string `json:"chain_name"`
}

func getChainsHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		rows, err := ctx.DB.Query(queries.GetChains)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var chains []hotelChain
		for rows.Next() {
			var chain hotelChain

			if err := rows.Scan(
				&chain.ID,
				&chain.Name,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			chains = append(chains, chain)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(chains)
	}
}

func deleteChainByID(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		param := chi.URLParam(r, "chain_ID")
		chainID, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided chain_ID '%v' is not a number", param).Error(), http.StatusBadRequest)
			return
		}

		res, err := ctx.DB.Exec(queries.DeleteChainByID, chainID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if rows, err := res.RowsAffected(); err != nil {
			log.Printf("Error getting affected rows: %v", err.Error())
			http.Error(w, "Error checking deletion status", http.StatusInternalServerError)
			return
		} else if rows == 0 {
			http.Error(w, fmt.Sprintf("Chain with ID %v not found", chainID), http.StatusNotFound)
			return
		} else if rows > 1 {
			log.Printf("Multiple chains deleted from id: %v", chainID)
		}

		json.NewEncoder(w).Encode(struct{ Message string }{Message: fmt.Sprintf("Successfully deleted chain with ID: %v", chainID)})
	}
}

func createHotelHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var payload struct {
			internal.Hotel
			ChainID   int `json:"chain_ID"`
			ManagerID int `json:"manager_ID"`
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tx, err := ctx.DB.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()

		err = tx.QueryRow(
			queries.InsertHotel,
			payload.ChainID,
			nil,
			payload.Name,
			payload.Address,
			payload.Category,
		).Scan(&payload.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = tx.Exec(queries.InsertHotelPhone, payload.ID, payload.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = tx.Exec(queries.InsertHotelEmail, payload.ID, payload.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tx.Commit(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(payload)
	}
}

func getHotelsHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		managerID := r.URL.Query().Get("manager_ID")

		if managerID == "" {
			http.Error(w, "manager_ID must be provided", http.StatusInternalServerError)
			return
		}

		param, err := strconv.Atoi(managerID)
		if err != nil {
			http.Error(w, fmt.Errorf("provided manager_ID '%v' is not a number", managerID).Error(), http.StatusInternalServerError)
			return
		}

		var hotel internal.Hotel
		err = ctx.DB.QueryRow(
			queries.GetHotelByManagerID,
			param,
		).Scan(
			&hotel.ID,
			&hotel.Name,
			&hotel.Address,
			&hotel.Phone,
			&hotel.Email,
			&hotel.Category,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(hotel)
	}
}

func getAllHotelsHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		rows, err := ctx.DB.Query(queries.GetHotels)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var hotels []internal.Hotel
		for rows.Next() {
			var hotel internal.Hotel
			if err := rows.Scan(
				&hotel.ID,
				&hotel.Name,
				&hotel.Address,
				&hotel.Phone,
				&hotel.Email,
				&hotel.Category,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			hotels = append(hotels, hotel)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(hotels)
	}
}

func deleteHotelByID(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		param := chi.URLParam(r, "hotel_ID")
		hotelID, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided hotel_ID '%v' is not a number", param).Error(), http.StatusBadRequest)
			return
		}

		res, err := ctx.DB.Exec(queries.DeleteHotelByID, hotelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if rows, err := res.RowsAffected(); err != nil {
			log.Printf("Error getting affected rows: %v", err.Error())
			http.Error(w, "Error checking deletion status", http.StatusInternalServerError)
			return
		} else if rows == 0 {
			http.Error(w, fmt.Sprintf("Hotel with ID %v not found", hotelID), http.StatusNotFound)
			return
		} else if rows > 1 {
			log.Printf("Multiple hotels deleted from id: %v", hotelID)
		}

		json.NewEncoder(w).Encode(struct{ Message string }{Message: fmt.Sprintf("Successfully deleted hotel with ID: %v", hotelID)})
	}
}

func updateHotel(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		param := chi.URLParam(r, "hotel_ID")
		hotelID, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided hotel_ID '%v' is not a number", param).Error(), http.StatusBadRequest)
			return
		}

		var hotel internal.Hotel
		if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if hotel.Name == "" || hotel.Address == "" || hotel.Phone == "" || hotel.Email == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		if hotel.Category < 1 || hotel.Category > 5 {
			http.Error(w, fmt.Sprintf("Invalid category value %v", hotel.Category), http.StatusBadRequest)
			return
		}

		tx, err := ctx.DB.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()

		res, err := tx.Exec(queries.UpdateHotel, hotel.Name, hotel.Address, hotel.Category, hotelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rows, err := res.RowsAffected(); err != nil {
			http.Error(w, "Error checking update status", http.StatusInternalServerError)
			return
		} else if rows == 0 {
			http.Error(w, fmt.Sprintf("Hotel with ID %v not found", hotelID), http.StatusNotFound)
			return
		}

		res, err = tx.Exec(queries.UpdateHotelPhone, hotel.Phone, hotelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rows, err := res.RowsAffected(); err != nil {
			http.Error(w, "Error checking update status", http.StatusInternalServerError)
			return
		} else if rows == 0 {
			http.Error(w, fmt.Sprintf("Hotel with ID %v not found", hotelID), http.StatusNotFound)
			return
		}

		res, err = tx.Exec(queries.UpdateHotelEmail, hotel.Email, hotelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rows, err := res.RowsAffected(); err != nil {
			http.Error(w, "Error checking update status", http.StatusInternalServerError)
			return
		} else if rows == 0 {
			http.Error(w, fmt.Sprintf("Hotel with ID %v not found", hotelID), http.StatusNotFound)
			return
		}

		if err := tx.Commit(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(struct{ Message string }{Message: fmt.Sprintf("Successfully updated hotel with ID: %v", hotelID)})
	}
}

func getViewsDataHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		roomsPerAreaRows, err := ctx.DB.Query(queries.GetRoomsPerAreaView)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer roomsPerAreaRows.Close()

		hotelCapacityRows, err := ctx.DB.Query(queries.GetHotelCapacityView)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer hotelCapacityRows.Close()

		var roomsPerArea []struct {
			Area       string `json:"area"`
			TotalRooms int    `json:"total_rooms"`
		}

		for roomsPerAreaRows.Next() {
			var data struct {
				Area       string `json:"area"`
				TotalRooms int    `json:"total_rooms"`
			}
			if err := roomsPerAreaRows.Scan(&data.Area, &data.TotalRooms); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			roomsPerArea = append(roomsPerArea, data)
		}

		var hotelCapacity []struct {
			HotelName     string `json:"hotel_name"`
			TotalCapacity int    `json:"total_capacity"`
		}

		for hotelCapacityRows.Next() {
			var data struct {
				HotelName     string `json:"hotel_name"`
				TotalCapacity int    `json:"total_capacity"`
			}
			if err := hotelCapacityRows.Scan(&data.HotelName, &data.TotalCapacity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			hotelCapacity = append(hotelCapacity, data)
		}

		if err = roomsPerAreaRows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = hotelCapacityRows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(struct {
			RoomsPerArea []struct {
				Area       string `json:"area"`
				TotalRooms int    `json:"total_rooms"`
			} `json:"rooms_per_area"`
			HotelCapacity []struct {
				HotelName     string `json:"hotel_name"`
				TotalCapacity int    `json:"total_capacity"`
			} `json:"hotel_capacity"`
		}{
			RoomsPerArea:  roomsPerArea,
			HotelCapacity: hotelCapacity,
		})
	}
}
