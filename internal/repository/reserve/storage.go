package reserve

import (
	"github.com/Masterminds/squirrel"
	"github.com/rs/zerolog"
	"github.com/sarastee/LamodaTest/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	reservedProductsTable            = "reserved_products"
	reservedProductCodeColumn        = "product_code"
	reservedProductWarehouseIDColumn = "warehouse_id"
	reservedProductAmountColumn      = "amount"
)

var _ repository.ReserveRepository = (*Repo)(nil)

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
