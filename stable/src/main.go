// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 || (os.Args[1] != "init" && os.Args[1] != "build") {
		fmt.Println("Usage: ./main <command>")
		fmt.Println("Commands: init, build")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		if err := handleInit(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "build":
		if err := handleBuild(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid command. Use 'init' or 'build'.")
		os.Exit(1)
	}
}
