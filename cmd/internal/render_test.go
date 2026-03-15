package internal

import (
	"strings"
	"testing"
)

func TestOutputPath(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"model.yaml", "model.html"},
		{".viz8/output/m-20260315.yaml", ".viz8/output/m-20260315.html"},
		{"path/to/spec.yml", "path/to/spec.html"},
		{"file.json", "file.html"},
	}
	for _, tt := range tests {
		got, err := outputPath(tt.input)
		if err != nil {
			t.Errorf("outputPath(%q) error: %v", tt.input, err)
			continue
		}
		if got != tt.want {
			t.Errorf("outputPath(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestOutputPath_noExtension(t *testing.T) {
	_, err := outputPath("noext")
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
