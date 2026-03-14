package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/tom-e-kid/viz8/internal"
)

//go:embed web/template.html
var templateHTML string

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: viz8 <file.yaml>\n")
		os.Exit(1)
	}

	path := os.Args[1]

	spec, err := internal.ParseFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := internal.Render(spec, templateHTML); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
