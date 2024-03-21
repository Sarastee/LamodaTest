package warehouse

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sarastee/LamodaTest/internal/service"
)

// GetAll is Service layer method, which process request from API layer
func (s *Service) GetAll(ctx context.Context, whID int32) (int64, error) {
	s.logger.Debug().Msg(fmt.Sprintf("attempt to get amount of products at the warehouse: %d", whID))

	warehouseExists, err := s.warehouseRepo.IsWarehouseExists(ctx, whID)
	if err != nil {
		s.logger.Err(err).Msg("unable to get amount")
		return 0, fmt.Errorf("failure while checking warehouse for existance: %w", err)
	}

	if !warehouseExists {
		s.logger.Err(service.ErrMsgWarehouseNotExists).Msgf("warehouse not exists. whID: %d", whID)
		return 0, errors.Wrapf(service.ErrMsgWarehouseNotExists, "whID: %d", whID)
	}

	amount, err := s.warehouseRepo.GetAll(ctx, whID)
	if err != nil {
		s.logger.Err(err).Msg("unable to get amount")
		return 0, fmt.Errorf("failure while getting amount of products at the warehouse: %w", err)
	}

	return amount, err
}
