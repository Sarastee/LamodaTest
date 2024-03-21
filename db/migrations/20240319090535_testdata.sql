-- +goose Up
-- +goose StatementBegin
INSERT INTO products
    (code, name, size)
VALUES
    ('101','Yellow T-Shirt','XS'),
    ('102','Yellow T-Shirt','S'),
    ('103','Yellow T-Shirt','M'),
    ('104','Yellow T-Shirt','L'),
    ('201','White T-Shirt','XS'),
    ('202','White T-Shirt','S'),
    ('203','White T-Shirt','M'),
    ('204','White T-Shirt','L'),
    ('301','Black T-Shirt','XS'),
    ('302','Black T-Shirt','S'),
    ('303','Black T-Shirt','M'),
    ('304','Black T-Shirt','L');

INSERT INTO warehouses
    (id, name, availability)
VALUES
    ('01','Moscow Unit',true),
    ('02','Saint-Petersburg Unit',true),
    ('03','Kazan Unit',false);

INSERT INTO warehouse_products
    (product_code, warehouse_id, amount)
VALUES
    ('204','01', 30),
    ('103','03', 25),
    ('303','02', 15),
    ('104','01', 40),
    ('302','03', 25),
    ('301','01', 5),
    ('102','02', 25),
    ('201', '02', 0),
    ('101', '03', 15);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM warehouse_products;
DELETE FROM warehouses;
DELETE FROM products;
-- +goose StatementEnd
