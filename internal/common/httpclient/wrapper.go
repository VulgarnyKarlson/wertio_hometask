package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

type Wrapper struct {
	domainURL string
	headers   map[string]string
}

func NewWrapper(domainURL string, headers map[string]string) *Wrapper {
	return &Wrapper{
		domainURL: domainURL,
		headers:   headers,
	}
}

func (c *Wrapper) NewRequest(ctx context.Context, opts *Request) (*Response, error) {
	client := &http.Client{}
	url := c.domainURL + "/" + opts.Path

	req, err := http.NewRequestWithContext(ctx, opts.Method, url, bytes.NewReader(opts.Body))
	if err != nil {
		return nil, fmt.Errorf("error create request %w", err)
	}

	q := req.URL.Query()
	for key, value := range opts.Query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	for key, value := range c.headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error send request %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error read response body %w", err)
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
	}, nil
}
