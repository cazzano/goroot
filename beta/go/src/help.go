// help.go
package main

import "fmt"

// DisplayHelp prints the usage instructions and available commands.
func DisplayHelp() {
	fmt.Println("Usage: goroot <command>")
	fmt.Println("Commands:")
	fmt.Println("  init         - Initialize the existing project with go.mod")
	fmt.Println("  new          - Initialize a new project structure")
	fmt.Println("  build        - Build the project")
	fmt.Println("  run          - Run Go files in the current directory (max 10 arguments)")
	fmt.Println("  run --1      - With specific file or module")
	fmt.Println("  --v          - Display the version information")
	fmt.Println("  --h          - Display this help message")
	fmt.Println("Examples:")
	fmt.Println("1. goroot new my-project")
	fmt.Println("   cd my-project/src")
	fmt.Println("   goroot run")
	fmt.Println("   goroot build")
	fmt.Println("   Get your compiled binary at my-project/target/release")
	fmt.Println("2. goroot init - This command creates go.mod in the current directory")
	fmt.Println("3. goroot new - This command creates this structure: my-project/{src,go.mod,target}")
	fmt.Println("4. goroot build - This command will build your source code into a compiled binary in my-project/target/release/")
	fmt.Println("5. goroot run - This will run your program with up to 10 arguments, like: goroot run --v --1 --dep")
	fmt.Println("6. goroot run --1 file.go - This will run your specific module or file, like: goroot run --1 help.go or goroot run --1 warn.go")
}
