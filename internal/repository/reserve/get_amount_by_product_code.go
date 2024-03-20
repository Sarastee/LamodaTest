package reserve

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

func (r *Repo) GetAmountByProductCode(ctx context.Context, code int32) (int64, error) {
	builderSelect := r.sq.Select(reservedProductAmountColumn).
		From(reservedProductsTable).
		Where(squirrel.Eq{reservedProductCodeColumn: code})

	query, args, err := builderSelect.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "reserve_repository.GetAmountByProductCode",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	amount, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, pgx.ErrNoRows
		}
		return 0, err
	}

	return amount, nil
}
