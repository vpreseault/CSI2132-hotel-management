package internal

import "database/sql"

type AppContext struct {
	DB *sql.DB
}

// TODO: cleanup user structs
type Customer struct {
	ID               int    `json:"ID"`
	Fullname         string `json:"fullname"`
	Address          string `json:"address"`
	IDType           string `json:"ID_type"`
	IDNumber         string `json:"ID_number"`
	RegistrationDate string `json:"registration_date"`
}

type Employee struct {
	ID       int    `json:"ID"`
	HotelID  string `json:"ID_number"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	IDType   string `json:"ID_type"`
	IDNumber string `json:"ID_number"`
	Role     string `json:"role"`
}
