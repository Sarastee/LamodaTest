package service

import "context"

type WarehouseService interface {
	Reserve(ctx context.Context, codes []int32) error
	Release(ctx context.Context, codes []int32) error
	GetAll(ctx context.Context, whID int32) (int64, error)
}
