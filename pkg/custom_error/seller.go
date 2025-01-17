package custom_error

import (
	"errors"
	"fmt"
)

type SellerError struct {
	Object     any
	Err        error
	StatusCode int
}

func (s SellerError) Error() string {
	return fmt.Sprintf("error: %v, message: %v", s.Object, s.Err.Error())
}

var (
	ErrorSellerNotFound          error = errors.New("seller not found in the database")
	ErrorCIDSellerAlreadyExist   error = errors.New("seller's CID already exists")
	ErrorMissingSellerID         error = errors.New("missing 'id' parameter in the request")
	ErrorInvalidSellerJSONFormat error = errors.New("invalid JSON format in the request body")
	ErrorNullSellerAttribute     error = errors.New("invalid request body: received empty or null value")
)
