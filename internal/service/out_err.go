package service

import "errors"

const (
	errWarehouseNotExists             = "warehouse not exists"
	errWarehouseNoAvailableWarehouses = "there are no available warehouses"
	errNotEnoughProductsOnWarehouse   = "not enough products on warehouses"

	errNotEnoughProductsInReserveList        = "not enough products on the reserve list"
	errNoProductsWithCodeInReserveList       = "there are no items with this code on the reserve list"
	errNoWarehousesWithThisCodeInReserveList = "there are no warehouses with this code in reserve list"
)

var ErrMsgWarehouseNotExists = errors.New(errWarehouseNotExists)                                       // ErrMsgWarehouseNotExists ...
var ErrMsgWarehouseNoAvailableWarehouses = errors.New(errWarehouseNoAvailableWarehouses)               // ErrMsgWarehouseNoAvailableWarehouses ...
var ErrMsgNotEnoughProductsOnWarehouse = errors.New(errNotEnoughProductsOnWarehouse)                   // ErrMsgNotEnoughProductsOnWarehouse ...
var ErrMsgNotEnoughProductsInReserveList = errors.New(errNotEnoughProductsInReserveList)               // ErrMsgNotEnoughProductsInReserveList ...
var ErrMsgNoProductsWithCodeInReserveList = errors.New(errNoProductsWithCodeInReserveList)             // ErrMsgNoProductsWithCodeInReserveList ...
var ErrMsgNoWarehousesWithThisCodeInReserveList = errors.New(errNoWarehousesWithThisCodeInReserveList) // ErrMsgNoWarehousesWithThisCodeInReserveList ...
