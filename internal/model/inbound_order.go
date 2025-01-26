package model

import (
	"time"
)

type InboundOrder struct {
	ID             int
	OrderDate      time.Time
	OrderNumber    string
	EmployeeID     int
	ProductBatchID int
	WareHouseID    int
}

func (i *InboundOrder) IsValid() bool {
	if i.EmployeeID <= 0 || i.ProductBatchID <= 0 || i.WareHouseID <= 0 {
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
