package errs

import (
	"fmt"
)

// ResourceAlreadyExistsError is for cases that tries create a resoure that already exists.
type ResourceAlreadyExistsError struct {
	resource string
}

func (e *ResourceAlreadyExistsError) Error() string {
	return fmt.Sprintf("%s already exists", e.resource)
}

// NewResourceAlreadyExistsError creates the ResourceAlreadyExistsError.
//
// Parameters:
//   - string: resource conflicting (e.g.:"devices", "users").
//
// Returns:
//   - error: an error if formatted message.
func NewResourceAlreadyExistsError(resource string) error {
	return &ResourceAlreadyExistsError{resource: resource}
}
