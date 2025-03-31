package queries

// Chain
var GetChains = `SELECT chain_ID, chain_name FROM HotelChains`
var DeleteChainByID = `DELETE FROM HotelChains WHERE chain_ID = $1`

// Hotel
var GetHotels = `SELECT 
	h.hotel_ID,
	h.hotel_name,
	h.address,
	hp.h_phone,
	he.h_email,
	h.category
FROM Hotels h
JOIN HotelPhones hp ON h.hotel_ID = hp.hotel_ID
JOIN HotelEmails he ON h.hotel_ID = he.hotel_ID
ORDER BY h.hotel_ID DESC
`
var GetHotelByManagerID = `SELECT 
	h.hotel_ID,
	h.hotel_name,
	h.address,
	hp.h_phone,
	he.h_email,
	h.category
FROM Hotels h
JOIN HotelPhones hp ON h.hotel_ID = hp.hotel_ID
JOIN HotelEmails he ON h.hotel_ID = he.hotel_ID
WHERE h.manager_ID = $1
`
var DeleteHotelByID = `DELETE FROM Hotels WHERE hotel_ID = $1`
var InsertHotel = `INSERT INTO Hotels (
	chain_ID,
	manager_ID,
	hotel_name,
	address,
	category
) VALUES ($1, $2, $3, $4, $5)
RETURNING hotel_ID
`
var InsertHotelPhone = `INSERT INTO HotelPhones (
	hotel_ID,
	h_phone
) VALUES ($1, $2)
`
var InsertHotelEmail = `INSERT INTO HotelEmails (
	hotel_ID,
	h_email
) VALUES ($1, $2)
`

var UpdateHotel = `UPDATE Hotels 
SET hotel_name = $1, 
    address = $2,
    category = $3
WHERE hotel_ID = $4`

var UpdateHotelPhone = `UPDATE HotelPhones 
SET h_phone = $1
WHERE hotel_ID = $2`

var UpdateHotelEmail = `UPDATE HotelEmails 
SET h_email = $1
WHERE hotel_ID = $2`

// Room
var CreateRoom = `
INSERT INTO Rooms (
	hotel_ID,
	room_number, 
	capacity,
	price,
	view_type,
	extendable,
	damaged
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING room_ID
`
var GetRooms = `
SELECT 
	r.room_ID,
	r.hotel_ID,
	r.room_number,
	r.capacity,
	r.price,
	r.view_type,
	r.extendable,
	r.damaged,
	COALESCE(
		json_agg(
			json_build_object(
				'amenity_ID', a.amenity_ID,
				'amenity_name', a.amenity_name
			)
		) FILTER (WHERE a.amenity_ID IS NOT NULL),
		'[]'::json
	) as amenities
FROM Rooms r
LEFT JOIN Room_Has_Amenities ra ON r.room_ID = ra.room_ID
LEFT JOIN Amenities a ON ra.amenity_ID = a.amenity_ID
WHERE r.hotel_ID = $1
GROUP BY r.room_ID
ORDER BY r.room_number
`
var DeleteRoomByID = `DELETE FROM Rooms WHERE room_ID = $1`
var UpdateRoom = `UPDATE Rooms 
SET 
    room_number = $2, 
    capacity = $3,
    price = $4,
    view_type = $5,
    extendable = $6,
    damaged = $7
WHERE room_ID = $1`

// Amenities
var GetRoomAmenities = `SELECT 
	a.amenity_ID,
	a.amenity_name
FROM (
	SELECT * FROM Rooms
	WHERE room_ID = $1
) r
JOIN RoomHasAmenities ra ON r.room_ID = ra.room_ID
JOIN Amenities a ON ra.amenity_ID = a.amenity_ID
`
var GetAmenities = `SELECT * FROM Amenities`
var DeleteRoomAmenities = `DELETE FROM Room_Has_Amenities WHERE room_ID = $1`
var InsertRoomAmenity = `INSERT INTO Room_Has_Amenities (
	room_ID,
	amenity_ID
) VALUES ($1, $2)
`

// Select from views
var GetRoomsPerAreaView = `SELECT * FROM RoomsPerAreaView`
var GetHotelCapacityView = `SELECT * FROM HotelCapacityView`
