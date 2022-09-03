package entities

type UserGroup struct {
	Name           string     `json:"name,omitempty" db:"name"`
	Slug           string     `json:"slug,omitempty" db:"slug"`
	Icon           string     `json:"icon,omitempty" db:"icon"`
	EnablePassAuth bool       `json:"enable_pass_auth" db:"enable_pass_auth"`
	Scopes         string     `json:"scopes,omitempty" db:"scopes"`
	OpenSignUp     bool       `json:"open_sign_up" db:"open_sign_up"`
	TenantID       string     `json:"tenant_id,omitempty" db:"tenant_id"`
	ExtraMeta      JsonStrMap `json:"extra_meta" db:"extra_meta,omitempty"`
}
