// init.go
package main

import (
	"fmt"
	"os/exec"
)

func handleInit() error {
	// Run the command to initialize the Go module
	cmd := exec.Command("go", "mod", "init", "main")
	output, err := cmd.CombinedOutput() // Capture combined output (stdout and stderr)
	if err != nil {
		return fmt.Errorf("error initializing module: %v\nOutput: %s", err, output)
	}

	fmt.Printf("Module initialized successfully:\n%s\n", output)
	return nil
}
