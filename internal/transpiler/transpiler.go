package transpiler

import (
	"fmt"
	"go/parser"
	"go/printer"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"strings"
)

// Transpiler takes Saika code with Chinese keywords and converts it to Go code
type Transpiler struct {
	keywordMap      map[string]string
	errorTranslator *ErrorTranslator
}

// New creates a new Transpiler instance with Chinese to Go keyword mappings
func New() *Transpiler {
	return &Transpiler{
		keywordMap:      ExtendedKeywordMap(),
		errorTranslator: NewErrorTranslator(),
	}
}

// TranspileFile reads a Saika file and converts it to Go code
func (t *Transpiler) TranspileFile(filename string) (string, error) {
	// Read the Saika source file
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %v", filename, err)
	}

	saikaCode := string(bytes)
	return t.Transpile(saikaCode)
}

// Transpile converts Saika code to Go code
func (t *Transpiler) Transpile(saikaCode string) (string, error) {
	// Preprocess the code to replace Chinese keywords with Go keywords
	// First try with the enhanced preprocessor
	goCode, err := t.EnhancedPreprocess(saikaCode)
	if err != nil {
		// Fall back to the original preprocessor if there's an error
		goCode, err = t.preprocess(saikaCode)
		if err != nil {
			return "", err
		}
	}

	// Handle special cases like package declarations
	goCode, err = t.HandleSpecialCases(goCode)
	if err != nil {
		return "", err
	}

	// Process the AST for additional transformations
	goCode, err = t.ProcessAST(goCode)
	if err != nil {
		// If there's an error in AST processing, just continue with the current code
		// This is non-fatal as the AST processing is an enhancement
	}

	// Parse using Go's parser to validate the resulting code
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", goCode, parser.ParseComments)
	if err != nil {
		// Translate the error message to use Chinese keywords
		translatedErr := t.errorTranslator.TranslateError(err)
		return "", fmt.Errorf("解析错误: %v", translatedErr)
	}

	// Format the code
	var buf strings.Builder
	if err := printer.Fprint(&buf, fset, file); err != nil {
		return "", fmt.Errorf("formatting error: %v", err)
	}

	return buf.String(), nil
}

// preprocess replaces Chinese keywords with their Go equivalents
func (t *Transpiler) preprocess(saikaCode string) (string, error) {
	// Create a scanner to tokenize the input
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(saikaCode))

	// Initialize the scanner with the Saika code
	var scanErrors scanner.ErrorList
	s.Init(file, []byte(saikaCode), scanErrors.Add, scanner.ScanComments)

	// Process tokens and build a transformed source
	var result strings.Builder
	var lastPos token.Pos

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}

		// Add any spacing between tokens
		if lastPos > 0 {
			gap := int(pos) - int(lastPos) - len(lastLit)
			if gap > 0 {
				result.WriteString(strings.Repeat(" ", gap))
			}
		}

		// Check if the token is an identifier that matches a Chinese keyword
		if tok == token.IDENT {
			if replacement, ok := t.keywordMap[lit]; ok {
				// Replace with Go keyword
				result.WriteString(replacement)
			} else {
				// Keep the original identifier
				result.WriteString(lit)
			}
		} else {
			// For other tokens, use the literal value
			result.WriteString(lit)
		}

		lastPos = pos
		lastLit = lit
	}

	return result.String(), nil
}

// Helper variables for the preprocessor
var lastLit string
