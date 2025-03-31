package queries

var CreateBooking = `INSERT INTO Bookings (
	customer_ID, 
	room_ID, 
	booking_date,
	start_date,
	end_date,
	total_price
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING booking_ID
`

var GetBookingsByCustomerID = `SELECT 
	b.booking_ID,
	c.full_name,
	h.hotel_name,
	r.room_number,
	b.start_date,
	b.end_date,
	b.total_price
FROM (
	SELECT * FROM Bookings
	WHERE customer_ID = $1
) b
JOIN Customers c ON b.customer_ID = c.customer_ID
JOIN Rooms r ON b.room_ID = r.room_ID
JOIN Hotels h ON r.hotel_ID = h.hotel_ID
`

var GetBookingsByHotelID = `SELECT 
	b.booking_ID,
	c.full_name,
	h.hotel_name,
	r.room_number,
	b.start_date,
	b.end_date,
	b.total_price
FROM Bookings b
JOIN Rooms r ON b.room_ID = r.room_ID
JOIN Hotels h ON r.hotel_ID = h.hotel_ID
JOIN Customers c ON b.customer_ID = c.customer_ID
WHERE h.hotel_ID = $1
`

var GetBookingByID = `SELECT 
	b.booking_ID,
	c.customer_ID,
	r.room_ID,
	b.start_date,
	b.end_date,
	b.total_price
FROM (
	SELECT * FROM Bookings
	WHERE booking_ID = $1
) b
JOIN Customers c ON b.customer_ID = c.customer_ID
JOIN Rooms r ON b.room_ID = r.room_ID
`

var CreateRenting = `INSERT INTO Rentings (
	employee_ID,
	customer_ID, 
	room_ID, 
	booking_ID, 
	check_in_date,
	check_out_date,
	payment,
	total_price
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING renting_ID
`

var GetRentingsByCustomerID = `SELECT 
	r.renting_ID,
	e.full_name,
	c.full_name,
	h.hotel_name,
	rm.room_number,
	r.check_in_date,
	r.check_out_date,
	r.payment,
	r.total_price
FROM (
	SELECT * FROM Rentings
	WHERE customer_ID = $1
) r
JOIN Customers c ON r.customer_ID = c.customer_ID
JOIN Employees e ON r.employee_ID = e.employee_ID
JOIN Rooms rm ON r.room_ID = rm.room_ID
JOIN Hotels h ON rm.hotel_ID = h.hotel_ID
`

var GetRentingsByHotelID = `SELECT 
	r.renting_ID,
	e.full_name as employee_name,
	c.full_name as customer_name,
	h.hotel_name,
	rm.room_number,
	r.check_in_date,
	r.check_out_date,
	r.payment,
	r.total_price
FROM Rentings r
JOIN Rooms rm ON r.room_ID = rm.room_ID
JOIN Hotels h ON rm.hotel_ID = h.hotel_ID
JOIN Customers c ON r.customer_ID = c.customer_ID
JOIN Employees e ON r.employee_ID = e.employee_ID
WHERE h.hotel_ID = $1
`

var GetRentingArchivesByCustomerID = `SELECT 
	a.archive_ID,
	c.full_name,
	a.check_in_date,
	a.check_out_date,
	a.total_price
FROM (
	SELECT * FROM Archives
	WHERE customer_ID = $1
) a
JOIN Customers c ON a.customer_ID = c.customer_ID
`
