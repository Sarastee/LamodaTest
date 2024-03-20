package warehouse

import (
	"context"
	"errors"

	"github.com/sarastee/LamodaTest/internal/service"
	"github.com/sarastee/LamodaTest/pkg/warehouse_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Reserve(ctx context.Context, request *warehouse_v1.ReserveRequest) (*warehouse_v1.ReserveResponse, error) {
	err := i.warehouseService.Reserve(ctx, request.Codes)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrMsgWarehouseNoAvailableWarehouses):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, service.ErrMsgNotEnoughProductsOnWarehouse):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, errInternal
		}
	}

	return &warehouse_v1.ReserveResponse{Message: "Products successfully reserved"}, nil
}
