package testutil

import (
	"HATCH_APP/pkg/o11y"
	"HATCH_APP/pkg/transport/httpx"
	"HATCH_APP/pkg/validator"
	"context"
	"encoding/json"
	"net/http"
)

func ParseResponse[T any](body []byte) (T, httpx.Response, error) {
	var raw httpx.Response
	err := json.Unmarshal(body, &raw)
	if err != nil {
		var zero T
		return zero, httpx.Response{}, err
	}

	data, err := json.Marshal(raw.Data)
	if err != nil {
		var zero T
		return zero, raw, err
	}

	var typed T
	err = json.Unmarshal(data, &typed)

	return typed, raw, err
}

func RequestInjection(req *http.Request) *http.Request {
	ctx := context.Background()

	ctx = o11y.WithLogger(ctx, o11y.Log)

	v := validator.New()

	ctx = validator.WithValidator(ctx, v)

	return req.WithContext(ctx)
}

func InjectPathParam(req *http.Request, k, v string) *http.Request {
	req.SetPathValue(k, v)

	return req
}
