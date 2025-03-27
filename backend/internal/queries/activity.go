package queries

var GetBookingsByCustomerID = `SELECT 
	b.booking_ID,
	c.full_name,
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
`

var GetBookingsByHotelID = `SELECT 
	b.booking_ID,
	c.full_name,
	r.room_number,
	b.start_date,
	b.end_date,
	b.total_price
FROM Bookings b
JOIN Rooms r ON r.room_ID = r.room_ID
JOIN Hotels h ON r.hotel_ID = h.hotel_ID
JOIN Customers c ON b.customer_ID = c.customer_ID
WHERE h.hotel_ID = $1
`

var GetRentingsByCustomerID = `SELECT 
	r.renting_ID,
	e.full_name,
	c.full_name,
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
`

var GetRentingsByHotelID = `SELECT 
	r.renting_ID,
	e.full_name as employee_name,
	c.full_name as customer_name,
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
	WHERE customer_ID = $1 AND renting_ID IS NOT NULL
) a
JOIN Customers c ON a.customer_ID = c.customer_ID
`
