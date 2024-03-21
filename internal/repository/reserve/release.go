package reserve

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

// Release is repo layer method, which release product
func (r *Repo) Release(ctx context.Context, code int32) error {
	tx, err := r.db.DB().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}

	defer func() {
		err = tx.Rollback(ctx)
		if err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			r.logger.Warn().Msg("error while rollback in reserve_repository.Release")
		}
	}()

	builderUpdate := r.sq.Update(reservedProductsTable).
		Set(reservedProductAmountColumn, squirrel.Expr(fmt.Sprintf("%s - 1", reservedProductAmountColumn))).
		Where(squirrel.Eq{reservedProductCodeColumn: code})

	query, args, err := builderUpdate.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	builderDelete := r.sq.Delete(reservedProductsTable).
		Where(squirrel.And{
			squirrel.Eq{reservedProductCodeColumn: code},
			squirrel.Eq{reservedProductAmountColumn: 0}})

	query, args, err = builderDelete.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
