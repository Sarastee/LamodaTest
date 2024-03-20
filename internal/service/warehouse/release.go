package warehouse

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/sarastee/LamodaTest/internal/service"
)

func (s *Service) Release(ctx context.Context, codes []int32) error {
	s.logger.Debug().Msg("attempt to release products")

	// m := make(map[int32]int64)

	for _, code := range codes {
		amount, err := s.reserveRepo.GetAmountByProductCode(ctx, code)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				s.logger.Err(service.ErrNoProductsWithCodeInReserveList).
					Msgf("there are no items with this code on the reserve list. product code: %d", code)
				return errors.Wrapf(service.ErrNoProductsWithCodeInReserveList, "product code: %d", code)
			}
			s.logger.Err(err).Msg("unable to check product amount")
			return fmt.Errorf("failure while checking product amount on reserve list: %w", err)
		}
		if amount == 0 {
			s.logger.Err(service.ErrNotEnoughProductsInReserveList).
				Msgf("not enough products in reserve list. product code: %d", code)
			return errors.Wrapf(service.ErrNotEnoughProductsInReserveList, "product code: %d", code)
		}
	}

	// TODO: метод уменьшающий значение резерва на 1, если amount становится == 0, то удалить элемент таблицы

	return nil
}
