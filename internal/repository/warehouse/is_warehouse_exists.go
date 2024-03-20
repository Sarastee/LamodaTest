package warehouse

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

func (r *Repo) IsWarehouseExists(ctx context.Context, whID int32) (bool, error) {
	builderSelect := r.sq.Select("TRUE").
		From(warehouseTable).
		Where(squirrel.Eq{warehouseIDColumn: whID})

	query, args, err := builderSelect.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return false, err
	}

	q := db.Query{
		Name:     "warehouse_repository.IsWarehouseExists",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	_, err = pgx.CollectOneRow(rows, pgx.RowTo[bool])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
