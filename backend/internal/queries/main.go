package queries

// Users
var CreateCustomer = `
INSERT INTO Customer (
	fullname, 
	address, 
	ID_type, 
	ID_number
	registration_date
) VALUES ($1, $2, $3, $4, $5)
`

var CreateEmployee = `
INSERT INTO Customer (
	fullname, 
	address, 
	ID_type, 
	ID_number
) VALUES ($1, $2, $3, $4)
`
