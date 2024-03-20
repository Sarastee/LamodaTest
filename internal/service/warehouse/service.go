package warehouse

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/LamodaTest/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

type Service struct {
	logger    *zerolog.Logger
	dbClient  db.Client
	txManager db.TxManager

	warehouseRepo repository.WarehouseRepository

	reserveRepo repository.ReserveRepository
}

func NewService(
	logger *zerolog.Logger,
	dbClient db.Client,
	txManager db.TxManager,
	warehouseRepo repository.WarehouseRepository,
	reserveRepo repository.ReserveRepository,
) *Service {
	return &Service{
		logger:        logger,
		dbClient:      dbClient,
		txManager:     txManager,
		warehouseRepo: warehouseRepo,
		reserveRepo:   reserveRepo,
	}
}
