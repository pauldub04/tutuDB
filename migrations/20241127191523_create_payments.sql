-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    ticket_id INT,
    amount DECIMAL(10, 2),
    payment_date TIMESTAMP,
    payment_method VARCHAR(50),
    status VARCHAR(50),
    FOREIGN KEY (ticket_id) REFERENCES tickets(ticket_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments;
-- +goose StatementEnd
