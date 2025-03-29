package queries

// Search
var BaseRoomSearch = `SELECT 
	room_ID,
	hotel_ID,
	hotel_name,
	room_number,
	capacity,
	price,
	view_type,
	extendable,
	damaged,
	chain_name,
	category,
	address,
	total_rooms
FROM RoomSearchView`

var GetEmployeeHotelID = `SELECT hotel_ID FROM Employees WHERE employee_ID = $1`
