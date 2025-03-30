package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

func updateCustomerByID(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerID := chi.URLParam(r, "customer_ID")

		var customer struct {
			Name    string `json:"full_name"`
			Address string `json:"address"`
		}

		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		result, err := ctx.DB.Exec(queries.UpdateCustomerByID,
			customer.Name,
			customer.Address,
			customerID,
		)

		if err != nil {
			http.Error(w, "Failed to update customer", http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, "Failed to get update result", http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Customer updated successfully",
		})
	}
}

func deleteCustomerByID(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		param := r.URL.Query().Get("customer_ID")
		if param == "" {
			http.Error(w, "customer_ID must be provided", http.StatusInternalServerError)
			return
		}

		customerID, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, fmt.Errorf("provided customer_ID '%v' is not a number", param).Error(), http.StatusBadRequest)
			return
		}

		res, err := ctx.DB.Exec(queries.DeleteCustomerByID, customerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rows, err := res.RowsAffected(); err != nil {
			log.Printf("Error getting affected rows: %v", err.Error())
			http.Error(w, "Error checking deletion status", http.StatusInternalServerError)
			return
		} else if rows == 0 {
			http.Error(w, fmt.Sprintf("customer with ID %v not found", customerID), http.StatusNotFound)
			return
		} else if rows > 1 {
			log.Printf("Multiple customers deleted from id: %v", customerID)
		}

		json.NewEncoder(w).Encode(struct {
			Message string
		}{
			Message: fmt.Sprintf("Successfully deleted customer with ID: %v", customerID),
		})
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

		managerID := r.URL.Query().Get("manager_ID")
		employeeName := r.URL.Query().Get("employee_name")

		if managerID == "" && employeeName == "" {
			http.Error(w, "Either manager_ID or employee_name must be provided", http.StatusInternalServerError)
			return
		}

		var query string
		var arg interface{}

		if employeeName != "" {
			query = queries.GetEmployeeByName
			arg = employeeName
		} else {
			query = queries.GetEmployeesByHotelID
			param, err := strconv.Atoi(managerID)
			if err != nil {
				http.Error(w, fmt.Errorf("provided manager_ID '%v' is not a number", managerID).Error(), http.StatusInternalServerError)
				return
			}

			hotelID, err := getHotelByEmployeeID(ctx, param)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			arg = hotelID
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

func updateEmployeeByID(ctx *internal.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employeeID := chi.URLParam(r, "employee_ID")

		var employee struct {
			Name    string `json:"full_name"`
			Address string `json:"address"`
		}

		if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		result, err := ctx.DB.Exec(queries.UpdateEmployeeByID,
			employee.Name,
			employee.Address,
			employeeID,
		)

		if err != nil {
			http.Error(w, "Failed to update employee", http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, "Failed to get update result", http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Employee updated successfully",
		})
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
