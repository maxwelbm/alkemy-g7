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
	NotFound             = errors.New("Not found")
	Conflict             = errors.New("It already exists")
	EmptyFields          = errors.New("No fields filled")
	InvalidErr           = errors.New("Invalid object")
	NotFoundErrorSection = errors.New("there's no section with this id")
	ConflictErrorSection = errors.New("section with this id already exists")
)
