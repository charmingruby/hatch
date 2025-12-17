package rest

import (
	"HATCH_APP/pkg/validator"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var ErrInvalidPayload = errors.New("invalid payload")

func ParseRequest[T any](c *gin.Context) (*T, error) {
	var obj T

	if err := c.ShouldBindJSON(&obj); err != nil {
		return nil, fmt.Errorf("%s: %s", ErrInvalidPayload.Error(), err.Error())
	}

	val := validator.FromContext(c)

	if err := val.Validate(obj); err != nil {
		return nil, fmt.Errorf("%s: %s", ErrInvalidPayload.Error(), err.Error())
	}

	return &obj, nil
}
