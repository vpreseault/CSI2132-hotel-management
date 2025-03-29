export type FormError = {
    type: string;
    message: string;
}

export type SearchResult = {
    room_ID: number,
    hotel_ID: number,
    hotel_name: string,
    room_number: string,
    capacity: number,
    price: number,
    view_type: string,
    extendable: boolean,
    damaged: boolean,
    chain_name: string,
    category: number,
    address: string,
    total_rooms: number,
}