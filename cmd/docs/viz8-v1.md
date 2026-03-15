# viz8/v1 Format Specification

viz8 is an architecture/system visualization tool. It parses YAML specifications following the `viz8/v1` format and renders them as interactive HTML/SVG diagrams in the browser.

## Document Structure

```yaml
format: viz8/v1        # Required. Must be exactly "viz8/v1"
title: My System       # Optional. Displayed in the header
description: Overview  # Optional. Displayed below the title
types: {}              # Optional. User-defined visual indicators
groups: []             # Required. Visual categories (columns/rows)
components: []         # Required. Cards to render
connections: []        # Optional. Directed edges between components
```

## Fields

### Type

Types define user-defined visual indicators that can be referenced by items and connections. Each type has a key (used for referencing) and visual properties.

```yaml
types:
  pk:                    # Key: referenced by item.type / connection.type
    label: PK            # Optional. Short display text (shown as badge/legend)
    color: "#fbbf24"     # Optional. CSS color for indicator/stroke
    style: solid         # Optional. Default connection style ("solid", "dashed", "dotted")
```

When a connection references a type:
- The type's `color` is used for the edge stroke and arrow (instead of the source group's color)
- The type's `style` is used as the default line style (only if the connection has no explicit `style`)

When an item references a type:
- The type's `color` is used for the row indicator (instead of the group color)
- The type's `label` is shown as a small badge next to the indicator

### Group

Groups define visual categories. In horizontal layout they become columns; in vertical layout they become rows.

```yaml
groups:
  - id: frontend       # Required. Unique identifier
    label: Frontend     # Optional. Display name
    description: ""     # Optional. Currently unused in rendering
    color: "#3b82f6"    # Optional. CSS color (hex recommended). Default: #6b7280
```

### Component

Components are cards rendered inside their assigned group.

```yaml
components:
  - id: api-gateway     # Required. Unique identifier
    label: API Gateway   # Optional. Card title
    group: backend       # Required. Must match a group id
    description: ""      # Optional. Shown in info panel on click
    items:               # Optional. Sub-elements shown as rows
      - label: /users    # Row content
        description: REST  # Optional. Right-aligned text
        type: api          # Optional. References a type key
```

### Connection

Connections are directed edges drawn between components.

```yaml
connections:
  - from: web-app       # Required. Source component id
    to: api-gateway      # Required. Target component id
    label: HTTP          # Optional. Text displayed at edge midpoint
    style: solid         # Optional. "solid", "dashed", or "dotted"
    type: sync           # Optional. References a type key
```

## Defaults

| Field | Default | Notes |
|---|---|---|
| `connection.style` | type's style, then `"solid"` | Explicit style > type-derived > "solid" |
| `title` (in UI) | `"viz8"` | |
| `group.color` (in UI) | `"#6b7280"` | |

## Constraints

- **Unknown group reference**: A component whose `group` does not match any group id is silently hidden.
- **Unknown component reference**: A connection whose `from` or `to` does not match any component id is silently dropped.
- **Unknown type reference**: An item or connection whose `type` does not match any type key is silently ignored (renders as if no type was set).
- **Duplicate ids**: Last definition wins (standard YAML behavior).
- **Circular connections**: Allowed and rendered normally.
- **Card width**: Fixed at 280px. Long labels may overflow.

## Interactive Features

The rendered HTML provides:

- **Dark/Light theme** toggle (default: dark)
- **Horizontal/Vertical layout** toggle (default: horizontal)
- **Pan & zoom** with mouse wheel and drag
- **Click** a component to highlight it and its connected edges; others dim
- **Filter by group** via header buttons
- **Info panel** showing component details on click
- **Legend** showing line styles and defined types

## Complete Example

```yaml
format: viz8/v1
title: Web Application Architecture
description: Three-tier architecture overview

types:
  api:
    label: API
    color: "#38bdf8"
  store:
    label: STR
    color: "#a78bfa"
  sync:
    label: SYNC
    color: "#10b981"
  async:
    label: ASYNC
    color: "#f59e0b"
    style: dashed

groups:
  - id: client
    label: Client
    color: "#3b82f6"
  - id: server
    label: Server
    color: "#10b981"
  - id: data
    label: Data
    color: "#f59e0b"

components:
  - id: spa
    label: SPA
    group: client
    description: Single page application built with React
    items:
      - label: Router
        description: react-router
      - label: State
        type: store
        description: zustand
  - id: api
    label: API Server
    group: server
    description: REST API server
    items:
      - label: /users
        type: api
        description: CRUD
      - label: /posts
        type: api
        description: CRUD
      - label: /auth
        type: api
        description: JWT
  - id: db
    label: PostgreSQL
    group: data
    description: Primary database
    items:
      - label: users
      - label: posts

connections:
  - from: spa
    to: api
    label: HTTPS
    type: sync
  - from: api
    to: db
    label: SQL
    type: async
```
