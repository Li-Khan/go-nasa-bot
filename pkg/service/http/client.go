package http

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	netUrl "net/url"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}

func (c *Client) SetHttpClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) Request(method string, url string) (*Request, error) {
	return c.request(method, url, nil)
}

func (c *Client) RequestJSON(method string, url string, payload any) (*Request, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.request(method, url, bytes.NewBuffer(b))
}

func (c *Client) RequestXML(method string, url string, payload any) (*Request, error) {
	b, err := xml.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.request(method, url, bytes.NewBuffer(b))
}

func (c *Client) request(method string, url string, body io.Reader) (*Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return &Request{
		req:    req,
		client: c.httpClient,
		vals:   make(netUrl.Values),
	}, nil
}
