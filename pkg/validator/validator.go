package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	val *validator.Validate
}

func New() *Validator {
	val := validator.New()

	return &Validator{val}
}

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
