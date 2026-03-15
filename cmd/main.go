package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/tom-e-kid/viz8/internal"
)

//go:embed web/template.html
var templateHTML string

//go:embed docs/viz8-v1.md
var schemaDoc string

const version = "0.1.0"

var subcommands = map[string]bool{
	"help": true, "schema": true, "version": true,
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	arg := os.Args[1]

	switch arg {
	case "help":
		if len(os.Args) >= 3 {
			printCommandHelp(os.Args[2])
		} else {
			printUsage()
		}
	case "schema":
		fmt.Print(schemaDoc)
	case "version":
		fmt.Printf("viz8 %s\n", version)
	default:
		if strings.HasPrefix(arg, "-") || subcommands[arg] == false && !looksLikeFile(arg) {
			fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", arg)
			printUsage()
			os.Exit(1)
		}
		runOpen(arg)
	}
}

// looksLikeFile returns true if the argument looks like a file path.
func looksLikeFile(s string) bool {
	return strings.ContainsAny(s, "./")
}

func runOpen(path string) {
	spec, err := internal.ParseFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	htmlPath, err := internal.Render(spec, templateHTML, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := internal.OpenBrowser(htmlPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("file://%s\n", htmlPath)
}

func printUsage() {
	fmt.Print(`viz8 - Architecture and system visualization tool

Usage:
  viz8 <file.yaml>    Open a viz8/v1 YAML file in the browser
  viz8 schema         Print the viz8/v1 format specification
  viz8 help [command] Show help for a command
  viz8 version        Print version

Examples:
  viz8 arch.yaml                     # Visualize a YAML spec
  viz8 .viz8/output/m-20260315.yaml  # Open from .viz8 output dir
  viz8 schema                        # Show format spec (useful for AI agents)
`)
}

func printCommandHelp(command string) {
	switch command {
	case "schema":
		fmt.Print(`viz8 schema - Print the viz8/v1 format specification

Usage:
  viz8 schema

Prints the complete viz8/v1 YAML format specification to stdout.
This is useful for AI agents and tools that need to understand
the expected YAML structure to generate valid viz8 input files.

The specification includes all field definitions, defaults,
constraints, and a complete example.
`)
	case "version":
		fmt.Print(`viz8 version - Print version information

Usage:
  viz8 version
`)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}
