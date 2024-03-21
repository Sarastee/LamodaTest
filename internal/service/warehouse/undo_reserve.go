package warehouse

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/sarastee/LamodaTest/internal/service"
)

// UndoReserve is Service layer method, which process request from API layer
func (s *Service) UndoReserve(ctx context.Context, codes []int32) error {
	s.logger.Debug().Msg("attempt to undo reserve products")

	// TODO: Проверить склад на доступность??? Можем ли мы отменить резерв если склад недоступен, по логике - да
	m := make(map[int32]int32)

	for _, code := range codes {
		whIDs, err := s.reserveRepo.GetWarehousesIDInReserveList(ctx, code)
		if err != nil {
			s.logger.Err(err).Msg("unable to get warehouse ID by code")
			return fmt.Errorf("failure while getting warehouse ID: %w", err)
		}
		if len(whIDs) == 0 {
			s.logger.Err(service.ErrMsgNoWarehousesWithThisCodeInReserveList).
				Msgf("no available warehouses in reserve list with this product: %d", code)
			return errors.Wrapf(service.ErrMsgNoWarehousesWithThisCodeInReserveList, "product code: %d", code)
		}

		for _, whID := range whIDs {
			amount, err := s.reserveRepo.GetAmountByProductCodeAndWarehouseID(ctx, code, whID)
			if err != nil {
				s.logger.Err(err).Msg("unable to check product amount")
				return fmt.Errorf("failure while checking product amount on reserve list: %w", err)
			}
			if amount > 0 {
				m[code] = whID
				break
			}
		}

		if m[code] == 0 {
			s.logger.Err(service.ErrMsgNotEnoughProductsInReserveList).
				Msgf("not enough products (product ID: %d) on reserve list", code)
			return errors.Wrapf(service.ErrMsgNotEnoughProductsInReserveList, "product code: %d", code)
		}
	}

	for code, whID := range m {
		log.Printf("%d : %d", code, whID)
		err := s.reserveRepo.UndoReserve(ctx, code, whID)
		if err != nil {
			s.logger.Err(err).Msg("unable to undo reservation")
			return fmt.Errorf("failure while trying to undo reserve product: %w", err)
		}
	}

	return nil
}
