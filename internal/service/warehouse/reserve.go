package warehouse

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sarastee/LamodaTest/internal/service"
)

// Reserve is Service layer method, which process request from API layer
func (s *Service) Reserve(ctx context.Context, codes []int32) error {
	s.logger.Debug().Msg("attempt to reserve products")

	m := make(map[int32]int32)
	// реализация резерва нескольких товаров с одинаковым уникальным кодом
	// заменить int32 на структуру с количеством

	for _, code := range codes {
		whIDs, err := s.warehouseRepo.GetWarehousesIDByProductCode(ctx, code)
		if err != nil {
			s.logger.Err(err).Msg("unable to get warehouse ID by code")
			return fmt.Errorf("failure while getting warehouse ID: %w", err)
		}
		if len(whIDs) == 0 {
			s.logger.Err(service.ErrMsgWarehouseNoAvailableWarehouses).
				Msgf("no available warehouses with this product: %d", code)
			return errors.Wrapf(service.ErrMsgWarehouseNoAvailableWarehouses, "product code: %d", code)
		}

		for _, whID := range whIDs {
			amount, err := s.warehouseRepo.GetWarehouseProductAmount(ctx, code, whID)
			if err != nil {
				s.logger.Err(err).Msg("unable to check product amount")
				return fmt.Errorf("failure while checking product amount on warehouse: %w", err)
			}
			if amount > 0 {
				m[code] = whID
				break
			}
		}

		if m[code] == 0 {
			s.logger.Err(service.ErrMsgNotEnoughProductsOnWarehouse).
				Msgf("not enough products (product ID: %d) on warehouses", code)
			return errors.Wrapf(service.ErrMsgNotEnoughProductsOnWarehouse, "product code: %d", code)
		}
	}

	for code, whID := range m {
		err := s.reserveRepo.ReserveProduct(ctx, code, whID)
		if err != nil {
			s.logger.Err(err).Msg("unable to reserve product")
			return fmt.Errorf("failure while reserving product: %w", err)
		}
	}

	return nil
}
