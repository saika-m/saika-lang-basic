package transpiler

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

// CustomTokenizer extends the Go scanner to better handle Chinese characters
type CustomTokenizer struct {
	src        []byte
	offset     int
	rdOffset   int
	ch         rune
	keywordMap map[string]string
}

// NewCustomTokenizer creates a new tokenizer for Saika code
func NewCustomTokenizer(src []byte, keywordMap map[string]string) *CustomTokenizer {
	t := &CustomTokenizer{
		src:        src,
		keywordMap: keywordMap,
	}
	t.next()
	return t
}

// next reads the next Unicode character
func (t *CustomTokenizer) next() {
	if t.rdOffset < len(t.src) {
		t.offset = t.rdOffset
		r, w := rune(t.src[t.rdOffset]), 1
		if r >= utf8.RuneSelf {
			r, w = utf8.DecodeRune(t.src[t.rdOffset:])
		}
		t.rdOffset += w
		t.ch = r
	} else {
		t.offset = len(t.src)
		t.ch = -1 // EOF
	}
}

// peek returns the next Unicode character without advancing
func (t *CustomTokenizer) peek() rune {
	if t.rdOffset < len(t.src) {
		r, _ := utf8.DecodeRune(t.src[t.rdOffset:])
		return r
	}
	return -1 // EOF
}

// isChineseChar checks if a rune is a Chinese character
func isChineseChar(r rune) bool {
	return unicode.Is(unicode.Han, r)
}

// isIdentifier checks if a rune can be part of an identifier
func isIdentifier(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) || isChineseChar(r)
}

// scanIdentifier scans an identifier
func (t *CustomTokenizer) scanIdentifier() string {
	offs := t.offset
	for isIdentifier(t.ch) {
		t.next()
	}
	return string(t.src[offs:t.offset])
}

// scanWhitespace scans whitespace
func (t *CustomTokenizer) scanWhitespace() string {
	offs := t.offset
	for unicode.IsSpace(t.ch) {
		t.next()
	}
	return string(t.src[offs:t.offset])
}

// scanComment scans a comment
func (t *CustomTokenizer) scanComment() (string, bool) {
	// Check for line comment
	if t.ch == '/' && t.peek() == '/' {
		offs := t.offset
		t.next() // consume first '/'
		t.next() // consume second '/'

		// Scan until end of line
		for t.ch != '\n' && t.ch >= 0 {
			t.next()
		}
		return string(t.src[offs:t.offset]), true
	}

	// Check for block comment
	if t.ch == '/' && t.peek() == '*' {
		offs := t.offset
		t.next() // consume '/'
		t.next() // consume '*'

		for {
			// Check for end of comment
			if t.ch == '*' && t.peek() == '/' {
				t.next() // consume '*'
				t.next() // consume '/'
				break
			}

			// EOF before end of comment
			if t.ch < 0 {
				return string(t.src[offs:t.offset]), true
			}

			t.next()
		}
		return string(t.src[offs:t.offset]), true
	}

	return "", false
}

// scanString scans a string literal
func (t *CustomTokenizer) scanString() (string, bool) {
	if t.ch == '"' || t.ch == '`' {
		quote := t.ch
		offs := t.offset
		t.next() // consume opening quote

		for {
			// Check for end of string
			if t.ch == quote {
				t.next() // consume closing quote
				break
			}

			// EOF before end of string
			if t.ch < 0 {
				return string(t.src[offs:t.offset]), true
			}

			// For double-quoted strings, handle escapes
			if quote == '"' && t.ch == '\\' {
				t.next() // consume backslash
				if t.ch < 0 {
					return string(t.src[offs:t.offset]), true
				}
			}

			t.next()
		}
		return string(t.src[offs:t.offset]), true
	}

	return "", false
}

// TokenizeAndTransform tokenizes Saika code and replaces Chinese keywords with Go equivalents
func (t *CustomTokenizer) TokenizeAndTransform() string {
	var buf bytes.Buffer

	for t.ch >= 0 {
		// Handle whitespace
		if unicode.IsSpace(t.ch) {
			ws := t.scanWhitespace()
			buf.WriteString(ws)
			continue
		}

		// Handle comments
		if comment, ok := t.scanComment(); ok {
			buf.WriteString(comment)
			continue
		}

		// Handle string literals
		if str, ok := t.scanString(); ok {
			buf.WriteString(str)
			continue
		}

		// Handle identifiers (including Chinese keywords)
		if unicode.IsLetter(t.ch) || t.ch == '_' || isChineseChar(t.ch) {
			ident := t.scanIdentifier()

			// Check if it's a Chinese keyword
			if replacement, ok := t.keywordMap[ident]; ok {
				buf.WriteString(replacement)
			} else {
				buf.WriteString(ident)
			}
			continue
		}

		// Handle operators and other characters
		buf.WriteRune(t.ch)
		t.next()
	}

	return buf.String()
}

// EnhancedPreprocess uses the custom tokenizer to better handle Chinese characters
func (t *Transpiler) EnhancedPreprocess(saikaCode string) (string, error) {
	tokenizer := NewCustomTokenizer([]byte(saikaCode), t.keywordMap)
	return tokenizer.TokenizeAndTransform(), nil
}
