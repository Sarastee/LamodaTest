# Lamoda Test API

// транзакции на товары внутри sql
// третья таблица склад-товар
// резервирование еще таблица транзакция склад-товар

// резервирование товара (codes []int) - транзакция:
    Проверка наличия товара складе (проверка склада на доступность + amount >= 1) 
    -> уменьшение количества на 1 в warehouse_products -> добавление в reserved_products + 1
    TODO: 
        ручка RESERVE(codes []int), 
        ручка REMOVE(code int) у warehouse_products, 
        ручка ADD(product_code int, warehouse_id) у reserved_products if not exists Insert else Set amount++
    
// возврат резерва товаров(code []int) - транзакция:
    Проверка наличия товара в резерве
    -> уменьшение количества на 1 в reserved_products -> увеличение количества на 1 в warehouse_products
    TODO: 
        ручка UNDORESERVE(codes []int), 
        ручка ADD(code int) у warehouse_products,
        ручка REMOVE(product_code int, warehouse_id) Set amount-- else error

// Передача товаров из резерва в доставку(code []int) - транзакции:
    Проверка наличия товара в резерве
    -> уменьшеник количества на 1 в reserved_products

// получение кол-ва оставшихся товаров на складе(warehouse_id int)
    GETALL(warehouse_id) returns Amount (DONE)

// TODO: linter, GitHub Actions, Tests, gRPC-Gateway, swagger???