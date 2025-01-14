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
	ErrorSellerNotFound          error = errors.New("Seller not found in the database")
	ErrorCIDSellerAlreadyExist   error = errors.New("Seller's CID already exists")
	ErrorMissingSellerID         error = errors.New("Missing 'id' parameter in the request")
	ErrorInvalidSellerJSONFormat error = errors.New("Invalid JSON format in the request body")
	ErrorNullSellerAttribute     error = errors.New("Invalid request body: received empty or null value")
)
