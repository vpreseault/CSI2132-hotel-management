package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/vpreseault/csi2132-project/backend/internal"
	"github.com/vpreseault/csi2132-project/backend/internal/queries"
)

func createCustomerHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		customer := internal.Customer{
			RegistrationDate: time.Now().UTC().Format("2006-01-02"),
		}
		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = ctx.DB.QueryRow(queries.CreateCustomer, customer.FullName, customer.IDType, customer.IDNumber, customer.Address, customer.RegistrationDate).Scan(&customer.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(customer)
	}
}

func getCustomerHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userName := chi.URLParam(r, "name")
		var customer internal.Customer

		err := ctx.DB.QueryRow(queries.GetCustomerByName, userName).Scan(&customer.ID, &customer.FullName, &customer.IDType, &customer.IDNumber, &customer.Address, &customer.RegistrationDate)
		if err != nil {
			if err == sql.ErrNoRows {
				// Treat 0 rows error differently
				// http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(customer)
	}
}

func createEmployeeHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var payload struct {
			ManagerID int `json:"manager_ID"`
			internal.Employee
		}

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var HotelID int
		err = ctx.DB.QueryRow(queries.GetEmployeeHotelID, payload.ManagerID).Scan(&HotelID)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not get HotelID from ManagerID. %v", err.Error()), http.StatusInternalServerError)
			return
		}

		employee := internal.Employee{
			FullName: payload.FullName,
			Role:     payload.Role,
		}
		err = ctx.DB.QueryRow(queries.CreateEmployee, HotelID, payload.FullName, payload.Address, payload.IDType, payload.IDNumber, payload.Role).Scan(&payload.ID)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not create new employee. %v", err.Error()), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(employee)
	}
}

func getEmployeeHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userName := chi.URLParam(r, "name")
		var employee internal.Employee

		err := ctx.DB.QueryRow(queries.GetEmployeeByName, userName).Scan(&employee.ID, &employee.HotelID, &employee.FullName, &employee.Address, &employee.IDType, &employee.IDNumber, &employee.Role)
		if err != nil {
			if err == sql.ErrNoRows {
				// Treat 0 rows error differently
				// http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(employee)
	}
}
