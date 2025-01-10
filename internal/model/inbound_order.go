package model

import (
	"time"
)

type InboundOrder struct {
	Id             int
	OrderDate      string
	OrderNumber    string
	EmployeeId     int
	ProductBatchId int
	WareHouseId    int
}

func (i *InboundOrder) IsValid() bool {
	if i.EmployeeId <= 0 || i.ProductBatchId <= 0 || i.WareHouseId <= 0 {
		return false
	}

	if _, err := time.Parse("2006-01-02", i.OrderDate); err != nil {
		return false
	}

	if i.OrderNumber == "" {
		return false
	}

	return true
}
