package entities

type UserGroupAuth struct {
	Id                int64      `json:"id,omitempty" db:"id,omitempty"`
	Name              string     `json:"name,omitempty" db:"name,omitempty"`
	Type              string     `json:"type,omitempty" db:"type,omitempty"`
	Provider          string     `json:"provider,omitempty" db:"provider,omitempty"`
	AuthURL           string     `json:"auth_url,omitempty" db:"auth_url,omitempty"`
	TokenURL          string     `json:"token_url,omitempty" db:"token_url,omitempty"`
	ClientId          string     `json:"client_id,omitempty" db:"client_id,omitempty"`
	ClientSecret      string     `json:"client_secret,omitempty" db:"client_secret,omitempty"`
	Scopes            string     `json:"scopes,omitempty" db:"scopes,omitempty"`
	NewUserIfNotExist bool       `json:"newuser_ifnot_exists,omitempty" db:"newuser_ifnot_exists,omitempty"`
	Payload           string     `json:"payload,omitempty" db:"payload,omitempty"`
	Policy            string     `json:"policy,omitempty"  db:"policy,omitempty"`
	UserGroup         string     `json:"user_group,omitempty" db:"user_group,omitempty"`
	TenantId          string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExtraMeta         JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type UserGroupPlug struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	PlugId    string     `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	AgentId   string     `json:"agent_id,omitempty" db:"agent_id,omitempty"`
	Icon      string     `json:"icon,omitempty" db:"icon,omitempty"`
	Policy    string     `json:"policy,omitempty" db:"policy,omitempty"`
	UserGroup string     `json:"user_group,omitempty" db:"user_group,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type UserGroupHook struct {
	Id         int64      `json:"id,omitempty" db:"id,omitempty"`
	Type       string     `json:"type,omitempty" db:"type,omitempty"`
	Target     string     `json:"target,omitempty" db:"target,omitempty"`
	Data       string     `json:"data,omitempty" db:"data,omitempty"`
	ClientSide bool       `json:"client_side,omitempty" db:"client_side,omitempty"`
	PlugId     string     `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	AgentId    string     `json:"agent_id,omitempty" db:"agent_id,omitempty"`
	UserGroup  string     `json:"user_group,omitempty" db:"user_group,omitempty"`
	TenantId   string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExtraMeta  JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type UserGroupData struct {
	Id         int64      `json:"id,omitempty" db:"id,omitempty"`
	DataSource string     `json:"data_source,omitempty" db:"data_source,omitempty"`
	DataGroup  string     `json:"data_group,omitempty" db:"data_group,omitempty"`
	Policy     string     `json:"policy,omitempty" db:"policy,omitempty"`
	UserGroup  string     `json:"user_group,omitempty" db:"user_group,omitempty"`
	TenantId   string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExtraMeta  JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}
