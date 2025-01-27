package customerror

import "fmt"

type BuyerError struct {
	Message string
	Code    int
	Entity  string
}

func (b *BuyerError) Error() string {
	return fmt.Sprintf("%s %s", b.Entity, b.Message)
}

func NewBuyerError(code int, message string, entity string) error {
	return &BuyerError{
		Entity:  entity,
		Code:    code,
		Message: message,
	}
}
