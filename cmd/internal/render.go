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

// Render converts a Spec to an HTML file in the specified output directory.
// If outputDir is empty, defaults to ~/.viz8/output/.
// Returns the absolute path of the generated HTML file.
func Render(spec *Spec, templateHTML string, inputPath string, outputDir string) (string, error) {
	html, err := renderHTML(spec, templateHTML)
	if err != nil {
		return "", err
	}

	outPath, err := outputPath(inputPath, outputDir)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return "", fmt.Errorf("creating output directory: %w", err)
	}

	if err := os.WriteFile(outPath, []byte(html), 0644); err != nil {
		return "", fmt.Errorf("writing %s: %w", outPath, err)
	}

	return outPath, nil
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

// DefaultOutputDir returns the default output directory (~/.viz8/output/).
func DefaultOutputDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolving home directory: %w", err)
	}
	return filepath.Join(home, ".viz8", "output"), nil
}

// outputPath derives the HTML output path from the input filename and output directory.
// If outputDir is empty, defaults to ~/.viz8/output/.
func outputPath(inputPath string, outputDir string) (string, error) {
	ext := filepath.Ext(inputPath)
	if ext == "" {
		return "", fmt.Errorf("input file has no extension: %s", inputPath)
	}

	if outputDir == "" {
		var err error
		outputDir, err = DefaultOutputDir()
		if err != nil {
			return "", err
		}
	}

	htmlName := strings.TrimSuffix(filepath.Base(inputPath), ext) + ".html"

	absPath, err := filepath.Abs(filepath.Join(outputDir, htmlName))
	if err != nil {
		return "", fmt.Errorf("resolving path: %w", err)
	}

	return absPath, nil
}

func renderHTML(spec *Spec, templateHTML string) (string, error) {
	data, err := json.Marshal(spec)
	if err != nil {
		return "", fmt.Errorf("marshaling JSON: %w", err)
	}

	html := strings.Replace(templateHTML, "__DATA__", string(data), 1)
	return html, nil
}
