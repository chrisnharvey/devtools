# DevTools

A collection of useful developer utilities in a single command-line application with both CLI and TUI (Terminal User Interface) modes.

> **⚠️ Note**: This project is currently in active development. Features and APIs may change between versions.

## Features

DevTools provides the following utilities:

- **SHA256** - Generate SHA256 hashes from input strings
- **URL Encode** - URL encode strings for safe transmission in URLs
- **URL Decode** - Decode URL-encoded strings back to their original form
- **JSON Minify** - Remove whitespace and formatting from JSON to reduce size
- **JSON Prettify** - Format JSON with proper indentation for readability

## Installation

### From Source

```bash
git clone https://github.com/chrisnharvey/devtools.git
cd devtools
go build -o devtools .
```

### Using Go Install

```bash
go install github.com/chrisnharvey/devtools@latest
```

## Usage

DevTools can be used in two modes:

### TUI Mode (Interactive)

Run the application without any arguments to launch the interactive Terminal User Interface:

```bash
./devtools
```

This will display a menu where you can select the tool you want to use and fill in the required fields interactively.

### CLI Mode (Command Line)

Each tool can also be used directly from the command line:

```bash
# Generate SHA256 hash
./devtools sha256 --string "hello world"

# URL encode a string
./devtools urlencode --string "hello world!"

# URL decode a string
./devtools urldecode --string "hello%20world%21"

# Minify JSON
./devtools jsonminify --json '{"name": "John", "age": 30}'

# Prettify JSON
./devtools jsonprettify --json '{"name":"John","age":30}'
```

### Help

Get help for any command:

```bash
./devtools --help
./devtools sha256 --help
```

## Examples

### SHA256 Hash Generation
```bash
$ ./devtools sha256 --string "hello world"
b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
```

### URL Encoding
```bash
$ ./devtools urlencode --string "hello world!"
hello%20world%21
```

### URL Decoding
```bash
$ ./devtools urldecode --string "hello%20world%21"
hello world!
```

### JSON Minification
```bash
$ ./devtools jsonminify --json '{
  "name": "John",
  "age": 30,
  "city": "New York"
}'
{"name":"John","age":30,"city":"New York"}
```

### JSON Prettification
```bash
$ ./devtools jsonprettify --json '{"name":"John","age":30,"city":"New York"}'
{
  "name": "John",
  "age": 30,
  "city": "New York"
}
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [tview](https://github.com/rivo/tview) - Terminal UI library

## Development

### Building

```bash
go build -o devtools .
```

### Testing

```bash
go test ./...
```

### Adding New Tools

To add a new tool:

1. Create a new package in the `cmd/` directory
2. Implement the `Command` interface with:
   - `Execute(values field.Values) error`
   - `GetName() string`
   - `GetDescription() string`
   - `GetFields() []field.Field`
3. Register the command in `main.go`

## License

This project is licensed under the terms specified in the LICENSE file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
