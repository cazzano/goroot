// build.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings" // Added to resolve import issue
)

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

	// Check for specific plain files in the current directory
	plainFiles := []string{"filename1", "filename2"} // Replace with your actual filenames
	hasPlainFile, err := checkForPlainFiles(currentDir, plainFiles)
	if err != nil {
		return err
	}

	// Call the debug function to print debug information
	DebugBuild(currentDir, hasGoFile, hasPlainFile)

	// If no Go files in current directory and specific plain files are found, try src directory
	if !hasGoFile && hasPlainFile {
		// Change to the parent directory twice
		parentDir := filepath.Dir(filepath.Dir(currentDir)) // Use parentDir to resolve unused variable
		if err := os.Chdir(parentDir); err != nil {
			return fmt.Errorf("error changing to parent directory: %v", err)
		}

		// Now change to src directory
		srcDir := filepath.Join(currentDir, "src")
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
		DebugBuildError(err)
		return fmt.Errorf("error building project: %v\nOutput: %s", err, output)
	}

	// Call the debug function to indicate success
	DebugBuildSuccess(filepath.Join(currentDir, filepath.Base(currentDir)))

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

// Helper function to check for specific plain files in a directory
func checkForPlainFiles(dir string, filenames []string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, fmt.Errorf("error reading directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			// Use strings.Contains to check for partial matches if needed
			for _, filename := range filenames {
				if entry.Name() == filename || strings.Contains(entry.Name(), filename) {
					fmt.Printf("[DEBUG] Found plain file in %s: %s\n", dir, entry.Name())
					return true, nil
				}
			}
		}
	}

	return false, nil
}
