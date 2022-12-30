package entities

type UserDevice struct {
	Id         int64      `json:"id,omitempty" db:"id,omitempty"`
	Name       string     `json:"name,omitempty" db:"name,omitempty"`
	UserId     string     `json:"user_id,omitempty" db:"user_id,omitempty"`
	DeviceType string     `json:"device_type,omitempty" db:"device_type,omitempty"`
	LastAddr   string     `json:"last_addr,omitempty" db:"last_addr,omitempty"`
	APNToken   string     `json:"apn_token,omitempty" db:"apn_token,omitempty"`
	Scopes     string     `json:"scopes,omitempty" db:"scopes,omitempty"`
	ExtraMeta  JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantID   string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}
