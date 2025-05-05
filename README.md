# Saika 编程语言 (Saika Programming Language)

Saika is a programming language that allows you to write Go code using Chinese keywords and syntax, making programming more accessible to Chinese speakers. The Saika transpiler converts your Chinese code into standard Go, enabling you to harness Go's performance and ecosystem while coding in Chinese.

![Saika](https://avatars.githubusercontent.com/u/165717332?v=4)

## 特性 (Features)

- **Chinese Keywords**: Write code using Chinese keywords for all Go constructs
- **Go Compatibility**: Fully compatible with the Go ecosystem and libraries
- **Simple Transpiler**: Easy conversion between Saika and Go code
- **Comprehensive Standard Library**: Access Go packages with Chinese names
- **Error Messages in Chinese**: Receive error feedback in Chinese
- **Zero Runtime Overhead**: Compiles directly to native executables with Go's performance

## 安装 (Installation)

### VSCode Extension
### <span style="color: yellow;">saika-m.saika-lang</span>

### 预构建二进制文件 (Pre-built Binaries)

Download the appropriate binary for your platform from the [Releases](https://github.com/saika-m/saika-lang-basic/releases) page.

#### Binary Types

| Archive Filename | System Type | Common Uses |
|-----------------|-------------|-------------|
| saika-linux-amd64.tar.gz | Linux on Intel/AMD 64-bit (x86_64) | Standard Linux desktops, servers |
| saika-linux-arm64.tar.gz | Linux on ARM 64-bit | Raspberry Pi 4, ARM servers, embedded devices |
| saika-linux-386.tar.gz | Linux on 32-bit Intel/AMD (x86) | Older Linux systems |
| saika-windows-amd64.zip | Windows on Intel/AMD 64-bit (x86_64) | Standard Windows PCs, laptops |
| saika-windows-arm64.zip | Windows on ARM 64-bit | Surface Pro X, Windows ARM devices |
| saika-windows-386.zip | Windows on 32-bit Intel/AMD (x86) | Older Windows systems |
| saika-macos-amd64.tar.gz | macOS on Intel processors | Intel Macs (pre-2020) |
| saika-macos-arm64.tar.gz | macOS on Apple Silicon | M1/M2/M3 Macs (2020 onward) |

#### Installation Instructions

Each archive includes an installation script that makes installation easy. After downloading, simply extract the archive and run the included installation script.

#### Manual Installation (Alternative)

[If you prefer to install manually, follow these platform-specific instructions](/binary-help/manual.md)

### 从源码构建 (Build from Source)

Prerequisites:
- Go 1.24 or later

```bash
# Clone the repository
git clone https://github.com/saika-m/saika-lang-basic.git
cd saika-lang-basic

# Build the saika binary
go build -o saika ./cmd/saika

# Install the binary to your GOPATH (optional)
go install ./cmd/saika
```

## 使用方法 (Usage)

Saika provides three main commands:

### 构建 (Build)

Compiles a Saika file to an executable:

```bash
saika build your_file.saika
```

This will create an executable with the same name as your Saika file (without the `.saika` extension).

### 运行 (Run)

Transpiles and runs a Saika file directly:

```bash
saika run your_file.saika
```

This is useful during development and testing.

### 测试 (Test)

Shows the transpiled Go code for a Saika file, which is helpful for learning and debugging:

```bash
saika test your_file.saika
```

The command will display both your original Saika code and the transpiled Go code side by side. It also gives you the option to save the Go code to a file.

## 代码示例 (Code Examples)

### 简单的Hello World (Simple Hello World)

```go
包 入口

导入 "工具"

数 入口() {
    工具.打印行("世界")
}
```

This will print "世界" (a play on "Hello World").

### 高级示例 (Advanced Example)

For more advanced features like structs, interfaces, concurrency, and error handling, see the `examples/advanced.saika` file in the repository.

## 关键字对照表 (Keyword Reference)

Here are the main Saika keywords and their Go equivalents:

| Saika (Chinese) | Go (English) |
|-----------------|--------------|
| 包              | package      |
| 导入             | import       |
| 数/函数          | func         |
| 变量             | var          |
| 常量             | const        |
| 返回             | return       |
| 整数             | int          |
| 浮点数           | float64      |
| 字符串           | string       |
| 布尔             | bool         |
| 真              | true         |
| 假              | false        |
| 如果             | if           |
| 否则             | else         |
| 否则如果          | else if      |
| 循环             | for          |
| 范围             | range        |
| 结构体           | struct       |
| 接口             | interface    |
| 类型             | type         |
| 映射             | map          |
| 转到             | go           |
| 延迟             | defer        |
| 空/无            | nil          |

See the `internal/transpiler/keywords.go` file for a complete list of keywords and mappings.

## 支持的包 (Supported Packages)

Saika provides Chinese wrappers for standard Go packages:

| Saika (Chinese) | Go (English)    |
|-----------------|-----------------|
| 工具             | fmt             |
| 时间             | time            |
| 系统             | os              |
| 输入输出          | io              |
| 同步             | sync            |
| 字符串工具        | strings         |
| 网络             | net             |
| 网络http         | net/http        |
| 数学             | math            |
| 随机             | math/rand       |
| 编码json         | encoding/json   |
| 文件路径          | path/filepath   |

## 如何贡献 (Contributing)

Contributions are welcome! Here are some ways you can help:

1. Report bugs or suggest features by opening issues
2. Improve documentation or add examples
3. Submit pull requests with bug fixes or new features
4. Add wrappers for more Go packages
5. Help translate error messages and documentation

Please follow the existing code style and include tests for new functionality.

## 许可证 (License)

Saika is released under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## 致谢 (Acknowledgments)

Thanks to all contributors and the Go community for making this project possible.

---

**友情提示：** Saika 是一个教育性和实验性的项目，旨在帮助中文母语的程序员学习编程概念。对于生产环境，我们建议使用标准的 Go 语言。

**Note:** Saika is an educational and experimental project designed to help Chinese speakers learn programming concepts. For production environments, we recommend using standard Go.