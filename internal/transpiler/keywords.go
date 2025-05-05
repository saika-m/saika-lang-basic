package transpiler

// ExtendedKeywordMap returns a comprehensive mapping of Chinese keywords to Go keywords
func ExtendedKeywordMap() map[string]string {
	return map[string]string{
		// Package names (special handling)
		"入口":     "main",          // main package
		"工具":     "fmt",           // fmt package
		"输入输出":   "io",            // io package
		"系统":     "os",            // os package
		"时间":     "time",          // time package
		"字符串工具":  "strings",       // strings package
		"排序":     "sort",          // sort package
		"同步":     "sync",          // sync package
		"上下文":    "context",       // context package
		"网络":     "net",           // net package
		"网络http": "net/http",      // net/http package
		"数学":     "math",          // math package
		"随机":     "math/rand",     // math/rand package
		"日志":     "log",           // log package
		"编码json": "encoding/json", // encoding/json package
		"编码xml":  "encoding/xml",  // encoding/xml package
		"文件路径":   "path/filepath", // path/filepath package
		"运行时":    "runtime",       // runtime package
		"缓冲区":    "bytes",         // bytes package
		"正则表达式":  "regexp",        // regexp package
		"测试":     "testing",       // testing package

		// Basic structure
		"包":  "package",
		"导入": "import",
		"数":  "func",
		"函数": "func",
		"返回": "return",

		// Data types
		"整数":      "int",
		"整数8":     "int8",
		"整数16":    "int16",
		"整数32":    "int32",
		"整数64":    "int64",
		"无符号整数":   "uint",
		"无符号整数8":  "uint8",
		"无符号整数16": "uint16",
		"无符号整数32": "uint32",
		"无符号整数64": "uint64",
		"字节":      "byte",
		"单精度浮点数":  "float32",
		"双精度浮点数":  "float64",
		"浮点数32":   "float32",
		"浮点数64":   "float64",
		"浮点数":     "float64",
		"复数64":    "complex64",
		"复数128":   "complex128",
		"复数":      "complex128",
		"字符串":     "string",
		"字符":      "rune",
		"布尔":      "bool",
		"真":       "true",
		"假":       "false",
		"空":       "nil",
		"无":       "nil",
		"任意":      "interface{}",
		"接口体":     "interface{}",
		"错误":      "error",

		// Variables and constants
		"变量": "var",
		"常量": "const",
		"类型": "type",
		"数值": "iota",

		// Control flow
		"如果":   "if",
		"否则":   "else",
		"否则如果": "else if",
		"循环":   "for",
		"范围":   "range",
		"开关":   "switch",
		"情况":   "case",
		"默认":   "default",
		"中断":   "break",
		"继续":   "continue",
		"跳转":   "goto",
		"回调":   "fallthrough",
		"选择":   "select",

		// Data structures
		"结构体": "struct",
		"接口":  "interface",
		"映射":  "map",
		"切片":  "slice",
		"数组":  "array",
		"通道":  "chan",

		// Error handling
		"延迟": "defer",
		"恢复": "recover",
		"抛出": "panic",

		// Modifiers
		"转到": "go",
		"指针": "*",
		"引用": "&",

		// Memory management
		"新建": "new",
		"创建": "make",
		"删除": "delete",
		"追加": "append",
		"复制": "copy",
		"长度": "len",
		"容量": "cap",

		// Methods
		"打印":     "Print",
		"打印行":    "Println",
		"打印格式":   "Printf",
		"扫描":     "Scan",
		"扫描行":    "Scanln",
		"扫描格式":   "Scanf",
		"错误打印":   "Fprint",
		"错误打印行":  "Fprintln",
		"错误打印格式": "Fprintf",
		"字符串格式":  "Sprintf",

		// Operators
		"和": "&&",
		"或": "||",
		"非": "!",

		// Access modifiers
		"公开": "", // Used to indicate exported (capitalized) names
		"私有": "", // Used to indicate unexported (lowercase) names
	}
}

// PackageNameMap returns a mapping of Chinese package names to Go package names
func PackageNameMap() map[string]string {
	return map[string]string{
		"入口":     "main",
		"工具":     "fmt",
		"输入输出":   "io",
		"系统":     "os",
		"时间":     "time",
		"字符串工具":  "strings",
		"排序":     "sort",
		"同步":     "sync",
		"上下文":    "context",
		"网络":     "net",
		"网络http": "net/http",
		"数学":     "math",
		"随机":     "math/rand",
		"日志":     "log",
		"编码json": "encoding/json",
		"编码xml":  "encoding/xml",
		"文件路径":   "path/filepath",
		"运行时":    "runtime",
		"缓冲区":    "bytes",
		"正则表达式":  "regexp",
		"测试":     "testing",
	}
}

// UpdateTranspilerKeywords updates the transpiler with a more comprehensive keyword map
func (t *Transpiler) UpdateKeywords() {
	t.keywordMap = ExtendedKeywordMap()
}
