package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func New() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(s any) error {
	err := v.validator.Struct(s)

	if err == nil {
		return nil
	}

	errs := v.unwrapValidationErr(err)

	return errors.New(strings.Join(errs, ", "))
}

func FromContext(c *gin.Context) *Validator {
	validator, _ := c.Get("validator")

	v, _ := validator.(*Validator)

	return v
}

func (v *Validator) unwrapValidationErr(err error) []string {
	var valErr *validator.ValidationErrors
	if !errors.As(err, &valErr) {
		return []string{err.Error()}
	}

	errs := *valErr

	reasonsWrapper := make([]string, 0, len(errs))

	for _, vErr := range errs {
		reason := fmt.Sprintf("field `%s` does not satisfy %s rule", vErr.Field(), vErr.Tag())

		reasonsWrapper = append(reasonsWrapper, reason)
	}

	return reasonsWrapper
}
