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
	HotelID  int    `json:"hotel_ID,omitempty"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	IDType   string `json:"ID_type,omitempty"`
	IDNumber string `json:"ID_number,omitempty"`
	Role     string `json:"role"`
}

type Room struct {
	ID         int       `json:"room_ID"`
	HotelID    int       `json:"hotel_ID"`
	RoomNumber string    `json:"room_number"`
	Capacity   int       `json:"capacity"`
	Price      float32   `json:"price"`
	ViewType   string    `json:"view_type"`
	Extendable bool      `json:"extendable"`
	Damaged    bool      `json:"damaged"`
	Amenities  []Amenity `json:"amenities"`
}

type Amenity struct {
	ID   int    `json:"amenity_ID"`
	Name string `json:"amenity_name"`
}

type Hotel struct {
	ID       int    `json:"hotel_ID"`
	Name     string `json:"hotel_name"`
	Address  string `json:"address"`
	Phone    string `json:"phone_number"`
	Email    string `json:"email"`
	Category int    `json:"category"`
}

type SearchResult struct {
	ID         int     `json:"room_ID"`
	HotelID    int     `json:"hotel_ID"`
	HotelName  string  `json:"hotel_name"`
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

type Booking struct {
	ID         int     `json:"booking_ID"`
	CustomerID int     `json:"customer_ID"`
	RoomID     int     `json:"room_ID"`
	StartDate  string  `json:"start_date"`
	EndDate    string  `json:"end_date"`
	TotalPrice float32 `json:"total_price"`
}

type Renting struct {
	ID           int     `json:"renting_ID"`
	EmployeeID   int     `json:"employee_ID"`
	CustomerID   int     `json:"customer_ID"`
	RoomID       int     `json:"room_ID"`
	BookingID    int     `json:"booking_ID,omitempty"`
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
	Payment      bool    `json:"payment"`
	TotalPrice   float32 `json:"total_price"`
}

type BookingDisplay struct {
	BookingID    int     `json:"booking_ID"`
	CustomerName string  `json:"customer_name"`
	HotelName    string  `json:"hotel_name"`
	RoomNumber   string  `json:"room_number"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	TotalPrice   float32 `json:"total_price"`
}

type RentingDisplay struct {
	RentingID    int     `json:"renting_ID"`
	EmployeeName string  `json:"employee_name"`
	CustomerName string  `json:"customer_name"`
	HotelName    string  `json:"hotel_name"`
	RoomNumber   string  `json:"room_number"`
	CheckInDate  string  `json:"check_in_date"`
	CheckOutDate string  `json:"check_out_date"`
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
