package queries

// Get Chain/Hotel/Room
var GetChains = `SELECT chain_ID, chain_name FROM HotelChains`
var GetHotels = `SELECT * FROM Hotels`
var GetRooms = `SELECT * FROM Rooms`

// Delete Chain/Hotel/Room
var DeleteChainByID = `DELETE FROM HotelChains WHERE chain_ID = $1`
var DeleteHotelByID = `DELETE FROM Hotels WHERE hotel_ID = $1`
var DeleteRoomByID = `DELETE FROM Rooms WHERE room_ID = $1`

// Get Employee List
var GetEmployeesByHotelID = `SELECT 
	employee_ID, 
	full_name, 
	role, 
	address
FROM Employees
WHERE hotel_ID = $1`

var DeleteEmployeeByID = `DELETE FROM Employees WHERE employee_ID = $1`
