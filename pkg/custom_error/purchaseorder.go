package custom_error

import "fmt"

type PurcahseOrderError struct {
	Message string
	Code    int
	Entity  string
}

func (b *PurcahseOrderError) Error() string {
	return fmt.Sprintf("%s %s", b.Entity, b.Message)
}

func NewPurcahseOrderError(code int, message string, entity string) error {
	return &PurcahseOrderError{
		Entity:  entity,
		Code:    code,
		Message: message,
	}
}
