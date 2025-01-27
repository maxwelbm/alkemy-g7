package model

type Employee struct {
	ID           int
	CardNumberID string
	FirstName    string
	LastName     string
	WarehouseID  int
}

type InboundOrdersReportByEmployee struct {
	ID                 int    `json:"id"`
	CardNumberID       string `json:"card_number_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	WarehouseID        int    `json:"warehouse_id"`
	InboundOrdersCount int    `json:"inbound_orders_count"`
}

func (e *Employee) IsValidEmployee() bool {
	if e.CardNumberID == "" {
		return false
	}

	if e.FirstName == "" || e.LastName == "" {
		return false
	}

	if e.WarehouseID == 0 {
		return false
	}

	return true
}

func (e *Employee) IsEmptyEmployee() bool {
	return e.CardNumberID == "" && e.FirstName == "" && e.LastName == "" && e.WarehouseID == 0
}
