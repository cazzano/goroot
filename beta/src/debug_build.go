// debug_build.go
package main

import (
	"fmt"
)

// DebugBuild prints debug information during the build process.
func DebugBuild(currentDir string, hasGoFile bool, hasPlainFile bool) {
	fmt.Printf("[DEBUG] Build process started in directory: %s\n", currentDir)
	if hasGoFile {
		fmt.Println("[DEBUG] Go files are present in the directory.")
	} else {
		fmt.Println("[DEBUG] No Go files found in the directory.")
	}

	if hasPlainFile {
		fmt.Println("[DEBUG] Plain files are present in the directory.")
	} else {
		fmt.Println("[DEBUG] No plain files found in the directory.")
	}
}

// DebugBuildSuccess prints a success message after the build.
func DebugBuildSuccess(binaryPath string) {
	fmt.Printf("[DEBUG] Build successful! Binary created at: %s\n", binaryPath)
}

// DebugBuildError prints an error message during the build process.
func DebugBuildError(err error) {
	fmt.Printf("[DEBUG] Build error: %v\n", err)
}
