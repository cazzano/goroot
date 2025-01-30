// build.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// handleBuild checks for Go files and runs the go build command.
func handleBuild() error {
	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %v", err)
	}

	// Check for Go files in the current directory
	hasGoFile, err := checkForGoFiles(currentDir)
	if err != nil {
		return err
	}

	if !hasGoFile {
		return fmt.Errorf("no Go files found in the current directory")
	}

	// Build the project in the current directory
	cmd := exec.Command("go", "build", ".") // Use "." to build all Go files in the current directory
	cmd.Dir = currentDir                    // Set the working directory for the command
	output, err := cmd.CombinedOutput()     // Capture combined output (stdout and stderr)
	if err != nil {
		return fmt.Errorf("error building project: %v\nOutput: %s", err, output)
	}

	fmt.Printf("Build successful! Output:\n%s\n", output)
	return nil
}

// Helper function to check for Go files in a directory
func checkForGoFiles(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, fmt.Errorf("error reading directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".go" {
			fmt.Printf("[DEBUG] Found Go file in %s: %s\n", dir, entry.Name())
			return true, nil
		}
	}

	return false, nil
}
