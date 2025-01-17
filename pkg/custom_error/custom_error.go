package custom_error

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Object     any
	Err        error
	StatusCode int
}

func (c CustomError) Error() string {
	return fmt.Sprintf("error: %v, message: %v", c.Object, c.Err.Error())
}

var (
	ErrNotFound             = errors.New("not found")
	ErrConflict             = errors.New("it already exists")
	ErrEmptyFields          = errors.New("no fields filled")
	ErrInvalid              = errors.New("invalid object")
	ErrDependencies         = errors.New("cannot be deleted because there are dependencies")
	ErrNotFoundErrorSection = errors.New("there's no section with this id")
	ErrConflictSection      = errors.New("section with this id already exists")
	ErrUnknow               = errors.New("unknow server error")
)
