package queries

// Users
var CreateUser = `
INSERT INTO users (
	fullname, 
	address, 
	id_type, 
	id_number
) VALUES ($1, $2, $3, $4)
`
