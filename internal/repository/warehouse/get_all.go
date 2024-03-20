package warehouse

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

func (r *Repo) GetAll(ctx context.Context, whID int32) (int64, error) {
	var amount int64

	builderSelect := r.sq.Select("SUM(amount) AS total_amount").
		From(WarehouseProductTable).
		Where(squirrel.Eq{WarehouseProductWarehouseIDColumn: whID})

	query, args, err := builderSelect.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "warehouse_repository.GetAll",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	amount, err = pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, err
	}

	return amount, nil
}
