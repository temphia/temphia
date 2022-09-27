package entities

type UserGroupAuth struct {
	Id                int64      `json:"id,omitempty" db:"id,omitempty"`
	Name              string     `json:"name,omitempty" db:"name,omitempty"`
	Type              string     `json:"type,omitempty" db:"type,omitempty"`
	Provider          string     `json:"provider,omitempty" db:"provider,omitempty"`
	ProviderOptions   JsonStrMap `json:"provider_opts,omitempty" db:"provider_opts,omitempty"`
	Scopes            string     `json:"scopes,omitempty" db:"scopes,omitempty"`
	NewUserIfNotExist bool       `json:"newuser_ifnot_exists,omitempty" db:"newuser_ifnot_exists,omitempty"`
	Policy            string     `json:"policy,omitempty"  db:"policy,omitempty"`
	UserGroup         string     `json:"user_group,omitempty" db:"user_group,omitempty"`
	TenantId          string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExtraMeta         JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

func (u *UserGroupAuth) AuthURL() string {
	if u.ProviderOptions == nil {
		return ""
	}
	return u.ProviderOptions["auth_url"]
}

func (u *UserGroupAuth) TokenURL() string {
	if u.ProviderOptions == nil {
		return ""
	}
	return u.ProviderOptions["token_url"]
}

func (u *UserGroupAuth) ClientId() string {
	if u.ProviderOptions == nil {
		return ""
	}
	return u.ProviderOptions["client_id"]
}

func (u *UserGroupAuth) ClientSecret() string {
	if u.ProviderOptions == nil {
		return ""
	}
	return u.ProviderOptions["client_secret"]
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

type UserGroupData struct {
	Id         int64      `json:"id,omitempty" db:"id,omitempty"`
	DataSource string     `json:"data_source,omitempty" db:"data_source,omitempty"`
	DataGroup  string     `json:"data_group,omitempty" db:"data_group,omitempty"`
	Policy     string     `json:"policy,omitempty" db:"policy,omitempty"`
	UserGroup  string     `json:"user_group,omitempty" db:"user_group,omitempty"`
	TenantId   string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExtraMeta  JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}
