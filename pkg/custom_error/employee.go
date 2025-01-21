package custom_error

import "net/http"

type EmployeerErr struct {
	Message    string
	StatusCode int
}

func (i *EmployeerErr) Error() string {
	return i.Message
}

func NewEmployeerErr(message string, statusCode int) *EmployeerErr {
	return &EmployeerErr{
		Message:    message,
		StatusCode: statusCode,
	}
}

var (
	EmployeeErrNotFound              = NewEmployeerErr("EEmployee not found", http.StatusNotFound)
	EmployeeErrDuplicatedCardNumber  = NewEmployeerErr("duplicated card number id", http.StatusConflict)
	EmployeeErrInvalid               = NewEmployeerErr("invalid employeee", http.StatusUnprocessableEntity)
	EmployeeErrInvalidWarehouseID    = NewEmployeerErr("invalid warehouse id", http.StatusUnprocessableEntity)
	EmployeeErrNotFoundInboundOrders = NewEmployeerErr("inboud orders not found", http.StatusNotFound)
)
