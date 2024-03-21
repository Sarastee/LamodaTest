package warehouse

import (
	"github.com/Masterminds/squirrel"
	"github.com/rs/zerolog"
	"github.com/sarastee/LamodaTest/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	warehouseTable    = "warehouses"
	warehouseIDColumn = "id"
	//warehouseNameColumn         = "name"
	//warehouseAvailabilityColumn = "availability"

	//productTable      = "products"
	//productCodeColumn = "code"
	//productNameColumn = "name"
	//productSizeColumn = "size"

	WarehouseProductTable             = "warehouse_products" // WarehouseProductTable contains warehouse_products table name
	WarehouseProductCodeColumn        = "product_code"       // WarehouseProductCodeColumn contains warehouse_products code column name
	WarehouseProductWarehouseIDColumn = "warehouse_id"       // WarehouseProductWarehouseIDColumn contains warehouse_products warehouse_id column
	WarehouseProductAmountColumn      = "amount"             // WarehouseProductAmountColumn contains warehouse_products amount column
)

var _ repository.WarehouseRepository = (*Repo)(nil)

// Repo is a struct, containing zerolog.logger, db.Client and squirrel.StatementBuilderType
type Repo struct {
	logger *zerolog.Logger
	db     db.Client
	sq     squirrel.StatementBuilderType
}

// NewRepo returns pointer for Repo struct
func NewRepo(logger *zerolog.Logger, dbClient db.Client) *Repo {
	return &Repo{
		logger: logger,
		db:     dbClient,
	}
}
