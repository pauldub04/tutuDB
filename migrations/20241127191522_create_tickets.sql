-- +goose Up
-- +goose StatementBegin
CREATE TABLE tickets (
    ticket_id SERIAL PRIMARY KEY,
    user_id INT,
    schedule_id INT,
    price DECIMAL(10, 2),
    seat_number VARCHAR(10),
    purchase_date TIMESTAMP,
    promotion_id INT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (schedule_id) REFERENCES schedules(schedule_id),
    FOREIGN KEY (promotion_id) REFERENCES promotions(promotion_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tickets;
-- +goose StatementEnd
