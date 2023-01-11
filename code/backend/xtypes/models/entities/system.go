package entities

type SystemEvent struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Type      string     `json:"type,omitempty" db:"type,omitempty"`
	Data      string     `json:"data,omitempty" db:"data,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type SystemKV struct {
	Id       int64  `json:"id,omitempty" db:"id,omitempty"`
	Key      string `json:"key,omitempty" db:"key,omitempty"`
	Type     string `json:"type,omitempty" db:"type,omitempty"`
	Value    string `json:"value,omitempty" db:"value,omitempty"`
	TenantId string `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}
