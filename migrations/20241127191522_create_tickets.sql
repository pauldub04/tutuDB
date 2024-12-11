-- +goose Up
-- +goose StatementBegin
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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS passengers;
-- +goose StatementEnd
