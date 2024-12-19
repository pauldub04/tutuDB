CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    phone VARCHAR(50),
    created_at TIMESTAMP
);
CREATE TABLE promotions (
    promotion_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    discount DECIMAL(5, 2),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    conditions TEXT
);
CREATE TABLE stations (
    station_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    city VARCHAR(255),
    type VARCHAR(50),
    latitude DECIMAL(9, 6),
    longitude DECIMAL(9, 6)
);
CREATE TABLE operators (
    operator_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    contact_info TEXT,
    country VARCHAR(100)
);
CREATE TABLE transport (
    transport_id SERIAL PRIMARY KEY,
    type VARCHAR(50),
    model VARCHAR(255),
    capacity INT,
    operator_id INT,
    FOREIGN KEY (operator_id) REFERENCES operators(operator_id)
);
CREATE TABLE routes (
    route_id SERIAL PRIMARY KEY,
    start_station_id INT,
    end_station_id INT,
    duration INT,
    transport_type VARCHAR(50),
    operator_id INT,
    FOREIGN KEY (start_station_id) REFERENCES stations(station_id),
    FOREIGN KEY (end_station_id) REFERENCES stations(station_id),
    FOREIGN KEY (operator_id) REFERENCES operators(operator_id)
);
CREATE TABLE schedules (
    schedule_id SERIAL PRIMARY KEY,
    route_id INT,
    departure_time TIMESTAMP,
    arrival_time TIMESTAMP,
    days_of_week VARCHAR(20),
    transport_id INT,
    FOREIGN KEY (route_id) REFERENCES routes(route_id),
    FOREIGN KEY (transport_id) REFERENCES transport(transport_id)
);
CREATE TABLE passengers (
    passenger_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    document_type VARCHAR(255) NOT NULL,
    document_number VARCHAR(255) NOT NULL
);
CREATE TABLE tickets (
    ticket_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    schedule_id INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    seat_number VARCHAR(10) NOT NULL,
    purchase_date TIMESTAMP NOT NULL,
    promotion_id INT NULL,
    passenger_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (schedule_id) REFERENCES schedules(schedule_id),
    FOREIGN KEY (promotion_id) REFERENCES promotions(promotion_id),
    FOREIGN KEY (passenger_id) REFERENCES passengers(passenger_id)
);
CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    ticket_id INT,
    amount DECIMAL(10, 2),
    payment_date TIMESTAMP,
    payment_method VARCHAR(50),
    status VARCHAR(50),
    FOREIGN KEY (ticket_id) REFERENCES tickets(ticket_id)
);
CREATE TABLE feedback (
    feedback_id SERIAL PRIMARY KEY,
    user_id INT,
    route_id INT,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (route_id) REFERENCES routes(route_id)
);
