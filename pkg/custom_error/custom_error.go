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
	NotFound        = errors.New("Not found")
	Conflict        = errors.New("It already exists")
	EmptyFields     = errors.New("No fields filled")
	InvalidErr      = errors.New("Invalid object")
	DependenciesErr = errors.New("Cannot be deleted because there are dependencies")
)
