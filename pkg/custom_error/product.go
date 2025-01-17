package custom_error

import (
	"fmt"
	"net/http"
)

// Só coloquei esse nome para testar, caso mais alguém queira usar esse é a struct
type GenericError struct {
	Message string
	Code    int
	Entity  string
}

const (
	ErrorNotFound = 1
	ErrorConflict = 2
	ErrorInvalid  = 3
	ErrorDep      = 4
	ErrorUnknown  = 5
)

func (b *GenericError) Error() string {
	return fmt.Sprintf("%s %s", b.Entity, b.Message)
}

func NewError(code int, message string, entity string, detail string) error {
	return &GenericError{
		Entity:  entity,
		Code:    code,
		Message: message,
	}
}

func HandleError(entityName string, errorCode int, validationErrors string) error {
	switch errorCode {
	case ErrorNotFound:
		return NewError(http.StatusNotFound, ErrNotFound.Error(), entityName, "")
	case ErrorConflict:
		return NewError(http.StatusConflict, ErrConflict.Error(), entityName, "")
	case ErrorInvalid:
		return NewError(http.StatusUnprocessableEntity, "had errors: "+validationErrors, entityName, "")
	case ErrorDep:
		return NewError(http.StatusConflict, ErrDependencies.Error(), entityName, "")
	default:
		return NewError(http.StatusInternalServerError, ErrUnknow.Error(), entityName, "")
	}
}
