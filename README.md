# mdcat

A terminal markdown renderer that makes your markdown files look beautiful in the terminal. Built specifically with iTerm2 users in mind.

![image](https://github.com/user-attachments/assets/955b80dc-a618-4515-84af-f1c714ad1377)

## Features

- 🎨 Syntax highlighting for code blocks
- 📱 Responsive terminal width handling
- 🔄 Pipe support (works with `cat` and other tools)
- 🎯 iTerm2 optimized styling
- 🚀 Fast rendering
- 📖 Support for all standard markdown features

## Installation

### Using Go

```bash
go install github.com/domano/mdcat@latest
```

### From Source

```bash
git clone https://github.com/domano/mdcat
cd mdcat
go build -o mdcat
mv mdcat /usr/local/bin/
```

## Usage

```bash
# Read from file
mdcat README.md

# Read from stdin
cat README.md | mdcat

# Read multiple files
mdcat README.md CONTRIBUTING.md
```

## Example

Input markdown:
````markdown
# Hello World

This is a **bold** statement with some `code`.

```go
fmt.Println("Hello, World!")
```
````

Will be rendered with proper styling in your terminal!

## Development

Requirements:
- Go 1.19 or later
- iTerm2 (recommended)

```bash
# Clone the repo
git clone https://github.com/domano/mdcat
cd mdcat

# Install dependencies
go mod download

# Build
go build

# Run tests
go test ./...
```

## Contributing

Contributions are welcome! Feel free to:
1. Fork the repository
2. Create a feature branch
3. Submit a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details

## Credits

Built with:
- [glamour](https://github.com/charmbracelet/glamour) - Markdown rendering
- [cobra](https://github.com/spf13/cobra) - CLI framework

## Author

Dino Omanovic ([@domano](https://github.com/domano))
