package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	FormatV1     = "viz8/v1"
	DefaultStyle = "solid"
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
	if spec.Format != FormatV1 {
		return nil, fmt.Errorf("unsupported format: %q (expected %q)", spec.Format, FormatV1)
	}
	for i := range spec.Connections {
		c := &spec.Connections[i]
		if c.Style == "" && c.Type != "" {
			if t, ok := spec.Types[c.Type]; ok && t.Style != "" {
				c.Style = t.Style
			}
		}
		if c.Style == "" {
			c.Style = DefaultStyle
		}
	}
	return &spec, nil
}
