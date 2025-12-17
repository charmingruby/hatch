package errs

import "fmt"

type NotFoundError struct {
	message string
}

func NewNotFoundError(resource string) error {
	return &NotFoundError{
		message: fmt.Sprintf("%s not found", resource),
	}
}

func (e *NotFoundError) Error() string {
	return e.message
}
