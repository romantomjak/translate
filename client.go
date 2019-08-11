package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	libraryVersion = "1.0.0"
	defaultBaseURL = "https://translation.googleapis.com/"
	userAgent      = "translate/" + libraryVersion + " (+https://github.com/romantomjak/translate)"
)

// Translation is a translation request result
type Translation struct {
	TranslatedText         string `json:"translatedText"`
	DetectedSourceLanguage string `json:"detectedSourceLanguage"`
}

// Client manages communication with Cloud Translation API
type Client struct {
	// HTTP client used to communicate with CT API
	client *http.Client

	// Google Cloud Translation API key
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

// Translate method translates a string with given parameters
func (c *Client) Translate(fromLang, toLang string, text []string) ([]Translation, error) {
	data := url.Values{
		"key":    {c.apiKey},
		"q":      text,
		"target": {toLang},
		"source": {fromLang},
		"format": {"text"},
	}

	req, err := c.newRequest(http.MethodPost, "language/translate/v2", data)
	if err != nil {
		return nil, fmt.Errorf("cannot initialise request: %v", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot execute request: %v", err)
	}

	var translationResp struct {
		Data struct {
			Translations []Translation `json:"translations"`
		} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&translationResp)
	if err != nil {
		return nil, fmt.Errorf("cannot decode json: %v", err)
	}

	return translationResp.Data.Translations, nil
}

func (c *Client) newRequest(method, path string, data url.Values) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}
