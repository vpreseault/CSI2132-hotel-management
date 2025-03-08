package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/vpreseault/csi2132-project/backend/internal"
	"github.com/vpreseault/csi2132-project/backend/internal/queries"
)

func createCustomerHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var customer internal.Customer
		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		registrationDate := time.Now().UTC().Format("2006-01-02")

		err = ctx.DB.QueryRow(queries.CreateCustomer, customer.Fullname, customer.Address, customer.IDType, customer.IDNumber, registrationDate).Scan(&customer.ID)
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

		err := ctx.DB.QueryRow(queries.GetCustomerByName, userName).Scan(&customer.ID, &customer.Fullname, &customer.IDType, &customer.IDNumber, &customer.RegistrationDate)
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

		var employee internal.Employee
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = ctx.DB.QueryRow(queries.CreateEmployee, employee.Fullname, employee.Address, employee.IDType, employee.IDNumber).Scan(&employee.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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

		err := ctx.DB.QueryRow(queries.GetEmployeeByName, userName).Scan(&employee.ID, &employee.HotelID, &employee.Fullname, &employee.Address, &employee.IDNumber, &employee.Role)
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
