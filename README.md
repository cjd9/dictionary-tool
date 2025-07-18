## Installation

To install the CLI tool to your system, run:

```
make install
```

This will build the binary and copy it to /usr/local/bin (you may be prompted for your password).

---

##  merriam-webster

This project provides a command line tool to query the Merriam-Webster Collegiate Dictionary API and receive a formatted response with the word definition.

- API Reference: [Merriam-Webster Collegiate Dictionary API](https://dictionaryapi.com/products/api-collegiate-dictionary)

## Command Line Tool

The CLI allows users to look up word definitions directly from the terminal.

### Usage

```
$ merriam-webster define <word>
```

For example, given the word `exercise`, your tool might return:

```
`ek-sər-sīz` (noun): the act of bringing into play or realizing in action
```

### Example

```
$ merriam-webster define exercise
`ek-sər-sīz` (noun): the act of bringing into play or realizing in action
```

### Requirements
- You must have a Merriam-Webster API key (see [API documentation](https://dictionaryapi.com/products/api-collegiate-dictionary)).
- The CLI will format the response to show the pronunciation, part of speech, and the first definition.


## CI
- [Buildkite](https://buildkite.com/) pipeline is configured in `buildkite.yml`
