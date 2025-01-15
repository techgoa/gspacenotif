// Package http provides internal HTTP client functionality
// for sending message to Google Spaces.
package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Client wraps standard http.Client for sending messages.
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new HTTP client instance.
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// SendMessage sends a POST request to the specified webhook URL
// with the provided request body.
func (c *Client) SendMessage(webhookURL string, request string) error {
	r := strings.NewReader(request)
	req, err := http.NewRequest(http.MethodPost, webhookURL, r)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
