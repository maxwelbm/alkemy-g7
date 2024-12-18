package custom_error

import (
	"fmt"
)

type CustomError struct {
	Object  any
	Message string
}

func (c *CustomError) Error() string {

	return fmt.Sprintf("error: %v, message: %v", c.Object, c.Message)

}
