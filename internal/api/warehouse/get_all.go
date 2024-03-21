package warehouse

import (
	"context"
	"errors"

	"github.com/sarastee/LamodaTest/internal/service"
	"github.com/sarastee/LamodaTest/pkg/warehouse_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAll is API layer method, returning amount of products
func (i *Implementation) GetAll(ctx context.Context, request *warehouse_v1.GetAllRequest) (*warehouse_v1.GetAllResponse, error) {
	amount, err := i.warehouseService.GetAll(ctx, request.WarehouseId)
	if err != nil {
		if errors.Is(err, service.ErrMsgWarehouseNotExists) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, errInternal
	}

	return &warehouse_v1.GetAllResponse{Amount: amount}, nil
}
