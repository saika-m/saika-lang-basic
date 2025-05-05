package transpiler

import (
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

// HandleSpecialCases processes special cases in the code that need custom handling
func (t *Transpiler) HandleSpecialCases(code string) (string, error) {
	// Handle the special case of package declarations
	code = t.handlePackageDeclarations(code)

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
