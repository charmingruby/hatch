package httpx

import (
	"HATCH_APP/pkg/validator"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrInvalidPayload = errors.New("invalid payload")

func ParseRequest[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	var obj T

	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		WriteBadRequestResponse(w, fmt.Sprintf("%s: %s", ErrInvalidPayload.Error(), err.Error()))
		return nil, err
	}

	val := validator.FromContext(r.Context())

	if err := val.Validate(obj); err != nil {
		WriteBadRequestResponse(w, fmt.Sprintf("%s: %s", ErrInvalidPayload.Error(), err.Error()))
		return nil, err
	}

	return &obj, nil
}
