package service

import "context"

// WarehouseService is interface for Warehouse Service
type WarehouseService interface {
	Reserve(ctx context.Context, codes []int32) error
	Release(ctx context.Context, codes []int32) error
	GetAll(ctx context.Context, whID int32) (int64, error)
	UndoReserve(ctx context.Context, codes []int32) error
}
