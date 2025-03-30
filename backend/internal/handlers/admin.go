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

		// Start transaction
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

		// Check if any rows were affected
		if rows, err := res.RowsAffected(); err != nil {
			http.Error(w, "Error checking update status", http.StatusInternalServerError)
			return
		} else if rows == 0 {
			http.Error(w, fmt.Sprintf("Hotel with ID %v not found", hotelID), http.StatusNotFound)
			return
		}

		// Commit transaction
		if err := tx.Commit(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(struct{ Message string }{Message: fmt.Sprintf("Successfully updated hotel with ID: %v", hotelID)})
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
			log.Printf("Error getting affected rows: %v", err.Error())
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
