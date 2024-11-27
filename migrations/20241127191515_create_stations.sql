-- +goose Up
-- +goose StatementBegin
CREATE TABLE stations (
    station_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    city VARCHAR(255),
    type VARCHAR(50),
    latitude DECIMAL(9, 6),
    longitude DECIMAL(9, 6)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stations;
-- +goose StatementEnd
