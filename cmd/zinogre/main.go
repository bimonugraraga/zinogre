package main

import (
	"fmt"
	"os"

	"github.com/bimonugraraga/zinogre/generator"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: zinogre new <name>")
		return
	}

	cmd := os.Args[1]
	if cmd == "new" {
		name := os.Args[2]

		err := generator.Create(name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Project created:", name)
		return
	}

	fmt.Println("Unknown command:", cmd)
}
