// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		DisplayHelp()
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
	case "run":
		if err := handleRun(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "new":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a folder name")
			os.Exit(1)
		}
		folderName := os.Args[2]
		if err := handleNew(folderName); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "--v":
		DisplayVersion()
		os.Exit(0)
	case "--h":
		DisplayHelp()
		os.Exit(0)
	default:
		DisplayHelp()
		os.Exit(1)
	}
}
