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
	fmt.Printf("[DEBUG] Current directory: %s\n", currentDir)

	// Check for Go files in the current directory
	hasGoFile, err := checkForGoFiles(currentDir)
	if err != nil {
		return err
	}

	// If no Go files in current directory, try src directory
	if !hasGoFile {
		srcDir := filepath.Join(currentDir, "src")
		// Check if src directory exists
		if _, err := os.Stat(srcDir); os.IsNotExist(err) {
			return fmt.Errorf("no Go files found and no 'src' directory exists")
		}

		// Change to src directory
		if err := os.Chdir(srcDir); err != nil {
			return fmt.Errorf("error changing to src directory: %v", err)
		}
		defer os.Chdir(currentDir) // Change back to original directory after build

		// Recheck for Go files in src directory
		hasGoFile, err = checkForGoFiles(srcDir)
		if err != nil {
			return err
		}

		if !hasGoFile {
			return fmt.Errorf("no Go files found in current directory or src directory")
		}

		fmt.Println("[DEBUG] Switching to src directory for build")
		currentDir = srcDir
	}

	// Build the project in the current directory
	cmd := exec.Command("go", "build", ".") // Use "." to build all Go files in the current directory
	cmd.Dir = currentDir                    // Set the working directory for the command
	output, err := cmd.CombinedOutput()     // Capture combined output (stdout and stderr)
	if err != nil {
		return fmt.Errorf("error building project: %v\nOutput: %s", err, output)
	}

	fmt.Println("Build successful!")

	// Define the target release directory
	releaseDir := filepath.Join(filepath.Dir(currentDir), "target", "release")
	if err := os.MkdirAll(releaseDir, 0755); err != nil {
		return fmt.Errorf("error creating release directory: %v", err)
	}

	// Move the compiled binary to the target/release directory
	binaryName := filepath.Base(currentDir) // Use the current directory name as the binary name
	srcBinaryPath := filepath.Join(currentDir, binaryName)
	destBinaryPath := filepath.Join(releaseDir, binaryName)

	if err := os.Rename(srcBinaryPath, destBinaryPath); err != nil {
		return fmt.Errorf("error moving binary to release directory: %v", err)
	}

	fmt.Printf("Binary moved to: %s\n", destBinaryPath)
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
