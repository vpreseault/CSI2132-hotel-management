package internal

import "database/sql"

type AppContext struct {
	DB *sql.DB
}

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
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	IDType   string `json:"ID_type"`
	IDNumber string `json:"ID_number"`
}
