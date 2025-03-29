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