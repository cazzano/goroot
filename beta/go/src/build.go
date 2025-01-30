// build.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

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

// handleBuild checks for Go files and runs the go build command.
func handleBuild() error {
	// Start tracking the build progress
	progress := NewBuildProgress()
	progress.StartBuildProcess()

	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		progress.HandleBuildError(fmt.Errorf("error getting current directory: %v", err))
		return err
	}

	// Check for Go files in the current directory
	hasGoFile, err := checkForGoFiles(currentDir)
	if err != nil {
		progress.HandleBuildError(err)
		return err
	}

	if !hasGoFile {
		err := fmt.Errorf("no Go files found in the current directory")
		progress.HandleBuildError(err)
		return err
	}

	// Update progress before building
	progress.UpdateProgress("Building the project")

	// Build the project in the current directory
	cmd := exec.Command("go", "build", ".") // Use "." to build all Go files in the current directory
	output, err := cmd.CombinedOutput()     // Capture combined output (stdout and stderr)
	if err != nil {
		err = fmt.Errorf("error building project: %v\nOutput: %s", err, output)
		progress.HandleBuildError(err)
		return err
	}

	fmt.Printf("Build successful! Output:\n%s\n", output)

	// Update progress after successful build
	progress.UpdateProgress("Moving binary to release directory")

	// Change to the parent directory
	parentDir := filepath.Dir(currentDir)
	if err := os.Chdir(parentDir); err != nil {
		err = fmt.Errorf("error changing to parent directory: %v", err)
		progress.HandleBuildError(err)
		return err
	}

	// Create the target/release directory if it doesn't exist
	releaseDir := filepath.Join(parentDir, "target", "release")
	if err := os.MkdirAll(releaseDir, 0755); err != nil {
		err = fmt.Errorf("error creating release directory: %v", err)
		progress.HandleBuildError(err)
		return err
	}

	// Move the compiled binary to the target/release directory
	binaryName := filepath.Base(currentDir) // Use the current directory name as the binary name
	srcBinaryPath := filepath.Join(currentDir, binaryName)
	destBinaryPath := filepath.Join(releaseDir, binaryName)

	if err := os.Rename(srcBinaryPath, destBinaryPath); err != nil {
		err = fmt.Errorf("error moving binary to release directory: %v", err)
		progress.HandleBuildError(err)
		return err
	}

	// Complete the build progress tracking
	progress.CompleteBuildProcess(destBinaryPath)
	return nil
}
