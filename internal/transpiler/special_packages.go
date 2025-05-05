package transpiler

import (
	"go/parser"
	"go/token"
	"regexp"
	"strconv"
	"strings"
)

// HandleSpecialCases processes special cases in the code that need custom handling
func (t *Transpiler) HandleSpecialCases(code string) (string, error) {
	// Handle the special case of package declarations
	code = t.handlePackageDeclarations(code)

	// Handle import statements with Chinese package names
	code = t.handleImportStatements(code)

	return code, nil
}

// handlePackageDeclarations handles package declarations with Chinese names
func (t *Transpiler) handlePackageDeclarations(code string) string {
	// Define a mapping for package declarations
	packageMap := PackageNameMap()

	// Look for package declarations in the code
	lines := strings.Split(code, "\n")
	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "package ") {
			packageName := strings.TrimPrefix(trimmedLine, "package ")
			packageName = strings.TrimSpace(packageName)

			// Check if the package name is in our Chinese package map
			if goPackageName, ok := packageMap[packageName]; ok {
				// Replace the line with the Go package name
				lines[i] = strings.Replace(line, "package "+packageName, "package "+goPackageName, 1)
			}
		}
	}

	return strings.Join(lines, "\n")
}

// handleImportStatements handles import statements with Chinese package names
func (t *Transpiler) handleImportStatements(code string) string {
	// Regular expressions to match different import formats
	singleImportRegex := regexp.MustCompile(`import\s+"([^"]+)"`)
	groupImportStartRegex := regexp.MustCompile(`import\s+\(`)
	groupImportLineRegex := regexp.MustCompile(`\s*"([^"]+)"`)

	packageMap := PackageNameMap()

	// Process single imports
	code = singleImportRegex.ReplaceAllStringFunc(code, func(match string) string {
		parts := singleImportRegex.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}

		importPath := parts[1]
		if goPackage, ok := packageMap[importPath]; ok {
			return strings.Replace(match, "\""+importPath+"\"", "\""+goPackage+"\"", 1)
		}

		return match
	})

	// Process grouped imports
	lines := strings.Split(code, "\n")
	inImportBlock := false

	for i, line := range lines {
		// Check if we're starting an import block
		if groupImportStartRegex.MatchString(line) {
			inImportBlock = true
			continue
		}

		// Check if we're ending an import block
		if inImportBlock && strings.TrimSpace(line) == ")" {
			inImportBlock = false
			continue
		}

		// Process lines inside import blocks
		if inImportBlock {
			parts := groupImportLineRegex.FindStringSubmatch(line)
			if len(parts) < 2 {
				continue
			}

			importPath := parts[1]
			if goPackage, ok := packageMap[importPath]; ok {
				lines[i] = strings.Replace(line, "\""+importPath+"\"", "\""+goPackage+"\"", 1)
			}
		}
	}

	return strings.Join(lines, "\n")
}

// ProcessAST processes the AST to handle special cases like imports
func (t *Transpiler) ProcessAST(code string) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		return code, err
	}

	// Process imports to handle Chinese package names
	packageMap := PackageNameMap()
	for _, imp := range file.Imports {
		// Get the import path
		importPath, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			continue
		}

		// Check if any part of the import path needs to be translated
		parts := strings.Split(importPath, "/")
		for i, part := range parts {
			if goPart, ok := packageMap[part]; ok {
				parts[i] = goPart
			}
		}

		// Update the import path
		newImportPath := strings.Join(parts, "/")
		if newImportPath != importPath {
			imp.Path.Value = strconv.Quote(newImportPath)
		}
	}

	// TODO: Add more AST processing as needed

	return code, nil
}
