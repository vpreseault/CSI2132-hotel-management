package queries

var CreateCustomer = `
INSERT INTO Customers (
	full_name,
	ID_type, 
	ID_number,
	address,
	registration_date
) VALUES ($1, $2, $3, $4, $5)
RETURNING customer_ID
`

var GetCustomerByName = `SELECT * FROM Customers WHERE full_name = $1`

var GetCustomerIDByName = `SELECT customer_ID FROM Customers WHERE full_name = $1`

var UpdateCustomerByID = `UPDATE Customers 
SET full_name = $1, address = $2
WHERE customer_ID = $3
`

var DeleteCustomerByID = `DELETE FROM Customers WHERE customer_ID = $1`

var CreateEmployee = `
INSERT INTO Employees (
	hotel_ID,
	full_name, 
	address, 
	ID_type, 
	ID_number,
	role
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING employee_ID
`

var GetEmployeeByName = `SELECT
	employee_ID, 
	full_name, 
	role, 
	address
FROM Employees WHERE full_name = $1`

var UpdateEmployeeByID = `UPDATE Employees 
SET full_name = $1, address = $2
WHERE employee_ID = $3
`

// Get Employee List
var GetEmployeesByHotelID = `SELECT 
	employee_ID, 
	full_name, 
	role, 
	address
FROM Employees
WHERE hotel_ID = $1`

var DeleteEmployeeByID = `DELETE FROM Employees WHERE employee_ID = $1`
