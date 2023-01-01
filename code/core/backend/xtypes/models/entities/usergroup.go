package entities

type UserGroup struct {
	Name        string     `json:"name,omitempty" db:"name"`
	Slug        string     `json:"slug,omitempty" db:"slug"`
	Icon        string     `json:"icon,omitempty" db:"icon"`
	Scopes      string     `json:"scopes,omitempty" db:"scopes"`
	Features    string     `json:"features,omitempty" db:"features"`
	FeatureOpts JsonStrMap `json:"feature_opts" db:"feature_opts,omitempty"`
	ExtraMeta   JsonStrMap `json:"extra_meta" db:"extra_meta,omitempty"`
	ModVersion  int64      `json:"mod_version,omitempty" db:"mod_version"`
	TenantID    string     `json:"tenant_id,omitempty" db:"tenant_id"`
}

// fixme => impl this

func (ug *UserGroup) HasFeature(feature string) bool {

	return true
	// if ug.Features == "" {
	// 	return false
	// }
	// return funk.ContainsString(strings.Split(ug.Features, ","), feature)
}

/*
	features
		sign_up
		open_sign_up
		pass_auth
		pubkey_auth
		device_key_auth

*/

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

type UserGroupData struct {
	Id         int64      `json:"id,omitempty" db:"id,omitempty"`
	DataSource string     `json:"data_source,omitempty" db:"data_source,omitempty"`
	DataGroup  string     `json:"data_group,omitempty" db:"data_group,omitempty"`
	Policy     string     `json:"policy,omitempty" db:"policy,omitempty"`
	UserGroup  string     `json:"user_group,omitempty" db:"user_group,omitempty"`
	TenantId   string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExtraMeta  JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}
