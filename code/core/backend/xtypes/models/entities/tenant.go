package entities

type Tenant struct {
	Name           string     `json:"name,omitempty" db:"name,omitempty"`
	Slug           string     `json:"slug,omitempty" db:"slug,omitempty"`
	OrgBio         string     `json:"org_bio,omitempty" db:"org_bio,omitempty"`
	RootPlug       string     `json:"root_plug,omitempty" db:"root_plug,omitempty"`
	RootAgent      string     `json:"root_agent,omitempty" db:"root_agent,omitempty"`
	RootHandler    string     `json:"root_handler,omitempty" db:"root_handler,omitempty"`
	SmtpUser       string     `json:"smtp_user,omitempty" db:"smtp_user,omitempty"`
	SmtpPass       string     `json:"smtp_pass,omitempty" db:"smtp_pass,omitempty"`
	DefaultUgroup  string     `json:"default_ugroup,omitempty" db:"default_ugroup,omitempty"`
	MasterSecret   string     `json:"master_secret,omitempty" db:"master_secret,omitempty"`
	DefaultDSource string     `json:"default_dyn,omitempty" db:"default_dyn,omitempty"`
	DisableP2P     bool       `json:"disable_p2p,omitempty" db:"disable_p2p,omitempty"`
	ExtraMeta      JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type TenantDomain struct {
	Id                  int64      `json:"id,omitempty" db:"id,omitempty"`
	Name                string     `json:"name,omitempty" db:"name,omitempty"`
	About               string     `json:"about,omitempty" db:"about,omitempty"`
	DefaultUgroup       string     `json:"default_ugroup,omitempty" db:"default_ugroup,omitempty"`
	TemplateBprint      string     `json:"template_bprint,omitempty" db:"template_bprint,omitempty"`
	ChildTemplateBprint string     `json:"child_template_bprint,omitempty" db:"child_template_bprint,omitempty"`
	AdapterType         string     `json:"renderer_type,omitempty" db:"renderer_type,omitempty"` // dynamic, plug_app, landing_page, static, alias, launcher
	AdapterPlugId       string     `json:"renderer_plug_id,omitempty" db:"renderer_plug_id,omitempty"`
	AdapterAgentId      string     `json:"renderer_agent_id,omitempty" db:"renderer_agent_id,omitempty"`
	AdapterOptions      JsonStrMap `json:"renderer_opts,omitempty" db:"renderer_opts,omitempty"`
	EditorAgentId       string     `json:"editor_agent_id,omitempty" db:"editor_agent_id,omitempty"`
	ServeSource         string     `json:"serve_source,omitempty" db:"serve_source,omitempty"`
	ServeFolder         string     `json:"serve_folder,omitempty" db:"serve_folder,omitempty"`
	TenantId            string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	ExtraMeta           JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type DomainWidget struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Type      string     `json:"type,omitempty" db:"type,omitempty"`
	Plug      string     `json:"plug,omitempty" db:"plug,omitempty"`
	Agent     string     `json:"agent,omitempty" db:"agent,omitempty"`
	ExecMeta  JsonMap    `json:"exec_meta,omitempty" db:"exec_meta,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
	DomainId  int64      `json:"domain_id,omitempty" db:"domain_id,omitempty"`
}
