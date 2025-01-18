// init.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func handleInit() error {
	// Get current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %v", err)
	}
	fmt.Printf("[DEBUG] Current directory: %s\n", currentDir)

	// Check if required files exist in current directory
	hasGoFile := false
	hasPlainFile := false

	entries, err := os.ReadDir(currentDir)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		baseName := strings.TrimSuffix(name, filepath.Ext(name))

		if filepath.Ext(name) == ".go" {
			hasGoFile = true
			fmt.Printf("[DEBUG] Found Go file: %s\n", name)
		}
		if name == baseName {
			hasPlainFile = true
			fmt.Printf("[DEBUG] Found plain file: %s\n", name)
		}
	}

	// Determine target directory for creating folders
	targetDir := currentDir
	if hasGoFile || hasPlainFile {
		targetDir = filepath.Dir(currentDir) // cd ..
		var fileStatus string
		if hasGoFile {
			fileStatus += ".go file present"
		}
		if hasPlainFile {
			if fileStatus != "" {
				fileStatus += ", "
			}
			fileStatus += "plain file present"
		}
		fmt.Printf("Found required file(s): %s\n", fileStatus)
	} else {
		fmt.Println("No required files found, creating directories in current location")
	}

	// Create src and target directories
	dirs := []string{"src", "target"}
	for _, dir := range dirs {
		dirPath := filepath.Join(targetDir, dir)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
		fmt.Printf("[DEBUG] Created directory: %s\n", dirPath)
	}

	return nil
}
