package warehouse

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/sarastee/LamodaTest/internal/service"
)

// Release is Service layer method, which process request from API layer
func (s *Service) Release(ctx context.Context, codes []int32) error {
	s.logger.Debug().Msg("attempt to release products")

	for _, code := range codes {
		amount, err := s.reserveRepo.GetAmountByProductCode(ctx, code)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				s.logger.Err(service.ErrMsgNoProductsWithCodeInReserveList).
					Msgf("there are no items with this code on the reserve list. product code: %d", code)
				return errors.Wrapf(service.ErrMsgNoProductsWithCodeInReserveList, "product code: %d", code)
			}
			s.logger.Err(err).Msg("unable to check product amount")
			return fmt.Errorf("failure while checking product amount on reserve list: %w", err)
		}
		if amount == 0 {
			s.logger.Err(service.ErrMsgNotEnoughProductsInReserveList).
				Msgf("not enough products on reserve list. product code: %d", code)
			return errors.Wrapf(service.ErrMsgNotEnoughProductsInReserveList, "product code: %d", code)
		}
	}

	for _, code := range codes {
		err := s.reserveRepo.Release(ctx, code)
		if err != nil {
			s.logger.Err(err).Msg("unable to release product")
			return fmt.Errorf("failure while releasing product: %w", err)
		}

	}

	return nil
}
