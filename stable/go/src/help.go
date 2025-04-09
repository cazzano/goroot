// help.go
package main

import io "fmt"

// DisplayHelp prints the usage instructions and available commands.
func DisplayHelp() {
	io.Println("Usage: goroot <command>")
	io.Println("Commands:")
	io.Println("  init         - Initialize the project structure")
	io.Println("  build        - Build the project")
	io.Println("  run          - Run Go files in the current directory (max 10 arguments)")
	io.Println("  run --1      - With specefic file or module")
	io.Println("  --v          - Display the version information")
	io.Println("  --h          - Display this help message")
}
