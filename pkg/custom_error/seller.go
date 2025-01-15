package custom_error

import "net/http"

type SellerError struct {
	Message string
	Code    int
}

func NewSellerErr(message string, statusCode int) *SellerError {
	return &SellerError{
		Message: message,
		Code:    statusCode,
	}
}

func (e *SellerError) Error() string {
	return e.Message
}

var (
	ErrorSellerNotFound          = NewSellerErr("Seller not found", http.StatusNotFound)
	ErrorCIDSellerAlreadyExist   = NewSellerErr("Seller's CID already exists", http.StatusConflict)
	ErrorMissingSellerID         = NewSellerErr("Missing 'id' parameter in the request", http.StatusBadRequest)
	ErrorInvalidSellerJSONFormat = NewSellerErr("Invalid JSON format in the request body", http.StatusUnprocessableEntity)
	ErrorNullSellerAttribute     = NewSellerErr("Invalid request body: received empty or null value", http.StatusUnprocessableEntity)
	ErrorDefaultSellerSQL        = NewSellerErr("SQL internal error", http.StatusInternalServerError)
)
