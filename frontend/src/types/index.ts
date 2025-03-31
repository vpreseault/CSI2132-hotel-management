export type FormError = {
    type: string;
    message: string;
}

export type Rental = {
    booking_ID?: number;
    check_in_date: string;
    check_out_date: string;
    customer_ID: number;
    customer_name: string;
    employee_ID: number;
    employee_name: string;
    hotel_name: string;
    payment: boolean;
    room_ID: number;
    room_number: string;
    total_price: number;
};

export type RentalItem = Pick<Rental,
    "customer_name" |
    "hotel_name" |
    "check_in_date" |
    "check_out_date" |
    "employee_name" |
    "room_number" |
    "total_price" |
    "payment"
>

export type RentalPayload = Pick<Rental,
    "check_in_date" |
    "check_out_date" |
    "customer_ID" |
    "employee_ID" |
    "room_ID" |
    "total_price"
>

export type RentalWithBookingPayload = Pick<Rental,
    "booking_ID" |
    "employee_ID" 
>

export type Booking = {
    booking_ID: number;
    customer_ID: number;
    customer_name: string;
    end_date: string;
    hotel_name: string;
    room_ID: number;
    room_number: string;
    start_date: string;
    total_price: number;
};

export type BookingItem = Pick<Booking,
    "booking_ID"|
    "customer_name"|
    "hotel_name"|
    "room_number"|
    "start_date"|
    "end_date"|
    "total_price"
>

export type BookingPayload = Pick<Booking,
    "customer_ID"|
    "room_ID"|
    "start_date"|
    "end_date"|
    "total_price"
>

export type ArchiveItem = {
    archive_ID: number,
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
    chain_ID: number,
    hotel_ID: number,
    hotel_name: string,
    manager_ID: number,
    address: string,
    phone_number: string,
    email: string,
    category: number,
}

export type HotelPayload = {
    chain_ID: number,
    hotel_name: string,
    manager_ID: number,
    address: string,
    phone_number: string,
    email: string,
    category: number,
}

export type HotelDisplay = {
    hotel_ID: number,
    hotel_name: string,
    address: string,
    phone_number: string,
    email: string,
    category: number,
}

export type Room = {
    room_ID?: number;
    room_number: string;
    capacity: number;
    price: number;
    view_type: string;
    extendable: boolean;
    damaged: boolean;
    amenities: Amenity[];
};

export type Amenity = {
    amenity_ID: number,
    amenity_name: string,
}
