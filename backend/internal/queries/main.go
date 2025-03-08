package queries

// Users
var CreateCustomer = `
INSERT INTO Customer (
	full_name, 
	address,
	ID_type, 
	ID_number,
	registration_date
) VALUES ($1, $2, $3, $4, $5)
`

var GetCustomerByName = `SELECT * FROM customer WHERE full_name = $1`

var CreateEmployee = `
INSERT INTO Employee (
	hotel_ID,
	full_name, 
	address, 
	ID_type, 
	ID_number,
	role
) VALUES ($1, $2, $3, $4, $5, $6)
`

var GetEmployeeByName = `SELECT * FROM employee WHERE full_name = $1`
