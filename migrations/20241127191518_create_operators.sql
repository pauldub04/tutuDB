-- +goose Up
-- +goose StatementBegin
CREATE TABLE operators (
    operator_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    contact_info TEXT,
    country VARCHAR(100)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS operators;
-- +goose StatementEnd
