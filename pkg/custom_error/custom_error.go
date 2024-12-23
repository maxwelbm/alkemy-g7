package custom_error

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Object any
	Err    error
}

func (c CustomError) Error() string {
	return fmt.Sprintf("error: %v, message: %v", c.Object, c.Err.Error())
}

var (
	NotFound      = errors.New("not found")
	Conflict      = errors.New("It already exists")
	EmptyFields   = errors.New("No fields filled")
	InvalidErr    = errors.New("invalid object")
	AlreadyExists = errors.New("already exists")
)