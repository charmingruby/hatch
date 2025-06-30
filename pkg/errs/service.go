package errs

import (
	"fmt"
)

type ResourceAlreadyExistsError struct {
	resource string
}

func (e *ResourceAlreadyExistsError) Error() string {
	return fmt.Sprintf("%s already exists", e.resource)
}

func NewErrResourceAlreadyExists(resource string) error {
	return &ResourceAlreadyExistsError{resource: resource}
}
