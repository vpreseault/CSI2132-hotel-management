DROP TABLE IF EXISTS Archive CASCADE;
DROP TABLE IF EXISTS Renting CASCADE;
DROP TABLE IF EXISTS Booking CASCADE;
DROP TABLE IF EXISTS Room_Has_Amenities CASCADE;
DROP TABLE IF EXISTS Amenities CASCADE;
DROP TABLE IF EXISTS Room CASCADE;
DROP TABLE IF EXISTS Employee CASCADE;
DROP TABLE IF EXISTS Hotel CASCADE;
DROP TABLE IF EXISTS HotelChain CASCADE;
DROP TABLE IF EXISTS Customer CASCADE;

/* Create tables without connecting the foreign keys */

-- HotelChain Table
CREATE TABLE HotelChain (
    chain_ID SERIAL PRIMARY KEY,
    central_office_address VARCHAR(255) NOT NULL, 
    number_of_hotels INT NOT NULL CHECK (number_of_hotels > 0)
);

-- Insert into HotelChain
INSERT INTO HotelChain (central_office_address, number_of_hotels) VALUES
('123 Main St, New York, NY', 10),
('456 Elm St, Los Angeles, CA', 7),
('789 Oak St, Chicago, IL', 5),
('101 Pine St, Houston, TX', 8),
('202 Maple St, Miami, FL', 6),
('303 Birch St, Seattle, WA', 4),
('404 Cedar St, Denver, CO', 3),
('505 Walnut St, Boston, MA', 9),
('606 Cherry St, Atlanta, GA', 5),
('707 Aspen St, San Francisco, CA', 7);

-- Hotel Table
CREATE TABLE Hotel (
    hotel_ID SERIAL PRIMARY KEY, 
    chain_ID INT NOT NULL,
    manager_ID INT UNIQUE,  
    number_of_rooms INT NOT NULL CHECK (number_of_rooms > 0),  
    hotel_address VARCHAR(255) NOT NULL,
    category INT NOT NULL CHECK (category BETWEEN 1 AND 5)
);

-- Insert into Hotel
INSERT INTO Hotel (chain_ID, manager_ID, number_of_rooms, hotel_address, category) VALUES
(1, NULL, 100, '1001 Broadway, New York, NY', 5),
(2, NULL, 80, '2022 Sunset Blvd, Los Angeles, CA', 4),
(3, NULL, 50, '3033 Lakeshore Dr, Chicago, IL', 3),
(4, NULL, 120, '4040 Westheimer Rd, Houston, TX', 5),
(5, NULL, 90, '5055 Ocean Dr, Miami, FL', 4),
(6, NULL, 70, '6066 Rainier Ave, Seattle, WA', 3),
(7, NULL, 60, '7077 Colfax Ave, Denver, CO', 3),
(8, NULL, 130, '8088 Beacon St, Boston, MA', 5),
(9, NULL, 110, '9099 Peachtree St, Atlanta, GA', 4),
(10, NULL, 95, '1010 Market St, San Francisco, CA', 4);

-- Phone and Emails for hotel and hotelchain
CREATE TABLE ChainPhone (
    chain_ID INT NOT NULL,
    c_phone VARCHAR(20) UNIQUE NOT NULL CHECK (c_phone ~ '^[0-9-]+$'),
    PRIMARY KEY (chain_ID, c_phone)
);

-- Insert into ChainPhone
INSERT INTO ChainPhone (chain_ID, c_phone) VALUES
(1, '212-555-1234'),
(2, '310-555-5678'),
(3, '312-555-9012'),
(4, '713-555-3456'),
(5, '305-555-7890'),
(6, '206-555-6543'),
(7, '303-555-2222'),
(8, '617-555-7777'),
(9, '404-555-8888'),
(10, '415-555-9999');

