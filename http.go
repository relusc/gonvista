package gonvista

import (
	"fmt"
	"io"
	"net/http"
)

// doHTTP executes HTTP requests with the submitted method and returns the response body
func (c *Client) doHTTP(url, method string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating HTTP request: %s", err)
	}

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while executing HTTP request: %s", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read HTTP response body: %s", err)
	}

	return bodyBytes, err
}
