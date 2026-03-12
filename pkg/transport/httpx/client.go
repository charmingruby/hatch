package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const defaultRequestTimeout = 5 * time.Second

type HTTPMethod = string

const (
	HTTPMethodGet    HTTPMethod = http.MethodGet
	HTTPMethodPost   HTTPMethod = http.MethodPost
	HTTPMethodPut    HTTPMethod = http.MethodPut
	HTTPMethodPatch  HTTPMethod = http.MethodPatch
	HTTPMethodDelete HTTPMethod = http.MethodDelete
)

var (
	ErrNilResponse   = errors.New("nil response")
	ErrNilTarget     = errors.New("nil target")
	ErrParseJSONBody = errors.New("parse json")
)

type Client struct {
	client  *http.Client
	baseURL string
}

type Request struct {
	URL     string
	Path    string
	Method  HTTPMethod
	Headers http.Header
	Body    []byte
}

func NewClient(baseURL string, timeout time.Duration) *Client {
	to := defaultRequestTimeout
	if timeout > 0 {
		to = timeout
	}

	client := http.Client{
		Timeout: to,
	}

	return &Client{
		client:  &client,
		baseURL: baseURL,
	}
}

func (c *Client) Do(ctx context.Context, req Request) (*http.Response, error) {
	method := HTTPMethodGet
	if req.Method != "" {
		method = req.Method
	}

	var reader io.Reader
	if len(req.Body) > 0 {
		reader = bytes.NewReader(req.Body)
	}

	baseURL := c.baseURL
	if req.URL != "" {
		baseURL = req.URL
	}

	reqURLStr := strings.TrimRight(baseURL, "/")
	if trimmedPath := strings.TrimLeft(req.Path, "/"); trimmedPath != "" {
		reqURLStr += "/" + trimmedPath
	}

	reqURL, err := url.Parse(reqURLStr)
	if err != nil {
		return nil, fmt.Errorf("invalid request url: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, method, reqURL.String(), reader)
	if err != nil {
		return nil, fmt.Errorf("invalid request url: %w", err)
	}

	if len(req.Headers) > 0 {
		httpReq.Header = req.Headers.Clone()
	}

	res, err := c.client.Do(httpReq) // #nosec G704 -- request URL already validated
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}

	return res, nil
}

func ParseResponse[T any](ctx context.Context, res *http.Response, target *T) error {
	if res == nil {
		return ErrNilResponse
	}

	if target == nil {
		return ErrNilTarget
	}

	if ctx == nil {
		ctx = context.Background()
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return fmt.Errorf("%w: %w", ErrParseJSONBody, err)
	}

	return nil
}
