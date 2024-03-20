package warehouse

import (
	"github.com/Masterminds/squirrel"
	"github.com/rs/zerolog"
	"github.com/sarastee/LamodaTest/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	warehouseTable              = "warehouses"
	warehouseIDColumn           = "id"
	warehouseNameColumn         = "name"
	warehouseAvailabilityColumn = "availability"

	productTable      = "products"
	productCodeColumn = "code"
	productNameColumn = "name"
	productSizeColumn = "size"

	WarehouseProductTable             = "warehouse_products"
	WarehouseProductCodeColumn        = "product_code"
	WarehouseProductWarehouseIDColumn = "warehouse_id"
	WarehouseProductAmountColumn      = "amount"
)

var _ repository.WarehouseRepository = (*Repo)(nil)

type Repo struct {
	logger *zerolog.Logger
	db     db.Client
	sq     squirrel.StatementBuilderType
}

func NewRepo(logger *zerolog.Logger, dbClient db.Client) *Repo {
	return &Repo{
		logger: logger,
		db:     dbClient,
	}
}
