package customError

import (
	"net/http"
)

type LocalityError struct {
	Message string
	Code    int
}

func NewLocalityErr(message string, statusCode int) *LocalityError {
	return &LocalityError{
		Message: message,
		Code:    statusCode,
	}
}

func (e *LocalityError) Error() string {
	return e.Message
}

var (
	ErrLocalityNotFound          = NewLocalityErr("locality not found", http.StatusNotFound)
	ErrMissingLocalityID         = NewLocalityErr("missing 'id' parameter in the request", http.StatusBadRequest)
	ErrInvalidLocalityJSONFormat = NewLocalityErr("invalid JSON format in the request body", http.StatusBadRequest)
	ErrInvalidLocalityPathParam  = NewLocalityErr("invalid value for request path parameter", http.StatusUnprocessableEntity)
	ErrNullLocalityAttribute     = NewLocalityErr("invalid request body: received empty or null value", http.StatusUnprocessableEntity)
	ErrDefaultLocalitySQL        = NewLocalityErr("sql internal error", http.StatusInternalServerError)
	ErrDefaultLocality           = NewSellerErr("internal server error", http.StatusInternalServerError)
)
