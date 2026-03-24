package httptest

import (
	"HATCH_APP/pkg/o11y"
	"HATCH_APP/pkg/transport/httpx"
	"HATCH_APP/pkg/validator"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	CheckResponse  func(t *testing.T, body []byte)
	ArrangeRequest func() *http.Request
	Headers        map[string]string
	ExpectStatus   int
}

func Run(t *testing.T, handler http.HandlerFunc, tc Case) {
	t.Helper()

	var req *http.Request

	if tc.ArrangeRequest != nil {
		req = tc.ArrangeRequest()
	}

	if req == nil {
		req = httptest.NewRequest(http.MethodGet, "/", nil)
	}

	req = injectContext(req)

	for k, v := range tc.Headers {
		req.Header.Set(k, v)
	}

	rec := httptest.NewRecorder()

	handler(rec, req)

	assert.Equal(t, tc.ExpectStatus, rec.Code)

	if tc.CheckResponse != nil {
		tc.CheckResponse(t, rec.Body.Bytes())
	}
}

func Init() {
	o11y.Init()
}

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

func WithParam(req *http.Request, k, v string) *http.Request {
	req.SetPathValue(k, v)

	return req
}

func NewRequest(method, target string) *http.Request {
	return httptest.NewRequest(method, target, nil)
}

func injectContext(req *http.Request) *http.Request {
	ctx := context.Background()

	ctx = o11y.WithLogger(ctx, o11y.Log)

	v := validator.New()

	ctx = validator.WithValidator(ctx, v)

	return req.WithContext(ctx)
}
