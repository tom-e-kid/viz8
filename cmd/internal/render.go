package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Render converts a Spec to an HTML file and opens it in the browser.
func Render(spec *Spec, templateHTML string) error {
	html, err := renderHTML(spec, templateHTML)
	if err != nil {
		return err
	}

	f, err := os.CreateTemp("", "viz8-*.html")
	if err != nil {
		return fmt.Errorf("creating temp file: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(html); err != nil {
		return fmt.Errorf("writing temp file: %w", err)
	}

	return openBrowser(f.Name())
}

func renderHTML(spec *Spec, templateHTML string) (string, error) {
	data, err := json.Marshal(spec)
	if err != nil {
		return "", fmt.Errorf("marshaling JSON: %w", err)
	}

	html := strings.Replace(templateHTML, "__DATA__", string(data), 1)
	return html, nil
}

func openBrowser(path string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	go cmd.Wait()
	return nil
}
