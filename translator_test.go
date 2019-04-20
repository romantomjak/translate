package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(testAPIKey)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func TestTranslatorSubmitsCorrectFormData(t *testing.T) {
	setup()
	defer teardown()

	fromLang := ""
	toLang := "en"
	text := []string{"blahblah"}

	mux.HandleFunc("/language/translate/v2", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		assertEqual(t, r.Form.Get("q"), text[0])
		assertEqual(t, r.Form.Get("target"), toLang)
		assertEqual(t, r.Form.Get("source"), fromLang)
		assertEqual(t, r.Form.Get("format"), "text")
	})

	translator := NewTranslator(client)
	translator.Translate(fromLang, toLang, text)
}
