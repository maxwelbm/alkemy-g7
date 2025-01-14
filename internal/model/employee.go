package model

type Employee struct {
	Id           int
	CardNumberId string
	FirstName    string
	LastName     string
	WarehouseId  int
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
