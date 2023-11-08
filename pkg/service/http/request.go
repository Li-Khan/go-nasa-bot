package http

import (
	"context"
	"net/http"
	"net/url"
)

type Request struct {
	req    *http.Request
	vals   url.Values
	client *http.Client
}

func (r *Request) Do() (*Response, error) {
	return r.do(context.Background())
}

func (r *Request) DoWithContext(ctx context.Context) (*Response, error) {
	return r.do(ctx)
}

func (r *Request) do(ctx context.Context) (*Response, error) {
	r.req.URL.RawQuery = r.vals.Encode()
	resp, err := r.client.Do(r.req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return &Response{
		resp: resp,
	}, nil
}

func (r *Request) SetQueryParam(key string, value string) *Request {
	r.vals.Set(key, value)
	return r
}

func (r *Request) SetHeader(key string, value string) *Request {
	r.req.Header.Set(key, value)
	return r
}

func (r *Request) AddQueryParam(key string, value string) *Request {
	r.vals.Add(key, value)
	return r
}

func (r *Request) AddHeader(key string, value string) *Request {
	r.req.Header.Add(key, value)
	return r
}
