-- +goose Up
-- +goose StatementBegin
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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS routes;
-- +goose StatementEnd
