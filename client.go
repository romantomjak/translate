package main

import (
	"net/http"
	"net/url"
	"strings"
)

const (
	libraryVersion = "1.0.0"
	defaultBaseURL = "https://translation.googleapis.com"
	userAgent      = "translate/" + libraryVersion + " (+https://github.com/romantomjak/translate)"
)

// Client manages communication with Cloud Translation API
type Client struct {
	// HTTP client used to communicate with CT API
	client *http.Client

	// Google's Cloud Translation API key
	apiKey string

	// User agent for client
	UserAgent string

	// Base URL for API requests.
	BaseURL *url.URL
}

// NewClient returns a new Cloud Translation API client
func NewClient(APIKey string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		client:    http.DefaultClient,
		apiKey:    APIKey,
		UserAgent: userAgent,
		BaseURL:   baseURL,
	}
}

// NewRequest returns a HTTP request ready for use with Client.Do
func (c *Client) NewRequest(urlStr string, data url.Values) (*http.Request, error) {
	data.Add("key", c.apiKey)

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(data.Encode()))
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
