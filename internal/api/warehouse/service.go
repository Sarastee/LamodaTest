package warehouse

import (
	"errors"

	"github.com/sarastee/LamodaTest/internal/service"
	"github.com/sarastee/LamodaTest/pkg/warehouse_v1"
)

const msgInternalError = "something went wrong, we already working on it"

var errInternal = errors.New(msgInternalError)

// Implementation is a struct, containing Interfaces for warehouse_v1.UnimplementedWarehouseV1Server and service.WarehouseService
type Implementation struct {
	warehouse_v1.UnimplementedWarehouseV1Server
	warehouseService service.WarehouseService
}

// NewImplementation is a method, returning pointer for Implementation struct
func NewImplementation(warehouseService service.WarehouseService) *Implementation {
	return &Implementation{
		warehouseService: warehouseService,
	}
}
