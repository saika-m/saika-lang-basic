package transpiler

import (
	"fmt"
	"go/scanner"
	"go/token"
	"regexp"
	"strings"
)

// ErrorTranslator translates Go error messages to Saika (Chinese) error messages
type ErrorTranslator struct {
	goToChineseKeywords map[string]string
	errorPatterns       map[string]string
}

// NewErrorTranslator creates a new ErrorTranslator
func NewErrorTranslator() *ErrorTranslator {
	// Create a reverse mapping from Go keywords to Chinese keywords
	chineseToGoKeywords := ExtendedKeywordMap()
	goToChineseKeywords := make(map[string]string)

	for chinese, go_kw := range chineseToGoKeywords {
		// Skip empty mappings
		if go_kw == "" {
			continue
		}

		// In case of duplicates, prefer the shorter Chinese keyword
		if existing, ok := goToChineseKeywords[go_kw]; ok {
			if len(chinese) < len(existing) {
				goToChineseKeywords[go_kw] = chinese
			}
		} else {
			goToChineseKeywords[go_kw] = chinese
		}
	}

	// Common error patterns and their Chinese translations
	errorPatterns := map[string]string{
		"expected":                    "预期",
		"found":                       "发现",
		"undefined":                   "未定义",
		"not declared":                "未声明",
		"syntax error":                "语法错误",
		"missing":                     "缺少",
		"cannot":                      "不能",
		"invalid":                     "无效",
		"redeclared":                  "重复声明",
		"constant":                    "常量",
		"variable":                    "变量",
		"function":                    "函数",
		"type":                        "类型",
		"package":                     "包",
		"import":                      "导入",
		"expression":                  "表达式",
		"statement":                   "语句",
		"assignment":                  "赋值",
		"declaration":                 "声明",
		"identifier":                  "标识符",
		"operand":                     "操作数",
		"operator":                    "操作符",
		"value":                       "值",
		"use of untyped nil":          "使用无类型的空值",
		"mismatched types":            "类型不匹配",
		"cannot use":                  "不能使用",
		"as type":                     "作为类型",
		"too many arguments":          "参数过多",
		"not enough arguments":        "参数不足",
		"undefined field":             "未定义的字段",
		"ambiguous selector":          "模糊的选择器",
		"non-name on left side of :=": "赋值符号左侧不是名称",
	}

	return &ErrorTranslator{
		goToChineseKeywords: goToChineseKeywords,
		errorPatterns:       errorPatterns,
	}
}

// TranslateErrorMessage translates a Go error message to a Saika error message
func (et *ErrorTranslator) TranslateErrorMessage(errorMsg string) string {
	// Replace Go keywords with Chinese keywords
	for goKeyword, chineseKeyword := range et.goToChineseKeywords {
		// Use word boundaries to avoid replacing parts of words
		pattern := fmt.Sprintf("\\b%s\\b", regexp.QuoteMeta(goKeyword))
		re := regexp.MustCompile(pattern)
		errorMsg = re.ReplaceAllString(errorMsg, chineseKeyword)
	}

	// Replace common error patterns
	for pattern, translation := range et.errorPatterns {
		errorMsg = strings.ReplaceAll(errorMsg, pattern, translation)
	}

	return errorMsg
}

// TranslateError translates a Go error to a Saika error
func (et *ErrorTranslator) TranslateError(err error) error {
	if err == nil {
		return nil
	}

	// For scanner.ErrorList, translate each error
	if errList, ok := err.(scanner.ErrorList); ok {
		for i := range errList {
			errList[i].Msg = et.TranslateErrorMessage(errList[i].Msg)
		}
		return errList
	}

	// For other errors, just translate the message
	return fmt.Errorf(et.TranslateErrorMessage(err.Error()))
}

// TranslateErrorWithMapping translates a Go error to a Saika error using the given position mapping
func (et *ErrorTranslator) TranslateErrorWithMapping(err error, posMap map[token.Pos]token.Pos) error {
	// TODO: Implement position mapping for errors
	return et.TranslateError(err)
}
