package customerror

import "fmt"

type WareHouseError struct {
	Message string
	Code    int
	Entity  string
}

func (w *WareHouseError) Error() string {
	return fmt.Sprintf("%s %s", w.Entity, w.Message)
}

func NewWareHouseError(message, entity string, code int) error {
	return &WareHouseError{
		Entity:  entity,
		Message: message,
		Code:    code,
	}
}
