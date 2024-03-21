package warehouse

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

// GetWarehousesIDByProductCode returns []int32 of available warehouses
func (r *Repo) GetWarehousesIDByProductCode(ctx context.Context, code int32) ([]int32, error) {
	builderSelect := r.sq.Select(WarehouseProductWarehouseIDColumn).
		From(WarehouseProductTable).
		Join(warehouseTable + " ON " + WarehouseProductTable + ".warehouse_id = warehouses.id").
		Where(squirrel.And{squirrel.Eq{WarehouseProductCodeColumn: code}, squirrel.Eq{"warehouses.availability": true}}).
		OrderBy(WarehouseProductAmountColumn + " DESC")

	query, args, err := builderSelect.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "warehouse_repository.GetWarehousesIDByProductCode",
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
