package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

		HotelID, err := getHotelByEmployeeID(ctx, payload.ManagerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
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

func getEmployeesHandler(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		hotelID := r.URL.Query().Get("hotel_ID")
		employeeName := r.URL.Query().Get("employee_name")

		if hotelID == "" && employeeName == "" {
			http.Error(w, "Either hotel_ID, employee_name must be provided", http.StatusInternalServerError)
			return
		}

		var query string
		var arg interface{}

		if employeeName != "" {
			query = queries.GetEmployeeByName
			arg = employeeName
		} else {
			query = queries.GetEmployeesByHotelID
			param, err := strconv.Atoi(hotelID)
			if err != nil {
				http.Error(w, fmt.Errorf("provided hotel_ID '%v' is not a number", hotelID).Error(), http.StatusInternalServerError)
				return
			}
			arg = param
		}

		rows, err := ctx.DB.Query(query, arg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var employees []internal.Employee
		for rows.Next() {
			var employee internal.Employee
			if err := rows.Scan(
				&employee.ID,
				&employee.FullName,
				&employee.Role,
				&employee.Address,
			); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			employees = append(employees, employee)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if employeeName != "" {
			if len(employees) > 1 {
				http.Error(w, fmt.Errorf("more than one employee with name: %v", employeeName).Error(), http.StatusInternalServerError)
			} else if len(employees) == 1 {
				json.NewEncoder(w).Encode(employees[0])
			} else {
				json.NewEncoder(w).Encode(nil)
			}
			return
		}

		json.NewEncoder(w).Encode(employees)
	}
}
