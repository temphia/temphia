package entities

type Tenant struct {
	Name           string     `json:"name,omitempty" db:"name,omitempty"`
	Slug           string     `json:"slug,omitempty" db:"slug,omitempty"`
	OrgBio         string     `json:"org_bio,omitempty" db:"org_bio,omitempty"`
	DefaultDSource string     `json:"default_dyn,omitempty" db:"default_dyn,omitempty"`
	DefaultUgroup  string     `json:"default_ugroup,omitempty" db:"default_ugroup,omitempty"`
	DisableP2P     bool       `json:"disable_p2p,omitempty" db:"disable_p2p,omitempty"`
	MasterSecret   string     `json:"master_secret,omitempty" db:"master_secret,omitempty"`
	ExtraMeta      JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type TenantDomain struct {
	Id             int64      `json:"id,omitempty" db:"id,omitempty"`
	Name           string     `json:"name,omitempty" db:"name,omitempty"`
	About          string     `json:"about,omitempty" db:"about,omitempty"`
	DefaultUgroup  string     `json:"default_ugroup,omitempty" db:"default_ugroup,omitempty"`
	CORSPolicy     string     `json:"cors_policy,omitempty" db:"cors_policy,omitempty"`
	AdapterPolicy  string     `json:"adapter_policy,omitempty" db:"adapter_policy,omitempty"`
	AdapterType    string     `json:"adapter_type,omitempty" db:"adapter_type,omitempty"` // dynamic, plug_app, landing_page, static, alias, launcher
	AdapterOptions JsonStrMap `json:"adapter_opts,omitempty" db:"adapter_opts,omitempty"`
	TenantId       string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExtraMeta      JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}
