# Saika Language Support for VS Code

This extension provides syntax highlighting and language support for the Saika programming language, a Chinese programming language built on top of Go.

## Features

- Syntax highlighting for Saika (.saika) files
- Recognition of Chinese keywords, types, and functions
- Proper handling of comments, strings, and other language constructs

## Installation

### From VS Code Marketplace (Coming Soon)
1. Open VS Code
2. Go to the Extensions view (Ctrl+Shift+X)
3. Search for "Saika Language Support"
4. Click Install

### Manual Installation
1. Download or clone this repository
2. Copy the `language-saika` folder to your VS Code extensions folder:
   - Windows: `%USERPROFILE%\.vscode\extensions`
   - macOS/Linux: `~/.vscode/extensions`
3. Restart VS Code

## Example

```saika
包 入口

导入 "格式化"

數 入口() {
    格式化.打印行("你好，世界！")
}
```

## Supported Keywords

This extension supports all Chinese keywords used in Saika, including:

### Basic Structure
- `包` (package)
- `导入` (import)
- `數`, `函数` (func)
- `返回` (return)

### Data Types
- `整数` (int)
- `浮点数` (float64)
- `字符串` (string)
- `布尔` (bool)
- `真` (true)
- `假` (false)
- `空` (nil)

### Control Flow
- `如果` (if)
- `否则` (else)
- `循环` (for)
- `范围` (range)

... and many more.

## Development

This extension is still in early development. If you encounter any issues or have suggestions for improvements, please submit an issue or pull request on GitHub.

## License

MIT