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
}

// NewClient returns a new Cloud Translation API client
func NewClient() *Client {
	return &Client{
		client:    http.DefaultClient,
		UserAgent: userAgent,
	}
}

// newRequest returns a HTTP request ready for use with Client.Do
func (c *Client) newRequest(data url.Values) (*http.Request, error) {
	req, err := http.NewRequest("POST", "https://translation.googleapis.com/language/translate/v2", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}
