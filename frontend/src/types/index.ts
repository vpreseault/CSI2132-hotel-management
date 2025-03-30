export type FormError = {
    type: string;
    message: string;
}

export type RentalItem = {
    customer_name: string;
    hotel_name: string;
    start_date: string;
    end_date: string;
    employee_name: string;
    room_number: number;
    total_price: number;
    payment: boolean;
};

export type BookingItem = {
    customer_name: string;
    hotel_name: string;
    room_number: number;
    start_date: string;
    end_date: string;
    total_price: number;
};

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

export type Hotel = {
    hotel_ID: number,
    hotel_name: string,
    address: string,
    phone_number: string,
    email: string,
    category: number,
}