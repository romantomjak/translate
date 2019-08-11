# translate

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
$ translate -key xxx -to en kuģis
ship
```

or if that doesn't work, specify language manually:

```shell
$ translate -key xxx -from lv -to en kuģis
ship
```

but manually specifying `key` and `to` parameters gets tedious, so lets set those values via environment variables:

```shell
export TRANSLATE_KEY=xxx
export TRANSLATE_TO=en
```

then you can execute the binary with a single argument:

```
$ translate kuģis
ship
```

You can, of course, explicitly override environment values via arguments:

```shell
$ translate -to fr kuģis
navire
```

## Contributing

You can contribute in many ways and not just by changing the code! If you have 
any ideas, just open an issue and tell me what you think.

Contributing code-wise - please fork the repository and submit a pull request.

## License

MIT
