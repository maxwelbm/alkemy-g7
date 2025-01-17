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
	ErrorLocalityNotFound          error = errors.New("locality Not found")
	ErrorMissingLocalityID         error = errors.New("missing 'id' parameter in the request")
	ErrorInvalidLocalityJSONFormat error = errors.New("invalid JSON format in the request body")
	ErrorInvalidLocalityPathParam  error = errors.New("invalid value for request path parameter")
	ErrorNullLocalityAttribute     error = errors.New("invalid request body: received empty or null value")
)
