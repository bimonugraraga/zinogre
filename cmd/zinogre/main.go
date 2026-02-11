package main

import (
	"fmt"
	"os"

	"github.com/bimonugraraga/zinogre/generator"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  zinogre new <project-name>")
		return
	}

	command := os.Args[1]

	switch command {

	case "new":
		name := os.Args[2]

		if err := generator.Create(name); err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Project created:", name)

	default:
		fmt.Println("Unknown command:", command)
	}
}
