### Технологии:
- Go 1.21.7
- PostgresSQL 16.1
- Docker

### Библиотеки:
- github.com/Masterminds/squirrel - удобный SQL builder для построения SQL запросов
- github.com/jackc/pgx/v5 - пакет для работы с базой данных
- github.com/rs/zerolog - хороший легковестный логгер с гибкой настройкой
- github.com/sarastee/platform_common - собственная платформенная библиотека с реализацией клиента для БД, Моков, Closer
- github.com/joho/godotenv - пакет для удобного доставания значений окружения из .env файла

- github.com/grpc-ecosystem/grpc-gateway/v2 
- google.golang.org/genproto/googleapis/api
- google.golang.org/grpc
- google.golang.org/protobuf - Пакеты для кодогенерации, реализации gRPC, gRPC-Gateway