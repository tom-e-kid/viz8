---
description: "Visualize a system structure as an interactive diagram using viz8"
argument-hint: "<target structure> (optional) — e.g., 'authentication flow', 'API layer'"
---

Visualize a system structure from the current codebase as an interactive HTML/SVG diagram.

## Steps

### Step 1: Get the viz8/v1 format specification

Run `viz8 schema` using the Bash tool and read the output carefully.
This is the YAML format you must produce.

### Step 2: Identify the target structure

**If arguments are provided** (`$ARGUMENTS`):
- Use the argument as the target structure description

**If no arguments**:
- Analyze the current codebase context (recent conversation, open files, project structure)
- Identify a meaningful structure to visualize (e.g., system architecture, data flow, module dependencies)

### Step 3: Confirm with the user

Present the identified structure to the user:
- What will be visualized (scope and components)
- How it will be organized (groups, connections)

Ask: "This is the structure I'll visualize. OK to proceed?"

**MUST NOT proceed until the user confirms.**

### Step 4: Generate viz8/v1 YAML

Based on the confirmed structure and the schema from Step 1:

1. Analyze the codebase to extract the relevant components, relationships, and groupings
2. Generate a valid viz8/v1 YAML document
3. Use meaningful `types` to distinguish different kinds of connections or items
4. Use `groups` to organize components into logical categories
5. Add `description` fields to components for additional context

### Step 5: Save the YAML file

1. Determine the filename: `m-<yyyyMMdd>.yaml` using today's date
2. If `.viz8/output/m-<yyyyMMdd>.yaml` already exists, append a suffix: `m-<yyyyMMdd>-2.yaml`, `m-<yyyyMMdd>-3.yaml`, etc.
3. Create the `.viz8/output/` directory if it doesn't exist: `mkdir -p .viz8/output`
4. Write the YAML to `.viz8/output/<filename>.yaml`

### Step 6: Open in browser

Run `viz8 .viz8/output/<filename>.yaml` using the Bash tool.
The command will:
- Generate an HTML file alongside the YAML
- Open it in the default browser
- Print the `file://` URL to stdout

Display the URL to the user.
