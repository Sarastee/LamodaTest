-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS warehouses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    availability BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    code INT PRIMARY KEY,
    name TEXT NOT NULL,
    size TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS warehouse_products (
    product_code INT REFERENCES products(code),
    warehouse_id INT REFERENCES warehouses(id),
    amount INT NOT NULL,
    PRIMARY KEY (product_code, warehouse_id)
);

CREATE TABLE IF NOT EXISTS reserved_products (
    product_code INT REFERENCES products(code),
    warehouse_id INT REFERENCES warehouses(id),
    amount INT NOT NULL,
    PRIMARY KEY (product_code, warehouse_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS warehouses;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS warehouse_products;
DROP TABLE IF EXISTS reserved_products;
-- +goose StatementEnd
