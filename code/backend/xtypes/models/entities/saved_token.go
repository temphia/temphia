package entities

import "time"

type SavedToken struct {
	Id        string     `json:"id,omitempty" db:"id,omitempty"`
	Type      string     `json:"type,omitempty" db:"type,omitempty"`
	UserId    string     `json:"user_id,omitempty" db:"user_id,omitempty"`
	UserGroup string     `json:"user_group,omitempty" db:"user_group,omitempty"`
	Target    string     `json:"target,omitempty" db:"target,omitempty"`
	Payload   string     `json:"payload,omitempty" db:"payload,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExpiresOn *time.Time `json:"expires_on,omitempty" db:"expires_on,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}
