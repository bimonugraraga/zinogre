package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bimonugraraga/zinogre/generator"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "new" {
		fmt.Println("Usage: zinogre new <project-name or module-path>")
		return
	}

	// full module path from CLI argument
	modulePath := os.Args[2]

	// folder name = last element of module path
	folderName := filepath.Base(modulePath)

	// create project folder with templates
	err := generator.Create(folderName, modulePath)
	if err != nil {
		panic(err)
	}

	// initialize go.mod inside the folder
	cmd := exec.Command("go", "mod", "init", modulePath)
	cmd.Dir = folderName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Project created: %s (module: %s)\n", folderName, modulePath)
}
