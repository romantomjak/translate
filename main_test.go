package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func assertContains(t *testing.T, got, want string) {
	t.Helper()
	if !strings.Contains(got, want) {
		t.Fatalf("expected %q to contain %q, but it didn't", got, want)
	}
}

func TestRun_ApiKeyNotSet(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	args := []string{}

	code := Run(stdin, stdout, stderr, args)
	assertEqual(t, code, 1)

	out := stderr.String()
	assertContains(t, out, "empty secret key")
}

func TestRun_TargetLanguageNotSet(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	args := []string{"-key=test"}

	code := Run(stdin, stdout, stderr, args)
	assertEqual(t, code, 1)

	out := stderr.String()
	assertContains(t, out, "empty target language code")
}

func TestRun_TextNotSet(t *testing.T) {
	stdin := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	args := []string{"-key=test", "-to=en"}

	code := Run(stdin, stdout, stderr, args)
	assertEqual(t, code, 1)

	out := stderr.String()
	assertContains(t, out, "nothing to translate")
}
