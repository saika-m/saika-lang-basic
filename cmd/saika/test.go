package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/saika-m/saika-lang-basic/internal/transpiler"
)

// TestCommand performs a test transpilation and displays the original Saika code
// alongside the transpiled Go code for inspection and debugging purposes.
//
// This command is intended for:
// - Learning how Saika maps to Go
// - Debugging transpilation issues
// - Understanding how Chinese keywords are translated
func TestCommand(t *transpiler.Transpiler, saikaFile string) {
	// Check if the file exists
	if _, err := os.Stat(saikaFile); os.IsNotExist(err) {
		fmt.Printf("Error: File %s does not exist\n", saikaFile)
		os.Exit(1)
	}

	// Read the Saika source file
	saikaCode, err := ioutil.ReadFile(saikaFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Transpile the Saika code to Go
	goCode, err := t.Transpile(string(saikaCode))
	if err != nil {
		fmt.Printf("Error transpiling code: %v\n", err)
		os.Exit(1)
	}

	// Print the original and transpiled code side by side
	fmt.Println("Original Saika code:")
	fmt.Println("--------------------")
	fmt.Println(string(saikaCode))
	fmt.Println("\nTranspiled Go code:")
	fmt.Println("------------------")
	fmt.Println(goCode)

	// Offer to save the Go code to a file
	fmt.Println("\nWould you like to save the transpiled Go code to a file? (y/n)")
	var response string
	fmt.Scanln(&response)

	if response == "y" || response == "Y" {
		goFilename := strings.TrimSuffix(saikaFile, ".saika") + ".go"
		err := ioutil.WriteFile(goFilename, []byte(goCode), 0644)
		if err != nil {
			fmt.Printf("Error writing Go file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Transpiled Go code saved to: %s\n", goFilename)
	}
}
