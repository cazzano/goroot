// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] != "init" {
		fmt.Println("Usage: ./main init")
		os.Exit(1)
	}

	// Call the init command handler
	if err := handleInit(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
