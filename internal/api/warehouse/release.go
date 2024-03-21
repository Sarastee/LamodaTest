package warehouse

import (
	"context"
	"errors"

	"github.com/sarastee/LamodaTest/internal/service"
	"github.com/sarastee/LamodaTest/pkg/warehouse_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Release is API layer method, returning message
func (i *Implementation) Release(ctx context.Context, request *warehouse_v1.ReleaseRequest) (*warehouse_v1.ReleaseResponse, error) {
	err := i.warehouseService.Release(ctx, request.Codes)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrMsgNotEnoughProductsInReserveList):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, service.ErrMsgNoProductsWithCodeInReserveList):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, errInternal
	}

	return &warehouse_v1.ReleaseResponse{Message: "Products successfully released"}, nil
}
