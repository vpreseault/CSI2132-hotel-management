package internal

import "database/sql"

type AppContext struct {
	DB *sql.DB
}

type Customer struct {
	ID               int    `json:"customer_ID"`
	FullName         string `json:"full_name"`
	IDType           string `json:"ID_type"`
	IDNumber         string `json:"ID_number"`
	Address          string `json:"address"`
	RegistrationDate string `json:"registration_date"`
}

type Employee struct {
	ID       int    `json:"employee_ID"`
	HotelID  string `json:"hotel_ID"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	IDType   string `json:"ID_type"`
	IDNumber string `json:"ID_number"`
	Role     string `json:"role"`
}
