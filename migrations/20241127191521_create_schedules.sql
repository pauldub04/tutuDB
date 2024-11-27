-- +goose Up
-- +goose StatementBegin
CREATE TABLE schedules (
    schedule_id SERIAL PRIMARY KEY,
    route_id INT,
    departure_time TIMESTAMP,
    arrival_time TIMESTAMP,
    days_of_week VARCHAR(7),
    transport_id INT,
    FOREIGN KEY (route_id) REFERENCES routes(route_id),
    FOREIGN KEY (transport_id) REFERENCES transport(transport_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS schedules;
-- +goose StatementEnd
