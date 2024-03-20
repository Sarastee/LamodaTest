package service

import "errors"

const (
	errWarehouseNotExists             = "warehouse not exists"
	errWarehouseNoAvailableWarehouses = "there are no available warehouses"
	errNotEnoughProductsOnWarehouse   = "not enough products on warehouses"

	errNotEnoughProductsInReserveList  = "not enough products on the reserve list"
	errNoProductsWithCodeInReserveList = "there are no items with this code on the reserve list"
)

var ErrMsgWarehouseNotExists = errors.New(errWarehouseNotExists)
var ErrMsgWarehouseNoAvailableWarehouses = errors.New(errWarehouseNoAvailableWarehouses)
var ErrMsgNotEnoughProductsOnWarehouse = errors.New(errNotEnoughProductsOnWarehouse)
var ErrNotEnoughProductsInReserveList = errors.New(errNotEnoughProductsInReserveList)
var ErrNoProductsWithCodeInReserveList = errors.New(errNoProductsWithCodeInReserveList)
