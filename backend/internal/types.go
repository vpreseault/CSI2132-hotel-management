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
	HotelID  int    `json:"hotel_ID"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	IDType   string `json:"ID_type"`
	IDNumber string `json:"ID_number"`
	Role     string `json:"role"`
}

type Room struct {
	ID         int     `json:"room_ID"`
	HotelID    int     `json:"hotel_ID"`
	RoomNumber string  `json:"room_number"`
	Capacity   int     `json:"capacity"`
	Price      float32 `json:"price"`
	ViewType   string  `json:"view_type"`
	Extendable bool    `json:"extendable"`
	Damaged    bool    `json:"damaged"`
}

type SearchResult struct {
	ID         int     `json:"room_ID"`
	HotelID    int     `json:"hotel_ID"`
	RoomNumber string  `json:"room_number"`
	Capacity   int     `json:"capacity"`
	Price      float32 `json:"price"`
	ViewType   string  `json:"view_type"`
	Extendable bool    `json:"extendable"`
	Damaged    bool    `json:"damaged"`
	ChainName  string  `json:"chain_name"`
	Category   int     `json:"category"`
	Address    string  `json:"address"`
	TotalRooms int     `json:"total_rooms"`
}

type BookingDisplay struct {
	BookingID    int     `json:"booking_ID"`
	CustomerName string  `json:"customer_name"`
	RoomNumber   int     `json:"room_number"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	TotalPrice   float32 `json:"total_price"`
}

type RentingDisplay struct {
	RentingID    int     `json:"renting_ID"`
	EmployeeName string  `json:"employee_name"`
	CustomerName string  `json:"customer_name"`
	RoomNumber   int     `json:"room_number"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	Payment      bool    `json:"payment"`
	TotalPrice   float32 `json:"total_price"`
}

type ArchiveDisplay struct {
	ArchiveID    int     `json:"archive_ID"`
	CustomerName string  `json:"customer_name"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	TotalPrice   float32 `json:"total_price"`
}
