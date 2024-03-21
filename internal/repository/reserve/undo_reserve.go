package reserve

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/LamodaTest/internal/repository/warehouse"
)

// UndoReserve is repo layer method, which undo reserve
func (r *Repo) UndoReserve(ctx context.Context, code int32, whID int32) error {
	tx, err := r.db.DB().BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}

	defer func() {
		err = tx.Rollback(ctx)
		if err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			r.logger.Warn().Msg("error while rollback in reserve_repository.UndoReserve")
		}
	}()

	builderInsertOrUpdate := r.sq.Insert(warehouse.WarehouseProductTable).
		Columns(
			warehouse.WarehouseProductCodeColumn,
			warehouse.WarehouseProductWarehouseIDColumn,
			warehouse.WarehouseProductAmountColumn).
		Values(code, whID, 1).
		Suffix("ON CONFLICT (product_code, warehouse_id) DO UPDATE SET amount = warehouse_products.amount + 1")

	query, args, err := builderInsertOrUpdate.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	builderUpdate := r.sq.Update(reservedProductsTable).
		Where(squirrel.And{
			squirrel.Eq{reservedProductCodeColumn: code},
			squirrel.Eq{reservedProductWarehouseIDColumn: whID},
		}).
		Set(reservedProductAmountColumn, squirrel.Expr("amount - 1"))

	query, args, err = builderUpdate.PlaceholderFormat(squirrel.Dollar).ToSql()
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
			squirrel.Eq{reservedProductWarehouseIDColumn: whID},
			squirrel.Eq{reservedProductAmountColumn: 0},
		})

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
