package internal

import "database/sql"

type AppContext struct {
	DB *sql.DB
}

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	IDType   string `json:"id_type"`
	IDNumber string `json:"id_number"`
}
