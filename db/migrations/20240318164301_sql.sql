-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS warehouses (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    availability BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    size TEXT NOT NULL,
    unique_code TEXT UNIQUE NOT NULL,
    quantity INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table warehouses;
drop table products;
-- +goose StatementEnd
