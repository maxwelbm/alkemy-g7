package model

type Employee struct {
	Id           int
	CardNumberId string
	FirstName    string
	LastName     string
	WarehouseId  int
}

type InboundOrdersReportByEmployee struct {
	Id                 int    `json:"id"`
	CardNumberId       string `json:"card_number_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	WarehouseId        int    `json:"warehouse_id"`
	InboundOrdersCount int    `json:"inbound_orders_count"`
}

func (e *Employee) IsValidEmployee() bool {
	if e.CardNumberId == "" {
		return false
	}
	if e.FirstName == "" || e.LastName == "" {
		return false
	}
	if e.WarehouseId == 0 {
		return false
	}
	return true
}

func (e *Employee) IsEmptyEmployee() bool {
	return e.CardNumberId == "" && e.FirstName == "" && e.LastName == "" && e.WarehouseId == 0
}
