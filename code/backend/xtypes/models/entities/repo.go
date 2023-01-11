package entities

type Repo struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Provider  string     `json:"provider,omitempty" db:"provider,omitempty"`
	URL       string     `json:"url,omitempty" db:"url,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}
