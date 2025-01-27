package customerror

import "fmt"

type CarrierError struct {
	Message string
	Code    int
	Entity  string
}

func (c *CarrierError) Error() string {
	return fmt.Sprintf("%s, %s", c.Entity, c.Message)
}

func NewCarrierError(message string, entity string, code int) error {
	return &CarrierError{
		Entity:  entity,
		Message: message,
		Code:    code,
	}
}
