package queries

// Auth
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

var GetEmployeeByName = `SELECT * FROM Employees WHERE full_name = $1`

// Activity
var GetCustomerBookings = `SELECT * FROM Bookings WHERE customer_id = $1`
var GetCustomerRentings = `SELECT * FROM Rentings WHERE customer_id = $1`
var GetCustomerArchives = `SELECT * FROM Archives WHERE customer_id = $1`
