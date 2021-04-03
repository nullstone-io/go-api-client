package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Config Config
}

func (c *Client) Do(method string, relativePath string, query url.Values, body io.Reader) (*http.Response, error) {
	req, err := c.CreateRequest(method, relativePath, query, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpClient := &http.Client{
		Transport: c.Config.CreateTransport(http.DefaultTransport),
	}
	return httpClient.Do(req)
}

func (c *Client) CreateRequest(method string, relativePath string, query url.Values, body io.Reader) (*http.Request, error) {
	u, err := c.Config.ConstructUrl(relativePath, query)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(method, u.String(), body)
}
