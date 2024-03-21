package reserve

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

// GetWarehousesIDInReserveList is repo layer method, returning slice of warehouse_id for given code
func (r *Repo) GetWarehousesIDInReserveList(ctx context.Context, code int32) ([]int32, error) {
	builderSelect := r.sq.Select(reservedProductWarehouseIDColumn).
		From(reservedProductsTable).
		Where(squirrel.Eq{reservedProductCodeColumn: code})

	query, args, err := builderSelect.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "reserve_repository.GetWarehousesIdInReserveList",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	whIDs, err := pgx.CollectRows(rows, pgx.RowTo[int32])
	if err != nil {
		return nil, err
	}

	return whIDs, nil
}
