package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/saika-m/saika-lang-basic/internal/transpiler"
)

const Version = "1.0.3"

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	saikaFile := os.Args[2]

	// Create a transpiler instance
	t := transpiler.New()

	switch command {
	case "build":
		buildCommand(t, saikaFile)
	case "run":
		runCommand(t, saikaFile)
	case "test":
		// The test command is implemented in test.go
		TestCommand(t, saikaFile)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  saika build <file.saika>  - Compile the Saika file to an executable")
	fmt.Println("  saika run <file.saika>    - Run the Saika file")
	fmt.Println("  saika test <file.saika>   - Show the transpiled Go code for a Saika file")
}

func buildCommand(t *transpiler.Transpiler, saikaFile string) {
	// Check if the file exists
	if _, err := os.Stat(saikaFile); os.IsNotExist(err) {
		fmt.Printf("Error: File %s does not exist\n", saikaFile)
		os.Exit(1)
	}

	// Transpile the Saika file to Go (silently)
	goCode, err := t.TranspileFile(saikaFile)
	if err != nil {
		fmt.Printf("Error transpiling file: %v\n", err)
		os.Exit(1)
	}

	// Create a temporary Go file (silently)
	tempGoFile, tempDir, err := createTempGoFile(goCode)
	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tempDir) // Clean up temporary directory

	// Setup Go environment for Chinese imports
	setupGoEnv()

	// Compile the Go file
	outputFile := strings.TrimSuffix(saikaFile, ".saika")
	cmd := exec.Command("go", "build", "-o", outputFile, tempGoFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Add the lib directory to GOPATH
	cmd.Env = append(os.Environ(), "GOPATH="+os.Getenv("GOPATH")+":"+getLibDir())

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error compiling file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully built: %s\n", outputFile)
}

func runCommand(t *transpiler.Transpiler, saikaFile string) {
	// Check if the file exists
	if _, err := os.Stat(saikaFile); os.IsNotExist(err) {
		fmt.Printf("Error: File %s does not exist\n", saikaFile)
		os.Exit(1)
	}

	// Transpile the Saika file to Go (silently)
	goCode, err := t.TranspileFile(saikaFile)
	if err != nil {
		fmt.Printf("Error transpiling file: %v\n", err)
		os.Exit(1)
	}

	// Create a temporary Go file (silently)
	tempGoFile, tempDir, err := createTempGoFile(goCode)
	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tempDir) // Clean up temporary directory

	// Setup Go environment for Chinese imports
	setupGoEnv()

	// Run the Go file
	cmd := exec.Command("go", "run", tempGoFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Add the lib directory to GOPATH
	cmd.Env = append(os.Environ(), "GOPATH="+os.Getenv("GOPATH")+":"+getLibDir())

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running file: %v\n", err)
		os.Exit(1)
	}
}

// Helper function to create a temporary Go file
func createTempGoFile(goCode string) (string, string, error) {
	// Create a temporary directory
	tempDir, err := ioutil.TempDir("", "saika-")
	if err != nil {
		return "", "", fmt.Errorf("failed to create temp directory: %v", err)
	}

	// Create a temporary Go file
	tempFile := filepath.Join(tempDir, "main.go")
	err = ioutil.WriteFile(tempFile, []byte(goCode), 0644)
	if err != nil {
		os.RemoveAll(tempDir) // Clean up
		return "", "", fmt.Errorf("failed to write temp file: %v", err)
	}

	return tempFile, tempDir, nil
}

// Helper function to setup the Go environment for Chinese imports
func setupGoEnv() {
	// No implementation needed since we handle translation in the transpiler
}

// Helper function to get the lib directory
func getLibDir() string {
	// Get the directory of the current executable
	execDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		// Fallback to the current directory
		execDir, _ = os.Getwd()
	}

	// Go up one directory and find the lib directory
	baseDir := filepath.Dir(execDir)
	libDir := filepath.Join(baseDir, "lib")

	return libDir
}
