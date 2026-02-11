package main

import (
	"fmt"
	"os"

	app "github.com/bimonugraraga/zinogre/internal/app/http"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  zinogre serve")
		return
	}

	switch os.Args[1] {

	case "serve":
		app.Main()

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
