package custom_error

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
	ErrorLocalityNotFound          = NewLocalityErr("Locality not found", http.StatusNotFound)
	ErrorMissingLocalityID         = NewLocalityErr("Missing 'id' parameter in the request", http.StatusBadRequest)
	ErrorInvalidLocalityJSONFormat = NewLocalityErr("Invalid JSON format in the request body", http.StatusBadRequest)
	ErrorInvalidLocalityPathParam  = NewLocalityErr("Invalid value for request path parameter", http.StatusUnprocessableEntity)
	ErrorNullLocalityAttribute     = NewLocalityErr("Invalid request body: received empty or null value", http.StatusUnprocessableEntity)
	ErrorDefaultLocalitySQL        = NewLocalityErr("SQL internal error", http.StatusInternalServerError)
)
