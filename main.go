package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	usage = `
Usage: translate [options] [args]

  Translates text from given language to target language using the
  Google Cloud Translation API.

  Translate words into english with automatic source language detection:

      $ translate -to en <word1> <word2>

  Translate words from Russian to Latvian:

      $ translate -from ru -to lv <word>

Options:

  -key=<key>
    The secret key to use to authenticate API requests with.
    Overrides the TRANSLATE_KEY environment variable if set.

  -to=<language>
    The ISO-639-1 language code to use for translation of the input text.
    Overrides the TRANSLATE_TO environment variable if set.

  -from=<language>
    The ISO-639-1 language code of the source text. If the source
    language is not specified, the API will attempt to detect
    the source language automatically.
`
)

func main() {
	os.Exit(Run(os.Stdin, os.Stdout, os.Stdout, os.Args[1:]))
}

func Run(stdin io.Reader, stdout, stderr io.Writer, args []string) int {
	var fromLang, toLang, apiKey string

	flags := flag.NewFlagSet("translate", flag.ContinueOnError)
	flags.StringVar(&fromLang, "from", "", "source language")
	flags.StringVar(&toLang, "to", "", "target language")
	flags.StringVar(&apiKey, "key", "", "secret key")
	flags.Usage = func() {
		fmt.Fprintln(stderr, strings.TrimSpace(usage))
	}

	if err := flags.Parse(args); err != nil {
		return 1
	}

	if apiKey == "" {
		apiKey = os.Getenv("TRANSLATE_KEY")
	}

	if apiKey == "" {
		fmt.Fprintln(stderr, "error: empty secret key")
		return 1
	}

	if toLang == "" {
		toLang = os.Getenv("TRANSLATE_TO")
	}

	if toLang == "" {
		fmt.Fprintln(stderr, "error: empty target language code")
		return 1
	}

	text := flags.Args()
	if len(text) == 0 {
		fmt.Fprintln(stderr, "error: nothing to translate")
		return 1
	}

	client := NewClient(apiKey)

	translations, err := client.Translate(fromLang, toLang, text)
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}

	for _, translation := range translations {
		fmt.Fprintf(stdout, "%s\n", translation.TranslatedText)
	}

	return 0
}
