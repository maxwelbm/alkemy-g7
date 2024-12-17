package handler

type EmployeeJSON struct {
	Id           int    `json:"id,omitempty"`
	CardNumberId string `json:"card_number_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	WarehouseId  int    `json:"warehouse_id,omitempty"`
}
