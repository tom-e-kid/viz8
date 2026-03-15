package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Render converts a Spec to an HTML file alongside the input YAML.
// Returns the absolute path of the generated HTML file.
func Render(spec *Spec, templateHTML string, inputPath string) (string, error) {
	html, err := renderHTML(spec, templateHTML)
	if err != nil {
		return "", err
	}

	outPath, err := outputPath(inputPath)
	if err != nil {
		return "", err
	}

	absPath, err := filepath.Abs(outPath)
	if err != nil {
		return "", fmt.Errorf("resolving path: %w", err)
	}

	if err := os.WriteFile(absPath, []byte(html), 0644); err != nil {
		return "", fmt.Errorf("writing %s: %w", absPath, err)
	}

	return absPath, nil
}

// OpenBrowser opens the given file path in the default browser.
func OpenBrowser(path string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	return cmd.Run()
}

// outputPath derives the HTML output path from the input YAML path.
// e.g., ".viz8/output/m-20260315.yaml" → ".viz8/output/m-20260315.html"
func outputPath(inputPath string) (string, error) {
	ext := filepath.Ext(inputPath)
	if ext == "" {
		return "", fmt.Errorf("input file has no extension: %s", inputPath)
	}
	return strings.TrimSuffix(inputPath, ext) + ".html", nil
}

func renderHTML(spec *Spec, templateHTML string) (string, error) {
	data, err := json.Marshal(spec)
	if err != nil {
		return "", fmt.Errorf("marshaling JSON: %w", err)
	}

	html := strings.Replace(templateHTML, "__DATA__", string(data), 1)
	return html, nil
}
