package main

import (
	"net/http"
	"net/url"
)

// Translator represents the Cloud Translation service
type Translator struct {
	Client *Client
}

// NewTranslator returns a new Cloud Translation client
func NewTranslator(client *Client) *Translator {
	return &Translator{
		Client: client,
	}
}

// Translate method translates a string with given parameters
func (t *Translator) Translate() (*http.Response, error) {
	data := url.Values{
		"q":      {"stabs"},
		"target": {"en"},
		"format": {"text"},
		"key":    {"blah"},
	}

	req, err := t.Client.NewRequest(data)
	if err != nil {
		return nil, err
	}

	resp, err := t.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
