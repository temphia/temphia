package entities

type BPrint struct {
	ID          string    `json:"id,omitempty" db:"id,omitempty"`
	Name        string    `json:"name,omitempty" db:"name,omitempty"`
	Slug        string    `json:"slug,omitempty" db:"slug,omitempty"`
	Type        string    `json:"type,omitempty" db:"type,omitempty"`
	SubType     string    `json:"sub_type,omitempty" db:"sub_type,omitempty"`
	Description string    `json:"description,omitempty" db:"description,omitempty"`
	Icon        string    `json:"icon,omitempty" db:"icon,omitempty"`
	Source      string    `json:"source,omitempty" db:"source,omitempty"`
	TenantID    string    `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	Tags        JsonArray `json:"tags,omitempty" db:"tags,omitempty"`
	Files       JsonArray `json:"files,omitempty" db:"files,omitempty"`
	ExtraMeta   JsonMap   `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}
