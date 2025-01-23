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
	ErrSellerNotFound          = NewSellerErr("seller not found", http.StatusNotFound)
	ErrCIDSellerAlreadyExist   = NewSellerErr("seller's CID already exists", http.StatusConflict)
	ErrMissingSellerID         = NewSellerErr("missing 'id' parameter in the request", http.StatusBadRequest)
	ErrInvalidSellerJSONFormat = NewSellerErr("invalid JSON format in the request body", http.StatusUnprocessableEntity)
	ErrNullSellerAttribute     = NewSellerErr("invalid request body: received empty or null value", http.StatusUnprocessableEntity)
	ErrNotSellerDelete         = NewSellerErr("cannot delete seller, it is necessary to delete locality first.", http.StatusBadRequest)
	ErrDefaultSQLSeller        = NewSellerErr("sql server error", http.StatusInternalServerError)
	ErrInternalServerSeller        = NewSellerErr("internal server error", http.StatusInternalServerError)
)
