package translate

import (
	"net/url"
	"testing"
)

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestNewPreparedRequest(t *testing.T) {
	client := NewClient()

	req, _ := client.NewRequest(url.Values{"hello": {"world"}})
	req.ParseForm()

	assertEqual(t, req.Method, "POST")
	assertEqual(t, req.Header.Get("User-Agent")[:9], "translate")
	assertEqual(t, req.Form.Get("hello"), "world")
}
