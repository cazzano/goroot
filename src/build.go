// build.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func handleBuild() error {
	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %v", err)
	}

	// Check for Go files in the current directory
	hasGoFile := false
	entries, err := os.ReadDir(currentDir)
	if err != nil {
		return fmt.Errorf("error reading current directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".go" {
			hasGoFile = true
			break
		}
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

	fmt.Println("Build successful!")
	return nil
}
