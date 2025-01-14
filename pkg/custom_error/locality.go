package custom_error

import (
	"errors"
	"fmt"
)

type LocalityError struct {
	Object     any
	Err        error
	StatusCode int
}

func (l LocalityError) Error() string {
	return fmt.Sprintf("error: %v, message: %v", l.Object, l.Err.Error())
}

var (
	ErrorLocalityNotFound          error = errors.New("Locality not found in the database")
	ErrorMissingLocalityID         error = errors.New("Missing 'id' parameter in the request")
	ErrorInvalidLocalityJSONFormat error = errors.New("Invalid JSON format in the request body")
	ErrorInvalidLocalityPathParam  error = errors.New("Invalid value for request path parameter")
	ErrorNullLocalityAttribute     error = errors.New("Invalid request body: received empty or null value")
)
