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

**Linux (x86_64):**
```bash
# Extract the archive
tar -xzf saika-linux-amd64.tar.gz

# Change to extracted directory
cd linux-amd64

# Make executable
chmod +x saika

# Copy to system path
sudo cp saika /usr/local/bin/

# Verify installation
saika

# Clean up
cd ..
rm -rf linux-amd64
rm saika-linux-amd64.tar.gz
```

**Linux (ARM64):**
```bash
# Extract the archive
tar -xzf saika-linux-arm64.tar.gz

# Change to extracted directory
cd linux-arm64

# Make executable
chmod +x saika

# Copy to system path
sudo cp saika /usr/local/bin/

# Verify installation
saika

# Clean up
cd ..
rm -rf linux-arm64
rm saika-linux-arm64.tar.gz
```

**Linux (386):**
```bash
# Extract the archive
tar -xzf saika-linux-386.tar.gz

# Change to extracted directory
cd linux-386

# Make executable
chmod +x saika

# Copy to system path
sudo cp saika /usr/local/bin/

# Verify installation
saika

# Clean up
cd ..
rm -rf linux-386
rm saika-linux-386.tar.gz
```

**macOS (Intel):**
```bash
# Extract the archive
tar -xzf saika-macos-amd64.tar.gz

# Change to extracted directory
cd macos-amd64

# Make executable
chmod +x saika

# Copy to system path
sudo cp saika /usr/local/bin/

# Verify installation
saika

# Clean up
cd ..
rm -rf macos-amd64
rm saika-macos-amd64.tar.gz
```

**macOS (Apple Silicon):**
```bash
# Extract the archive
tar -xzf saika-macos-arm64.tar.gz

# Change to extracted directory
cd macos-arm64

# Make executable
chmod +x saika

# Copy to system path
sudo cp saika /usr/local/bin/

# Verify installation
saika

# Clean up
cd ..
rm -rf macos-arm64
rm saika-macos-arm64.tar.gz
```

**Windows (x86_64):**
```powershell
# Extract the archive
Expand-Archive saika-windows-amd64.zip

# Change to extracted directory
cd saika-windows-amd64

# Copy to a directory in your PATH
copy saika.exe C:\Windows\System32\

# Verify installation
saika

# Clean up
cd ..
Remove-Item -Recurse -Force saika-windows-amd64
Remove-Item saika-windows-amd64.zip
```

**Windows (ARM64):**
```powershell
# Extract the archive
Expand-Archive saika-windows-arm64.zip

# Change to extracted directory
cd saika-windows-arm64

# Copy to a directory in your PATH
copy saika.exe C:\Windows\System32\

# Verify installation
saika

# Clean up
cd ..
Remove-Item -Recurse -Force saika-windows-arm64
Remove-Item saika-windows-arm64.zip
```

**Windows (386):**
```powershell
# Extract the archive
Expand-Archive saika-windows-386.zip

# Change to extracted directory
cd saika-windows-386

# Copy to a directory in your PATH
copy saika.exe C:\Windows\System32\

# Verify installation
saika

# Clean up
cd ..
Remove-Item -Recurse -Force saika-windows-386
Remove-Item saika-windows-386.zip
```

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
    工具.打印行("我不好世界")
}
```

This will print "我不好世界" (a play on "Hello World").

### 更复杂的示例 (More Complex Example)

```go
包 入口

导入 "工具"

// 数值计算函数
数 计算(x 整数, y 整数) 整数 {
    变量 结果 = x + y
    返回 结果
}

// 打印信息函数
数 打印信息(消息 字符串) {
    工具.打印行(消息)
}

// 判断是否是偶数
数 是偶数(数字 整数) 布尔 {
    如果 数字 % 2 == 0 {
        返回 真
    } 否则 {
        返回 假
    }
}

// 主函数
数 入口() {
    // 定义变量
    变量 姓名 = "赵明"
    变量 年龄 = 25
    常量 问候语 = "你好，世界！"

    // 输出问候语
    工具.打印行(问候语)
    
    // 使用条件语句
    如果 年龄 > 18 {
        工具.打印行(姓名, "已经成年")
    } 否则 {
        工具.打印行(姓名, "未成年")
    }
    
    // 使用循环
    变量 总和 = 0
    循环 i := 1; i <= 10; i = i + 1 {
        总和 = 总和 + i
    }
    工具.打印行("1到10的总和是:", 总和)
    
    // 调用函数
    变量 计算结果 = 计算(5, 7)
    工具.打印行("5 + 7 =", 计算结果)
    
    // 检查奇偶性
    变量 数字 = 4
    如果 是偶数(数字) {
        工具.打印行(数字, "是偶数")
    } 否则 {
        工具.打印行(数字, "是奇数")
    }
    
    // 打印自定义信息
    打印信息("程序执行完毕")
}
```

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