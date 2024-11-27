-- +goose Up
-- +goose StatementBegin
CREATE TABLE promotions (
    promotion_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    discount DECIMAL(5, 2),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    conditions TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS promotions;
-- +goose StatementEnd
