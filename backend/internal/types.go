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

type Booking struct {
	BookingID   int     `json:"booking_ID"`
	CustomerID  int     `json:"customer_ID"`
	RoomID      int     `json:"room_ID"`
	BookingDate string  `json:"booking_date"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	TotalPrice  float32 `json:"total_price"`
}

type Renting struct {
	RentingID    int     `json:"renting_ID"`
	EmployeeID   int     `json:"employee_ID"`
	CustomerID   int     `json:"customer_ID"`
	RoomID       int     `json:"room_ID"`
	BookingID    int     `json:"booking_ID"`
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
	Payment      bool    `json:"payment"`
	TotalPrice   float32 `json:"total_price"`
}

type Archive struct {
	ArchiveID    int     `json:"archive_ID"`
	RentingID    int     `json:"renting_ID"`
	BookingID    int     `json:"booking_ID"`
	CustomerID   int     `json:"customer_ID"`
	TotalPrice   float32 `json:"total_price"`
	BookingDate  string  `json:"booking_date"`
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
	ArchiveDate  string  `json:"archive_date"`
}
