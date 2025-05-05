// 包 工具 provides Chinese wrappers for the fmt package
package 工具

import (
	"fmt"
	"io"
)

// 打印 wraps fmt.Print
func 打印(参数 ...interface{}) (n int, err error) {
	return fmt.Print(参数...)
}

// 打印行 wraps fmt.Println
func 打印行(参数 ...interface{}) (n int, err error) {
	return fmt.Println(参数...)
}

// 打印格式 wraps fmt.Printf
func 打印格式(格式 string, 参数 ...interface{}) (n int, err error) {
	return fmt.Printf(格式, 参数...)
}

// 扫描 wraps fmt.Scan
func 扫描(参数 ...interface{}) (n int, err error) {
	return fmt.Scan(参数...)
}

// 扫描行 wraps fmt.Scanln
func 扫描行(参数 ...interface{}) (n int, err error) {
	return fmt.Scanln(参数...)
}

// 扫描格式 wraps fmt.Scanf
func 扫描格式(格式 string, 参数 ...interface{}) (n int, err error) {
	return fmt.Scanf(格式, 参数...)
}

// 错误打印 wraps fmt.Fprint
func 错误打印(输出 io.Writer, 参数 ...interface{}) (n int, err error) {
	return fmt.Fprint(输出, 参数...)
}

// 错误打印行 wraps fmt.Fprintln
func 错误打印行(输出 io.Writer, 参数 ...interface{}) (n int, err error) {
	return fmt.Fprintln(输出, 参数...)
}

// 错误打印格式 wraps fmt.Fprintf
func 错误打印格式(输出 io.Writer, 格式 string, 参数 ...interface{}) (n int, err error) {
	return fmt.Fprintf(输出, 格式, 参数...)
}

// 字符串格式 wraps fmt.Sprintf
func 字符串格式(格式 string, 参数 ...interface{}) string {
	return fmt.Sprintf(格式, 参数...)
}

// 字符串扫描 wraps fmt.Sscanf
func 字符串扫描(字符串 string, 格式 string, 参数 ...interface{}) (n int, err error) {
	return fmt.Sscanf(字符串, 格式, 参数...)
}

// Sprint wraps fmt.Sprint
func 字符串打印(参数 ...interface{}) string {
	return fmt.Sprint(参数...)
}

// Sprintln wraps fmt.Sprintln
func 字符串打印行(参数 ...interface{}) string {
	return fmt.Sprintln(参数...)
}

// Errorf wraps fmt.Errorf
func 错误格式(格式 string, 参数 ...interface{}) error {
	return fmt.Errorf(格式, 参数...)
}
