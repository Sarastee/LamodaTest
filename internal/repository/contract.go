package repository

import "context"

// WarehouseRepository is interface for Warehouse Repository
type WarehouseRepository interface {
	GetAll(ctx context.Context, whID int32) (int64, error)
	GetWarehousesIDByProductCode(ctx context.Context, code int32) ([]int32, error)
	GetWarehouseProductAmount(ctx context.Context, code int32, whID int32) (int64, error)
	IsWarehouseExists(ctx context.Context, whID int32) (bool, error)
}

// ReserveRepository is interface for Reserve Repository
type ReserveRepository interface {
	ReserveProduct(ctx context.Context, code int32, whID int32) error
	Release(ctx context.Context, code int32) error
	GetAmountByProductCode(ctx context.Context, code int32) (int64, error)
	GetAmountByProductCodeAndWarehouseID(ctx context.Context, code int32, whID int32) (int64, error)
	UndoReserve(ctx context.Context, code int32, whID int32) error
	GetWarehousesIDInReserveList(ctx context.Context, code int32) ([]int32, error)
}
