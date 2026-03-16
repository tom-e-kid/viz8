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
	args := os.Args[1:]

	// Parse flags
	var outputDir string
	var remaining []string
	for i := 0; i < len(args); i++ {
		if args[i] == "-o" {
			if i+1 >= len(args) {
				fmt.Fprintln(os.Stderr, "Error: -o requires a directory argument")
				os.Exit(1)
			}
			outputDir = args[i+1]
			i++
		} else {
			remaining = append(remaining, args[i])
		}
	}

	if len(remaining) == 0 {
		printUsage()
		os.Exit(0)
	}

	arg := remaining[0]

	switch arg {
	case "help":
		if len(remaining) >= 2 {
			printCommandHelp(remaining[1])
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
		runOpen(arg, outputDir)
	}
}

// looksLikeFile returns true if the argument looks like a file path.
func looksLikeFile(s string) bool {
	return strings.ContainsAny(s, "./")
}

func runOpen(path string, outputDir string) {
	spec, err := internal.ParseFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	htmlPath, err := internal.Render(spec, templateHTML, path, outputDir)
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
	defaultDir, err := internal.DefaultOutputDir()
	if err != nil {
		defaultDir = "~/.viz8/output"
	}
	fmt.Printf(`viz8 - Architecture and system visualization tool

Usage:
  viz8 [options] <file.yaml>  Open a viz8/v1 YAML file in the browser
  viz8 schema                 Print the viz8/v1 format specification
  viz8 help [command]         Show help for a command
  viz8 version                Print version

Options:
  -o <dir>  Output directory for generated HTML (default: %s)

Examples:
  viz8 arch.yaml             # Visualize a YAML spec
  viz8 -o ./out arch.yaml    # Output HTML to ./out/
  viz8 schema                # Show format spec (useful for AI agents)
`, defaultDir)
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
