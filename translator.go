package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Translator represents the Cloud Translation service
type Translator struct {
	Client *Client
}

// Translation is a translation request result
type Translation struct {
	TranslatedText         string `json:"translatedText"`
	DetectedSourceLanguage string `json:"detectedSourceLanguage"`
}

// NewTranslator returns a new Cloud Translation client
func NewTranslator(client *Client) *Translator {
	return &Translator{
		Client: client,
	}
}

// Translate method translates a string with given parameters
func (t *Translator) Translate(fromLang, toLang string, text []string) ([]Translation, error) {
	data := url.Values{
		"q":      text,
		"target": {toLang},
		"source": {fromLang},
		"format": {"text"},
	}

	req, err := t.Client.NewRequest(data)
	if err != nil {
		return nil, fmt.Errorf("cannot initialise request: %v", err)
	}

	resp, err := t.Client.Do(req)
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
