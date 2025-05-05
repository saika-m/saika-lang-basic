# Saika Programming Language (赛卡编程语言)

Saika is a programming language that lets you write Go code using Chinese keywords and syntax. It's designed to make programming more accessible to Chinese speakers while leveraging the power and performance of the Go ecosystem.

![Saika Logo](https://via.placeholder.com/200x100?text=Saika)

## Features

- **Native Chinese syntax**: Write code using Chinese keywords, making it more intuitive for Chinese speakers
- **Go compatibility**: Transpiles to standard Go code that leverages the entire Go ecosystem
- **VS Code support**: Includes syntax highlighting and language tools
- **Simple toolchain**: Build and run Saika files with easy-to-use commands
- **Error messages in Chinese**: Get helpful error messages in Chinese when debugging

## Installation

### Prerequisites

- Go 1.18 or later
- Git

### Installing Saika

```bash
# Clone the repository
git clone https://github.com/saika-m/saika-lang-basic.git

# Build the Saika compiler
cd saika-lang-basic
go build -o saika ./cmd/saika

# Add to your PATH (optional)
# On Linux/macOS:
sudo mv saika /usr/local/bin/
# On Windows: Add the directory containing saika.exe to your PATH
```

### Installing the VS Code Extension

1. Open VS Code
2. Go to Extensions (Ctrl+Shift+X)
3. Click on "..." in the Extensions panel and select "Install from VSIX..."
4. Navigate to the `saika-extension` folder and select the VSIX file (or install from the marketplace once published)

## Quick Start

Create your first Saika program:

```saika
包 入口

导入 "工具"

数 入口() {
    工具.打印行("你好，世界！")
}
```

Save this as `hello.saika` and run it:

```bash
saika run hello.saika
```

## Commands

- `saika build <file.saika>`: Compile a Saika file to an executable
- `saika run <file.saika>`: Run a Saika program
- `saika test <file.saika>`: Show the transpiled Go code

## Language Reference

### Basic Structure

| Saika (Chinese) | Go (English) |
|-----------------|--------------|
| 包              | package      |
| 导入             | import       |
| 数 / 函数         | func         |
| 返回             | return       |

### Data Types

| Saika (Chinese) | Go (English) |
|-----------------|--------------|
| 整数             | int          |
| 浮点数            | float64      |
| 字符串            | string       |
| 布尔             | bool         |
| 真              | true         |
| 假              | false        |
| 空 / 无          | nil          |

### Control Flow

| Saika (Chinese) | Go (English) |
|-----------------|--------------|
| 如果             | if           |
| 否则             | else         |
| 循环             | for          |
| 范围             | range        |

For a complete language reference, see the [documentation](docs/language-reference.md).

## Examples

Check out the `examples/` directory for sample Saika programs:

- [hello.saika](examples/hello.saika): Basic "Hello World" and simple functions
- [simple.saika](examples/simple.saika): Minimal example
- [advanced.saika](examples/advanced.saika): Demonstrates more complex features like structs, interfaces, and concurrency

## Project Structure

```
saika-lang-basic/
├── cmd/saika/               # Command-line interface
├── examples/                # Example Saika programs
├── internal/                # Internal packages
│   └── transpiler/          # Transpiler implementation
├── lib/                     # Library packages
│   └── 工具/                 # Chinese wrappers for Go packages
└── saika-extension/         # VS Code extension
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- The Go team for creating such an excellent language
- Everyone who has contributed to making programming more accessible to non-English speakers

## Contact

Project Link: [https://github.com/saika-m/saika-lang-basic](https://github.com/saika-m/saika-lang-basic)