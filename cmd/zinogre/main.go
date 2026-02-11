package main

import (
	"fmt"
	"os"

	"github.com/bimonugraraga/zinogre/generator"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "new" {
		fmt.Println("Usage: zinogre new <project-name>")
		return
	}

	err := generator.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
	}
}
