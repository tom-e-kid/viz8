package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ParseFile reads a YAML file and returns a Spec.
func ParseFile(path string) (*Spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	return Parse(data)
}

// Parse decodes YAML bytes into a Spec.
func Parse(data []byte) (*Spec, error) {
	var spec Spec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("parsing YAML: %w", err)
	}
	if spec.Format != "viz8/v1" {
		return nil, fmt.Errorf("unsupported format: %q (expected \"viz8/v1\")", spec.Format)
	}
	// Default connection style to "solid".
	for i := range spec.Connections {
		if spec.Connections[i].Style == "" {
			spec.Connections[i].Style = "solid"
		}
	}
	return &spec, nil
}
