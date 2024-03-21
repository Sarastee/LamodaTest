package warehouse

import (
	"context"
	"errors"

	"github.com/sarastee/LamodaTest/internal/service"
	"github.com/sarastee/LamodaTest/pkg/warehouse_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UndoReserve is API layer method, returning message
func (i *Implementation) UndoReserve(ctx context.Context, request *warehouse_v1.UndoReserveRequest) (*warehouse_v1.UndoReserveResponse, error) {
	err := i.warehouseService.UndoReserve(ctx, request.Codes)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrMsgNoWarehousesWithThisCodeInReserveList):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, service.ErrMsgNotEnoughProductsInReserveList):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, errInternal
	}

	return &warehouse_v1.UndoReserveResponse{Message: "Products reservation successfully undo"}, nil
}
