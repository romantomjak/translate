package main

import (
	"net/http"
	"net/url"
	"strings"
)

const (
	libraryVersion = "1.0.0"
	userAgent      = "translate/" + libraryVersion + " (+https://github.com/romantomjak/translate)"
)

// Client manages communication with Cloud Translation API
type Client struct {
	// HTTP client used to communicate with CT API
	client *http.Client

	// User agent for client
	UserAgent string

	// Base URL for API requests.
	BaseURL *url.URL
}

// NewClient returns a new Cloud Translation API client
func NewClient() *Client {
	return &Client{
		client:    http.DefaultClient,
		UserAgent: userAgent,
	}
}

// NewRequest returns a HTTP request ready for use with Client.Do
func (c *Client) NewRequest(data url.Values) (*http.Request, error) {
	req, err := http.NewRequest("POST", "https://translation.googleapis.com/language/translate/v2", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

// Do sends an API request and returns the API response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
