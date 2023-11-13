package gonvista

import "net/http"

// NewClientDefault creates an API client using the Go default HTTP client
func NewClientDefault() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}

// NewClient creates an API client using a custom created Go HTTP client
func NewClient(c *http.Client) *Client {
	return &Client{
		httpClient: c,
	}
}
