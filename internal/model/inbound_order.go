package model

import (
	"time"
)

type InboundOrder struct {
	Id             int
	OrderDate      time.Time
	OrderNumber    string
	EmployeeId     int
	ProductBatchId int
	WareHouseId    int
}

func (i *InboundOrder) IsValid() bool {
	if i.EmployeeId <= 0 || i.ProductBatchId <= 0 || i.WareHouseId <= 0 {
		return false
	}

	if i.OrderDate.IsZero() {
		return false
	}

	if i.OrderNumber == "" {
		return false
	}

	return true
}
