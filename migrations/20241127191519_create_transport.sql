-- +goose Up
-- +goose StatementBegin
CREATE TABLE transport (
    transport_id SERIAL PRIMARY KEY,
    type VARCHAR(50),
    model VARCHAR(255),
    capacity INT,
    operator_id INT,
    FOREIGN KEY (operator_id) REFERENCES operators(operator_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transport;
-- +goose StatementEnd
