package reserve

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

// GetAmountByProductCodeAndWarehouseID is a repo layer method, which returns amount of products in reserved_products table
// by given code and warehouse_id
func (r *Repo) GetAmountByProductCodeAndWarehouseID(ctx context.Context, code int32, whID int32) (int64, error) {
	builderSelect := r.sq.Select(reservedProductAmountColumn).
		From(reservedProductsTable).
		Where(squirrel.And{
			squirrel.Eq{reservedProductCodeColumn: code},
			squirrel.Eq{reservedProductWarehouseIDColumn: whID}})

	query, args, err := builderSelect.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "reserve_repository.GetAmountByProductCodeAndWarehouseID",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	amount, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, err
	}

	return amount, nil
}
