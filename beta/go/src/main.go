// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 || (os.Args[1] != "build" && os.Args[1] != "-v" && os.Args[1] != "help") {
		DisplayHelp() // Call the help function if the command is invalid
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build":
		if err := handleBuild(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "-v":
		DisplayVersion() // Call the version display function
		os.Exit(0)
	case "help":
		DisplayHelp() // Call the help function
		os.Exit(0)
	default:
		DisplayHelp() // Call the help function if the command is invalid
		os.Exit(1)
	}
}
