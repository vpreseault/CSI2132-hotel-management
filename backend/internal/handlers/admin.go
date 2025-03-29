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

func deleteEmployeeByID(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		param := r.URL.Query().Get("employee_ID")
		if param == "" {
			http.Error(w, "employee_ID must be provided", http.StatusInternalServerError)
			return
		}

		employeeID, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided employee_ID '%v' is not a number", param).Error(), http.StatusBadRequest)
			return
		}

		res, err := ctx.DB.Exec(queries.DeleteEmployeeByID, employeeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rows, err := res.RowsAffected(); err != nil {
			log.Printf("Error getting affected rows: %v", err.Error())
			http.Error(w, "Error checking deletion status", http.StatusInternalServerError)
			return
		} else if rows == 0 {
			http.Error(w, fmt.Sprintf("Employee with ID %v not found", employeeID), http.StatusNotFound)
			return
		} else if rows > 1 {
			log.Printf("Multiple employees deleted from id: %v", employeeID)
		}

		json.NewEncoder(w).Encode(struct {
			Message string
		}{
			Message: fmt.Sprintf("Successfully deleted employee with ID: %v", employeeID),
		})
	}
}
