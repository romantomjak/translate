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

	client = NewClient()
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func assertHttpMethod(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestTranslatorSubmitsCorrectFormData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/language/translate/v2", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		assertHttpMethod(t, r.Method, "POST")
		assertEqual(t, r.Header.Get("User-Agent")[:9], "translate")
		assertEqual(t, r.Form.Get("q"), "stabs")
		assertEqual(t, r.Form.Get("target"), "en")
		assertEqual(t, r.Form.Get("format"), "text")
		assertEqual(t, r.Form.Get("key"), "blahllll")
	})

	translator := NewTranslator(client)
	translator.Translate()
}
