package repository

import "context"

type WarehouseRepository interface {
	GetAll(ctx context.Context, whID int32) (int64, error)
	GetWarehousesIDByProductCode(ctx context.Context, code int32) ([]int32, error)
	GetWarehouseProductAmount(ctx context.Context, code int32, whID int32) (int64, error)
	IsWarehouseExists(ctx context.Context, whID int32) (bool, error)
}

type ReserveRepository interface {
	ReserveProduct(ctx context.Context, code int32, whID int32) error
	Release(ctx context.Context, code int32, whID int32) error
	GetAmountByProductCode(ctx context.Context, code int32) (int64, error)
}
