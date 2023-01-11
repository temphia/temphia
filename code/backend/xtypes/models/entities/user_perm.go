package entities

type Role struct {
	Slug     string `json:"slug,omitempty" db:"slug,omitempty"`
	TenantID string `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type UserRole struct {
	RoleId   string `json:"role_id,omitempty" db:"role_id,omitempty"`
	UserId   string `json:"user_id,omitempty" db:"user_id,omitempty"`
	TenantID string `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type Permission struct {
	ID         int64      `json:"id,omitempty" db:"id,omitempty"`
	ObjectType string     `json:"object_type,omitempty" db:"object_type,omitempty"`
	ObjectId   string     `json:"object_id,omitempty" db:"object_id,omitempty"`
	RoleID     string     `json:"role_id,omitempty" db:"role_id,omitempty"`
	ExtraMeta  JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantID   string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}
