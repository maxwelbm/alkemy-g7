package customError

import "net/http"

type InboundOrderErr struct {
	Message    string
	StatusCode int
}

func (i *InboundOrderErr) Error() string {
	return i.Message
}

func NewInboundOrderErr(message string, statusCode int) *InboundOrderErr {
	return &InboundOrderErr{
		Message:    message,
		StatusCode: statusCode,
	}
}

var (
	InboundErrInvalidEntry          = NewInboundOrderErr("invalid inbound entry", http.StatusUnprocessableEntity)
	InboundErrInvalidEmployee       = NewInboundOrderErr("invalid employee id", http.StatusConflict)
	InboundErrInvalidWarehouse      = NewInboundOrderErr("invalid warehouse id", http.StatusConflict)
	InboundErrInvalidProductBatch   = NewInboundOrderErr("invalid product batch id", http.StatusConflict)
	InboundErrDuplicatedOrderNumber = NewInboundOrderErr("order number already exists", http.StatusConflict)
)
