package reserve

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/LamodaTest/internal/repository/warehouse"
)

// ReserveProduct is repo layer method, which reserve product
func (r *Repo) ReserveProduct(ctx context.Context, code int32, whID int32) error {
	tx, err := r.db.DB().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}

	defer func() {
		err = tx.Rollback(ctx)
		if err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			r.logger.Warn().Msg("error while rollback in reserve_repository.Reserve")
		}
	}()

	builderInsertOrUpdate := r.sq.Insert(reservedProductsTable).
		Columns(reservedProductCodeColumn, reservedProductWarehouseIDColumn, reservedProductAmountColumn).
		Values(code, whID, 1).
		Suffix("ON CONFLICT (product_code, warehouse_id) DO UPDATE SET amount = reserved_products.amount + 1")

	query, args, err := builderInsertOrUpdate.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	builderUpdate := r.sq.Update(warehouse.WarehouseProductTable).
		Where(squirrel.And{
			squirrel.Eq{warehouse.WarehouseProductCodeColumn: code},
			squirrel.Eq{warehouse.WarehouseProductWarehouseIDColumn: whID}}).
		Set(warehouse.WarehouseProductAmountColumn, squirrel.Expr(warehouse.WarehouseProductAmountColumn+" - 1"))

	query, args, err = builderUpdate.PlaceholderFormat(squirrel.Dollar).ToSql()
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
