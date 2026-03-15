package internal

// Spec is the top-level viz8/v1 document.
type Spec struct {
	Format      string          `yaml:"format" json:"format"`
	Title       string          `yaml:"title" json:"title"`
	Description string          `yaml:"description" json:"description"`
	Types       map[string]Type `yaml:"types" json:"types"`
	Groups      []Group         `yaml:"groups" json:"groups"`
	Components  []Component     `yaml:"components" json:"components"`
	Connections []Connection    `yaml:"connections" json:"connections"`
}

// Type is a user-defined visual indicator for items and connections.
type Type struct {
	Label string `yaml:"label" json:"label"`
	Color string `yaml:"color" json:"color"`
	Style string `yaml:"style" json:"style"`
}

// Group is a visual category (column) for components.
type Group struct {
	ID          string `yaml:"id" json:"id"`
	Label       string `yaml:"label" json:"label"`
	Description string `yaml:"description" json:"description"`
	Color       string `yaml:"color" json:"color"`
}

// Component is a card rendered inside a group column.
type Component struct {
	ID          string `yaml:"id" json:"id"`
	Label       string `yaml:"label" json:"label"`
	Group       string `yaml:"group" json:"group"`
	Description string `yaml:"description" json:"description"`
	Items       []Item `yaml:"items" json:"items"`
}

// Item is a sub-element inside a component card.
type Item struct {
	Label       string `yaml:"label" json:"label"`
	Description string `yaml:"description" json:"description"`
	Type        string `yaml:"type" json:"type"`
}

// Connection is a directed edge between two components.
type Connection struct {
	From  string `yaml:"from" json:"from"`
	To    string `yaml:"to" json:"to"`
	Label string `yaml:"label" json:"label"`
	Style string `yaml:"style" json:"style"`
	Type  string `yaml:"type" json:"type"`
}
