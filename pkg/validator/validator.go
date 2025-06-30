// Package validator provides capabilities to validate data structures.
package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator is a wrapper around the go-playground validator instance.
// It is used to validate structs based on struct tags.
type Validator struct {
	val *validator.Validate
}

// New creates and returns a new Validator instance.
//
// Returns:
//   - *Validator: validator instance.
func New() *Validator {
	val := validator.New()

	return &Validator{val}
}

// Validate validates the given struct based on `validate` struct tags.
//
// Parameters:
//   - any: struct to be validated.
//
// Returns:
//   - error: if there is any validation error, with formatted string message.
//
// Example struct:
//
//	type User struct {
//	    Email string `validate:"required,email"`
//	}
//
// Example usage:
//
//	err := v.Validate(user)
//	if err != nil {
//	    log.Println(err)
//	}
func (v *Validator) Validate(obj any) error {
	err := v.val.Struct(obj)
	if err == nil {
		return nil
	}

	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		var messages []string
		for _, e := range validationErrs {
			message := fmt.Sprintf("'%s' failed on '%s'", e.Field(), e.Tag())
			messages = append(messages, message)
		}

		return fmt.Errorf("validation failed: %s", strings.Join(messages, "; "))
	}

	return err
}
