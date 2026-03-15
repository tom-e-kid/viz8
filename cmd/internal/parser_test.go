package internal

import (
	"testing"
)

func TestParse_valid(t *testing.T) {
	yaml := `
format: viz8/v1
title: Test
groups:
  - id: g1
    label: Group1
components:
  - id: c1
    label: Comp1
    group: g1
`
	spec, err := Parse([]byte(yaml))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec.Title != "Test" {
		t.Errorf("title = %q, want %q", spec.Title, "Test")
	}
	if len(spec.Groups) != 1 {
		t.Errorf("groups = %d, want 1", len(spec.Groups))
	}
	if len(spec.Components) != 1 {
		t.Errorf("components = %d, want 1", len(spec.Components))
	}
}

func TestParse_unsupportedFormat(t *testing.T) {
	yaml := `format: viz8/v2`
	_, err := Parse([]byte(yaml))
	if err == nil {
		t.Fatal("expected error for unsupported format")
	}
}

func TestParse_missingFormat(t *testing.T) {
	yaml := `title: No Format`
	_, err := Parse([]byte(yaml))
	if err == nil {
		t.Fatal("expected error for missing format")
	}
}

func TestParse_invalidYAML(t *testing.T) {
	_, err := Parse([]byte(`{{{`))
	if err == nil {
		t.Fatal("expected error for invalid YAML")
	}
}

func TestParse_connectionStyleFromType(t *testing.T) {
	yaml := `
format: viz8/v1
types:
  async:
    style: dashed
groups:
  - id: g1
components:
  - id: a
    group: g1
  - id: b
    group: g1
connections:
  - from: a
    to: b
    type: async
`
	spec, err := Parse([]byte(yaml))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec.Connections[0].Style != "dashed" {
		t.Errorf("style = %q, want %q", spec.Connections[0].Style, "dashed")
	}
}

func TestParse_connectionExplicitStyleOverridesType(t *testing.T) {
	yaml := `
format: viz8/v1
types:
  async:
    style: dashed
groups:
  - id: g1
components:
  - id: a
    group: g1
  - id: b
    group: g1
connections:
  - from: a
    to: b
    type: async
    style: dotted
`
	spec, err := Parse([]byte(yaml))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec.Connections[0].Style != "dotted" {
		t.Errorf("style = %q, want %q", spec.Connections[0].Style, "dotted")
	}
}

func TestParse_connectionTypeWithNoStyle(t *testing.T) {
	yaml := `
format: viz8/v1
types:
  info:
    label: INFO
    color: "#38bdf8"
groups:
  - id: g1
components:
  - id: a
    group: g1
  - id: b
    group: g1
connections:
  - from: a
    to: b
    type: info
`
	spec, err := Parse([]byte(yaml))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec.Connections[0].Style != "solid" {
		t.Errorf("style = %q, want %q", spec.Connections[0].Style, "solid")
	}
}

func TestParse_connectionUnknownType(t *testing.T) {
	yaml := `
format: viz8/v1
groups:
  - id: g1
components:
  - id: a
    group: g1
  - id: b
    group: g1
connections:
  - from: a
    to: b
    type: nonexistent
`
	spec, err := Parse([]byte(yaml))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec.Connections[0].Style != "solid" {
		t.Errorf("style = %q, want %q", spec.Connections[0].Style, "solid")
	}
}

func TestParse_connectionDefaultStyle(t *testing.T) {
	yaml := `
format: viz8/v1
groups:
  - id: g1
components:
  - id: a
    group: g1
  - id: b
    group: g1
connections:
  - from: a
    to: b
`
	spec, err := Parse([]byte(yaml))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec.Connections[0].Style != "solid" {
		t.Errorf("style = %q, want %q", spec.Connections[0].Style, "solid")
	}
}
