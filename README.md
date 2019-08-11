# translate [![Build Status](https://travis-ci.org/romantomjak/translate.svg?branch=master)](https://travis-ci.org/romantomjak/translate) [![Coverage Status](https://coveralls.io/repos/github/romantomjak/translate/badge.svg?branch=master)](https://coveralls.io/github/romantomjak/translate?branch=master) [![GoDoc](https://godoc.org/github.com/romantomjak/translate?status.svg)](https://godoc.org/github.com/romantomjak/translate)

Command line client for Google's Cloud Translation API

---

## Requirements

You'll need to sign up for [Cloud Translation](https://cloud.google.com/translate/) to get your API Key.

## Install

```shell
go get -u github.com/romantomjak/translate
```

## Usage

To translate with automatic source language detection:

```shell
$ export TRANSLATE_KEY=xxx
$ export TRANSLATE_TO=en
$ translate kuģis
ship
```

if that doesn't work, you can specify source language manually:

```shell
$ translate -from lv kuģis
ship
```

You can, of course, explicitly override environment values via arguments:

```shell
$ translate -key xxx -to fr kuģis
navire
```

### Translating whole sentences

By default each space separated argument is treated as a word and will get translated on its own line like so:

```shell
$ translate mans kuģis ir visskaistākais
my
ship
and
the most beautiful
```

To translate whole sentences, wrap it in quotes like so:

```shell
$ translate "mans kuģis ir visskaistākais"
my ship is the most beautiful
```

## Contributing

You can contribute in many ways and not just by changing the code! If you have 
any ideas, just open an issue and tell me what you think.

Contributing code-wise - please fork the repository and submit a pull request.

## License

MIT
