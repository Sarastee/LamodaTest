package warehouse

import (
	"errors"

	"github.com/sarastee/LamodaTest/internal/service"
	"github.com/sarastee/LamodaTest/pkg/warehouse_v1"
)

const msgInternalError = "something went wrong, we already working on it"

var errInternal = errors.New(msgInternalError)

type Implementation struct {
	warehouse_v1.UnimplementedWarehouseV1Server
	warehouseService service.WarehouseService
}

func NewImplementation(warehouseService service.WarehouseService) *Implementation {
	return &Implementation{
		warehouseService: warehouseService,
	}
}
