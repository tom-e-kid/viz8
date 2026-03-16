package internal

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestOutputPath(t *testing.T) {
	tests := []struct {
		input     string
		outputDir string
		want      string
	}{
		{"model.yaml", "/tmp/out", "/tmp/out/model.html"},
		{".viz8/output/m-20260315.yaml", "/tmp/out", "/tmp/out/m-20260315.html"},
		{"path/to/spec.yml", "/tmp/out", "/tmp/out/spec.html"},
		{"file.json", "/tmp/out", "/tmp/out/file.html"},
	}
	for _, tt := range tests {
		got, err := outputPath(tt.input, tt.outputDir)
		if err != nil {
			t.Errorf("outputPath(%q, %q) error: %v", tt.input, tt.outputDir, err)
			continue
		}
		if got != tt.want {
			t.Errorf("outputPath(%q, %q) = %q, want %q", tt.input, tt.outputDir, got, tt.want)
		}
	}
}

func TestOutputPath_defaultDir(t *testing.T) {
	defaultDir, err := DefaultOutputDir()
	if err != nil {
		t.Fatalf("DefaultOutputDir() error: %v", err)
	}
	want := filepath.Join(defaultDir, "model.html")
	got, err := outputPath("model.yaml", "")
	if err != nil {
		t.Fatalf("outputPath error: %v", err)
	}
	if got != want {
		t.Errorf("outputPath with default dir = %q, want %q", got, want)
	}
}

func TestOutputPath_noExtension(t *testing.T) {
	_, err := outputPath("noext", "/tmp/out")
	if err == nil {
		t.Fatal("expected error for file without extension")
	}
}

func TestRenderHTML(t *testing.T) {
	spec := &Spec{
		Format: FormatV1,
		Title:  "Test",
		Groups: []Group{{ID: "g1", Label: "G1"}},
	}
	template := `<html>__DATA__</html>`

	html, err := renderHTML(spec, template)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.Contains(html, "__DATA__") {
		t.Error("template placeholder was not replaced")
	}
	if !strings.Contains(html, `"title":"Test"`) {
		t.Error("rendered HTML does not contain spec data")
	}
	if !strings.Contains(html, "<html>") {
		t.Error("rendered HTML does not contain template structure")
	}
}