CREATE TABLE ChainEmail (
    chain_ID INT NOT NULL,
    c_email VARCHAR(255) UNIQUE NOT NULL CHECK (c_email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    PRIMARY KEY (chain_ID, c_email)
);

-- Insert into ChainEmail
INSERT INTO ChainEmail (chain_ID, c_email) VALUES
(1, 'info@chain1.com'),
(2, 'contact@chain2.com'),
(3, 'support@chain3.com'),
(4, 'info@chain4.com'),
(5, 'contact@chain5.com'),
(6, 'support@chain6.com'),
(7, 'info@chain7.com'),
(8, 'contact@chain8.com'),
(9, 'support@chain9.com'),
(10, 'info@chain10.com');

CREATE TABLE HotelPhone (
    hotel_ID INT NOT NULL,
    h_phone VARCHAR(20) UNIQUE NOT NULL CHECK (h_phone ~ '^[0-9-]+$'),
    PRIMARY KEY (hotel_ID, h_phone)
);

-- Insert into HotelPhone
INSERT INTO HotelPhone (hotel_ID, h_phone) VALUES
(1, '212-555-1111'),
(2, '310-555-2222'),
(3, '312-555-3333'),
(4, '713-555-4444'),
(5, '305-555-5555'),
(6, '206-555-6666'),
(7, '303-555-7777'),
(8, '617-555-8888'),
(9, '404-555-9999'),
(10, '415-555-0000');

CREATE TABLE HotelEmail (
    hotel_ID INT NOT NULL,
    h_email VARCHAR(255) UNIQUE NOT NULL CHECK (h_email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    PRIMARY KEY (hotel_ID, h_email)
);

-- Insert into HotelEmail
INSERT INTO HotelEmail (hotel_ID, h_email) VALUES
(1, 'nyc@hotel1.com'),
(2, 'la@hotel2.com'),
(3, 'chicago@hotel3.com'),
(4, 'houston@hotel4.com'),
(5, 'miami@hotel5.com'),
(6, 'seattle@hotel6.com'),
(7, 'denver@hotel7.com'),
(8, 'boston@hotel8.com'),
(9, 'atlanta@hotel9.com'),
(10, 'sf@hotel10.com');

-- Employee Table
CREATE TABLE Employee (
    employee_ID SERIAL PRIMARY KEY,
    hotel_ID INT NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    ID_type VARCHAR(50) NOT NULL CHECK (ID_type IN ('SSN', 'SIN')),
    ID_number VARCHAR(255) UNIQUE NOT NULL,
    role VARCHAR(50) NOT NULL CHECK (role IN ('Manager', 'Employee'))
);

-- Insert into Employee
INSERT INTO Employee (hotel_ID, full_name, address, SSN_SIN, role) VALUES
(1, 'John Doe', '500 Fifth Ave, New York, NY', '123-45-6789', 'Manager'),
(2, 'Jane Smith', '123 Hollywood Blvd, Los Angeles, CA', '987-65-4321', 'Manager'),
(3, 'Michael Brown', '789 Michigan Ave, Chicago, IL', '567-89-1234', 'Manager'),
(4, 'Emily Davis', '400 Westheimer Rd, Houston, TX', '111-22-3333', 'Manager'),
(5, 'James Wilson', '505 Ocean Dr, Miami, FL', '222-33-4444', 'Manager'),
(6, 'Sarah Miller', '606 Rainier Ave, Seattle, WA', '333-44-5555', 'Manager'),
(7, 'David Taylor', '707 Colfax Ave, Denver, CO', '444-55-6666', 'Manager'),
(8, 'Emma Anderson', '808 Beacon St, Boston, MA', '555-66-7777', 'Manager'),
(9, 'Daniel Martinez', '909 Peachtree St, Atlanta, GA', '666-77-8888', 'Manager'),
(10, 'Sophia Thomas', '1010 Market St, San Francisco, CA', '777-88-9999', 'Manager');


-- Customer Table
CREATE TABLE Customer (
    customer_ID SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    ID_type VARCHAR(50) NOT NULL CHECK (ID_type IN ('SSN', 'SIN', 'Driver License')),
    ID_number VARCHAR(255) UNIQUE NOT NULL,
    registration_date VARCHAR(255) UNIQUE NOT NULL
);

-- Insert into Customer
INSERT INTO Customer (full_name, ID_type, ID_number, registration_date) VALUES
('Alice Johnson', 'SSN', 'A12345678', '2024-01-15'),
('Bob Williams', 'Driver License', 'B98765432', '2024-02-10'),
('Charlie Davis', 'SIN', 'C56789012', '2024-03-05'),
('Diana White', 'SSN', 'D11223344', '2024-04-20'),
('Eric Harris', 'Driver License', 'E22334455', '2024-05-18'),
('Fiona Lewis', 'SIN', 'F33445566', '2024-06-25'),
('George Clark', 'SSN', 'G44556677', '2024-07-12'),
('Hannah Young', 'Driver License', 'H55667788', '2024-08-30'),
('Ian Hall', 'SIN', 'I66778899', '2024-09-10'),
('Jessica Allen', 'SSN', 'J77889900', '2024-10-22');


-- Room Table
CREATE TABLE Room (
    room_ID SERIAL PRIMARY KEY,
    hotel_ID INT NOT NULL,
    room_number VARCHAR(10) NOT NULL,
    capacity INT NOT NULL CHECK (capacity > 0),
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    view_type VARCHAR(20) CHECK (view_type IN ('Mountain', 'Sea')),
    extendable BOOLEAN NOT NULL,
    damaged BOOLEAN NOT NULL
);

-- Insert into Room
INSERT INTO Room (hotel_ID, room_number, capacity, price, view_type, extendable, damaged) VALUES
(1, '101', 2, 150.00, 'Sea', TRUE, FALSE),
(2, '202', 3, 200.00, 'Mountain', FALSE, FALSE),
(3, '303', 2, 175.00, 'Sea', TRUE, FALSE),
(4, '404', 4, 250.00, 'Mountain', TRUE, FALSE),
(5, '505', 2, 160.00, 'Sea', FALSE, FALSE),
(6, '606', 3, 190.00, 'Mountain', TRUE, FALSE),
(7, '707', 2, 140.00, 'Sea', FALSE, FALSE),
(8, '808', 4, 230.00, 'Mountain', TRUE, FALSE),
(9, '909', 3, 180.00, 'Sea', FALSE, FALSE),
(10, '1010', 2, 155.00, 'Mountain', TRUE, FALSE);

-- Amenities Table
CREATE TABLE Amenities (
    amenity_ID SERIAL PRIMARY KEY,
    amenity_name VARCHAR(255) NOT NULL UNIQUE
);

-- Insert into Amenities
INSERT INTO Amenities (amenity_name) VALUES
('WiFi'),
('Air Conditioning'),
('Television'),
('Mini Fridge'),
('Room Service'),
('Gym Access'),
('Pool Access'),
('Spa Access'),
('Balcony'),
('Coffee Maker');

-- Room_Has_Amenities 
CREATE TABLE Room_Has_Amenities (
    room_ID INT NOT NULL,
    amenity_ID INT NOT NULL,
    PRIMARY KEY (room_ID, amenity_ID)
);

-- Insert into Room_Has_Amenities
INSERT INTO Room_Has_Amenities (room_ID, amenity_ID) VALUES
(1, 1), (1, 2), (1, 3),
(2, 1), (2, 4), (2, 5),
(3, 1), (3, 6), (3, 7),
(4, 2), (4, 8), (4, 9),
(5, 3), (5, 4), (5, 10);

-- Booking Table
CREATE TABLE Booking (
    booking_ID SERIAL PRIMARY KEY,
    customer_ID INT NOT NULL,
    room_ID INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    total_price DECIMAL(10,2) NOT NULL CHECK (total_price >= 0),
    CHECK (start_Date < end_Date)
);

-- Insert into Booking
INSERT INTO Booking (customer_ID, room_ID, start_date, end_date, total_price) VALUES
(1, 1, '2025-04-01', '2025-04-05', 600.00),
(2, 2, '2025-05-10', '2025-05-15', 1000.00),
(3, 3, '2025-06-20', '2025-06-25', 875.00),
(4, 4, '2025-07-15', '2025-07-20', 1250.00),
(5, 5, '2025-08-05', '2025-08-10', 800.00);

-- Renting Table
CREATE TABLE Renting (
    renting_ID SERIAL PRIMARY KEY,
    employee_ID INT NOT NULL,
    customer_ID INT NOT NULL,
    room_ID INT NOT NULL,
    booking_ID INT UNIQUE,
    check_in_date DATE NOT NULL,
    check_out_date DATE NOT NULL,
    total_price DECIMAL(10,2) NOT NULL CHECK (total_price >= 0),
    CHECK (check_in_date <= check_out_date)
);

-- Insert into Renting
INSERT INTO Renting (employee_ID, customer_ID, room_ID, booking_ID, check_in_date, check_out_date, total_price) VALUES
(1, 1, 1, 1, '2025-04-01', '2025-04-05', 600.00),
(2, 2, 2, 2, '2025-05-10', '2025-05-15', 1000.00),
(3, 3, 3, 3, '2025-06-20', '2025-06-25', 875.00),
(4, 4, 4, 4, '2025-07-15', '2025-07-20', 1250.00),
(5, 5, 5, 5, '2025-08-05', '2025-08-10', 800.00);

-- Archive Table  
CREATE TABLE Archive (
    archive_ID SERIAL PRIMARY KEY,
    renting_ID INT,
    booking_ID INT,
    total_price DECIMAL(10,2) NOT NULL CHECK (total_price >= 0)
);

-- Insert into Archive
INSERT INTO Archive (renting_ID, booking_ID, total_price) VALUES
(1, 1, 600.00),
(2, 2, 1000.00),
(3, 3, 875.00),
(4, 4, 1250.00),
(5, 5, 800.00);


/* Alters tables to connect the foreign keys */

-- Hotel references HotelChain and Employee (Manager)
ALTER TABLE Hotel
	ADD FOREIGN KEY (chain_ID) REFERENCES HotelChain(chain_ID) ON DELETE CASCADE,
    ADD FOREIGN KEY (manager_ID) REFERENCES Employee(employee_ID) ON DELETE SET NULL;

-- ChainPhone references HotelChain 
ALTER TABLE ChainPhone
    ADD FOREIGN KEY (chain_ID) REFERENCES HotelChain(chain_ID) ON DELETE CASCADE;

-- ChainEmail references HotelChain 
ALTER TABLE ChainEmail
    ADD FOREIGN KEY (chain_ID) REFERENCES HotelChain(chain_ID) ON DELETE CASCADE;

-- HotelPhone references HotelChain 
ALTER TABLE HotelPhone
    ADD FOREIGN KEY (hotel_ID) REFERENCES Hotel(hotel_ID) ON DELETE CASCADE;

-- HotelEmail references HotelChain 
ALTER TABLE HotelEmail
    ADD FOREIGN KEY (hotel_ID) REFERENCES Hotel(hotel_ID) ON DELETE CASCADE;

-- Employee references Hotel
ALTER TABLE Employee
    ADD FOREIGN KEY (hotel_ID) REFERENCES Hotel(hotel_ID) ON DELETE CASCADE;

-- Room references Hotel
ALTER TABLE Room
    ADD FOREIGN KEY (hotel_ID) REFERENCES Hotel(hotel_ID) ON DELETE CASCADE;

-- Room_Has_Amenities references Room and Amenities
ALTER TABLE Room_Has_Amenities
    ADD FOREIGN KEY (room_ID) REFERENCES Room(room_ID) ON DELETE CASCADE,
    ADD FOREIGN KEY (amenity_ID) REFERENCES Amenities(amenity_ID) ON DELETE CASCADE;

-- Booking references Customer and Room
ALTER TABLE Booking
    ADD FOREIGN KEY (customer_ID) REFERENCES Customer(customer_ID) ON DELETE CASCADE,
    ADD FOREIGN KEY (room_ID) REFERENCES Room(room_ID) ON DELETE CASCADE;

-- Renting references Employee, Customer, Room, and Booking
ALTER TABLE Renting
    ADD FOREIGN KEY (employee_ID) REFERENCES Employee(employee_ID) ON DELETE CASCADE,
    ADD FOREIGN KEY (customer_ID) REFERENCES Customer(customer_ID) ON DELETE CASCADE,
    ADD FOREIGN KEY (room_ID) REFERENCES Room(room_ID) ON DELETE CASCADE,
    ADD FOREIGN KEY (booking_ID) REFERENCES Booking(booking_ID) ON DELETE SET NULL;

-- Archive references Renting and Booking
ALTER TABLE Archive
    ADD FOREIGN KEY (renting_ID) REFERENCES Renting(renting_ID) ON DELETE SET NULL,
    ADD FOREIGN KEY (booking_ID) REFERENCES Booking(booking_ID) ON DELETE SET NULL;