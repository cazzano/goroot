# 🌲 GoRoot

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/goroot)](https://goreportcard.com/report/github.com/cazzano/goroot)
[![Go Version](https://img.shields.io/github/go-mod/go-version/yourusername/goroot)](https://github.com/cazzano/goroot)
[![Release](https://img.shields.io/github/v1.0.0/yourusername/goroot)](https://github.com/cazzano/goroot/releases)

**GoRoot** is a powerful command-line tool designed to simplify Go project management and execution.

## ✨ Features

- **Project Initialization** - Set up new Go projects with proper structure
- **Build Automation** - Compile your Go projects with ease
- **Execution Control** - Run Go files directly with support for arguments
- **Flexible Usage** - Target specific files or modules for execution

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/cazzano/goroot.git

# Navigate to the directory
cd goroot/stable/go/src/

# Install the tool
go build && mv hello goroot && sudo mv goroot /usr/bin/
```

## 🚀 Usage

GoRoot provides several commands to simplify your Go development workflow:

### Initialize a new project

```bash
goroot init
```

This will create a standard Go project structure in the current directory:

```
my-project/
├── cmd/
│   └── main.go
├── internal/
├── pkg/
├── go.mod
└── README.md
```

### Build your project

```bash
goroot build
```

Compiles your Go project and produces an executable binary.

### Run Go files

Run a Go file in the current directory:

```bash
goroot run
```

Run with arguments (supports up to 10 arguments):

```bash
goroot run arg1 arg2 arg3
```

### Run a specific file or module

```bash
goroot run --1 ./path/to/file.go
```

```bash
goroot run --1 ./cmd/mymodule
```

### Display version information

```bash
goroot --v
```

### Display help message

```bash
goroot --h
```

## 📝 Examples

### Example 1: Quick Start Project

```bash
# Initialize a new project
mkdir my-awesome-app
cd my-awesome-app
goroot init

# Build the project
goroot build

# Run the project
goroot run
```

### Example 2: Running with Arguments

```bash
# Run a file with command-line arguments
goroot run config.json --verbose debug
```

### Example 3: Working with Specific Files

```bash
# Run a specific file
goroot run --1 ./examples/hello.go

# Build and run a specific module
goroot build
goroot run --1 ./cmd/api
```

## 🛠️ Command Reference

| Command | Description |
|---------|-------------|
| `init` | Initialize the project structure |
| `build` | Build the project |
| `run` | Run Go files in the current directory |
| `run --1` | Run a specific file or module |
| `--v` | Display version information |
| `--h` | Display help message |

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📞 Support

For support, please open an issue in the GitHub repository or contact the maintainers.
