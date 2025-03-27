package queries

// Delete Chain/Hotel/Room
var DeleteChainByID = `DELETE FROM HotelChains WHERE chain_ID = $1`
var DeleteHotelByID = `DELETE FROM Hotels WHERE hotel_ID = $1`
var DeleteRoomByID = `DELETE FROM Rooms WHERE room_ID = $1`
