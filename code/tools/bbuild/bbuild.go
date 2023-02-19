package bbuild

type BPrint struct {
	Name        string            `json:"name,omitempty"`
	Slug        string            `json:"slug,omitempty"`
	Type        string            `json:"type,omitempty"`
	Description string            `json:"description,omitempty"`
	Icon        string            `json:"icon,omitempty"`
	Version     string            `json:"version,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Files       map[string]string `json:"files,omitempty"`
	ExtraMeta   map[string]any    `json:"extra_meta,omitempty"`
}
