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
    ErrNotFound = 1
    ErrConflict = 2
    ErrInvalid  = 3
    ErrDep      = 4
    ErrUnknown  = 5
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
	case ErrNotFound:
		return NewError(http.StatusNotFound, NotFound.Error(), entityName, "")
	case ErrConflict:
		return NewError(http.StatusConflict, Conflict.Error(), entityName, "")
	case ErrInvalid:
		return NewError(http.StatusUnprocessableEntity, "had errors: "+validationErrors, entityName, "")
	case ErrDep:
		return NewError(http.StatusConflict, DependenciesErr.Error(), entityName, "")
	default:
		return NewError(http.StatusInternalServerError, UnknowErr.Error(), entityName, "")
	}
}