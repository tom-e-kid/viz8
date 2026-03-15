# viz8

Architecture and system visualization tool. Parses YAML specifications (`viz8/v1` format) and renders interactive HTML/SVG diagrams in the browser.

## Features

- Dark / Light theme toggle
- Horizontal / Vertical layout toggle
- Pan & zoom with mouse
- Click components to highlight connections
- Filter by group

## Install

Requires [mise](https://mise.jdx.dev/) and Go.

```bash
mise run install
```

This builds the binary and installs it to `~/.local/bin/viz8`.

## Usage

```bash
viz8 spec.yaml
```

The tool opens the rendered diagram in your default browser.

## Format

See [cmd/docs/viz8-v1.md](cmd/docs/viz8-v1.md) for the full specification.

```yaml
format: viz8/v1
title: My System
description: Overview

groups:
  - id: frontend
    label: Frontend
    color: "#3b82f6"

components:
  - id: app
    label: Web App
    group: frontend
    items:
      - label: Router
      - label: Store

connections:
  - from: app
    to: api
    label: HTTPS
    style: solid  # solid | dashed | dotted
```

## Development

```bash
mise run build    # Build to bin/viz8
mise run install  # Build and install to ~/.local/bin/viz8
```
