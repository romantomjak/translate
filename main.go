package main

import (
	"flag"
	"fmt"
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
	Overrides the RT_TRANSLATE_KEY environment variable if set.

  -to=<language>
    The ISO-639-1 language code to use for translation of the input text.
	Overrides the RT_TRANSLATE_TO environment variable if set.

  -from=<language>
	The ISO-639-1 language code of the source text. If the source
	language is not specified, the API will attempt to detect
	the source language automatically.
`
)

func main() {
	var fromLang, toLang, apiKey string

	flag.StringVar(&fromLang, "from", "", "source language")
	flag.StringVar(&toLang, "to", "", "target language")
	flag.StringVar(&apiKey, "key", "", "secret key")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, strings.TrimSpace(usage))
	}
	flag.Parse()

	text := flag.Args()

	if toLang == "" {
		fmt.Fprintln(os.Stderr, "error: empty target language code")
		return
	}

	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "error: empty secret key")
		return
	}

	if len(text) == 0 {
		fmt.Fprintln(os.Stderr, "error: nothing to translate")
		return
	}

	client := NewClient(apiKey)
	translator := NewTranslator(client)

	translations, err := translator.Translate(fromLang, toLang, text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	for _, translation := range translations {
		fmt.Fprintf(os.Stdout, "%s\n", translation.TranslatedText)
	}
}
