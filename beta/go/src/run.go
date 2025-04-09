// run.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// handleRun runs Go files with optional arguments
func handleRun() error {
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

	// Prepare the command arguments
	args := os.Args[2:] // Skip the "run" command

	// Check if the first argument is "-one"
	if len(args) > 0 && args[0] == "--1" {
		// Limit arguments to 10 (including the "-one" flag)
		if len(args) > 11 { // 10 arguments + 1 for "-one"
			return fmt.Errorf("maximum of 10 arguments allowed, received %d", len(args)-1)
		}

		// The Go file to run is the second argument
		if len(args) < 2 {
			return fmt.Errorf("please provide a filename after '-one'")
		}

		goFile := args[1]
		if filepath.Ext(goFile) != ".go" {
			return fmt.Errorf("the specified file must have a .go extension")
		}

		// Construct the command
		cmdArgs := []string{"run", goFile}

		// Add user-provided arguments (excluding the first two)
		cmdArgs = append(cmdArgs, args[2:]...)

		// Create the command
		cmd := exec.Command("go", cmdArgs...)

		// Set the command's stdout and stderr to the current process's
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Print the command being executed (for debugging)
		fmt.Printf("[DEBUG] Running command: go %v\n", cmdArgs)

		// Run the command
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("error running Go file: %v", err)
		}
	} else {
		// Original run functionality
		// Limit arguments to 10
		if len(args) > 10 {
			return fmt.Errorf("maximum of 10 arguments allowed, received %d", len(args))
		}

		// Construct the full command
		cmdArgs := []string{"run"}

		// Find all .go files in the current directory
		goFiles, err := filepath.Glob(filepath.Join(currentDir, "*.go"))
		if err != nil {
			return fmt.Errorf("error finding Go files: %v", err)
		}

		// Add Go files to the command
		cmdArgs = append(cmdArgs, goFiles...)

		// Add user-provided arguments
		cmdArgs = append(cmdArgs, args...)

		// Create the command
		cmd := exec.Command("go", cmdArgs...)

		// Set the command's stdout and stderr to the current process's
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Print the command being executed (for debugging)
		fmt.Printf("[DEBUG] Running command: go %v\n", cmdArgs)

		// Run the command
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("error running Go files: %v", err)
		}
	}

	return nil
}
