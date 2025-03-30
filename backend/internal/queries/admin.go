package queries

// Chain
var GetChains = `SELECT chain_ID, chain_name FROM HotelChains`
var DeleteChainByID = `DELETE FROM HotelChains WHERE chain_ID = $1`

// Hotel
var GetHotels = `SELECT * FROM Hotels`
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

// var GetRooms = `SELECT * FROM Rooms WHERE room_ID = $1`
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
var DeleteRoomByID = `DELETE FROM Rooms WHERE room_ID = $1`
