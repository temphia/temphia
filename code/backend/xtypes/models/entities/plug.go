package entities

type Plug struct {
	Id   string `json:"id,omitempty" db:"id,omitempty"`
	Name string `json:"name,omitempty" db:"name,omitempty"`
	Live bool   `json:"live,omitempty" db:"live,omitempty"`
	Dev  bool   `json:"dev,omitempty"  db:"dev,omitempty"`

	BprintId         string     `json:"bprint_id,omitempty"  db:"bprint_id,omitempty"`
	BprintItemId     string     `json:"bprint_item_id,omitempty"  db:"bprint_item_id,omitempty"`
	InstancedObjects JsonStrMap `json:"instanced_objects,omitempty" db:"instanced_objects,omitempty"`
	StepHead         string     `json:"step_head,omitempty"  db:"step_head,omitempty"`

	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type Agent struct {
	Id         string     `json:"id,omitempty" db:"id,omitempty"`
	Name       string     `json:"name,omitempty" db:"name,omitempty"`
	Type       string     `json:"type,omitempty" db:"type,omitempty"`         // none, era spa gossr static
	Executor   string     `json:"executor,omitempty" db:"executor,omitempty"` // none, javascript1, wasm1
	IfaceFile  string     `json:"iface_file,omitempty" db:"iface_file,omitempty"`
	EntryFile  string     `json:"entry_file,omitempty" db:"entry_file,omitempty"`
	WebFiles   JsonStrMap `json:"web_files,omitempty" db:"web_files,omitempty"`
	WebOptions JsonStrMap `json:"web_options,omitempty" db:"web_options,omitempty"`
	ExtraMeta  JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	ModVersion int64      `json:"mod_version,omitempty" db:"mod_version,omitempty"`
	PlugId     string     `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	TenantId   string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type SPAOptions struct {
	WebEntry      string `json:"web_entry,omitempty"`
	WebScript     string `json:"web_script,omitempty"`
	WebStyle      string `json:"web_style,omitempty"`
	WebLoader     string `json:"web_loader,omitempty"`
	ServiceWorker string `json:"service_worker,omitempty"`
}

type SSROptions struct {
	AutoInject     string `json:"auto_inject,omitempty"`
	InjectKey      string `json:"inject_key,omitempty"`
	InjectFiles    string `json:"inject_files,omitempty"`
	ActionRedirect string `json:"action_redirect,omitempty"`
	ServiceWorker  string `json:"service_worker,omitempty"`
}

type AgentResource struct {
	Slug       string `json:"slug,omitempty" db:"slug,omitempty"`
	PlugId     string `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	AgentId    string `json:"agent_id,omitempty" db:"agent_id,omitempty"`
	ResourceId string `json:"resource_id,omitempty" db:"resource_id,omitempty"`
	Actions    string `json:"actions,omitempty" db:"actions,omitempty"`
	Policy     string `json:"policy,omitempty" db:"policy,omitempty"`
	TenantId   string `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type AgentLink struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	FromPlug  string     `json:"from_plug_id,omitempty" db:"from_plug_id,omitempty"`
	FromAgent string     `json:"from_agent_id,omitempty" db:"from_agent_id,omitempty"`
	ToPlug    string     `json:"to_plug_id,omitempty" db:"to_plug_id,omitempty"`
	ToAgent   string     `json:"to_agent_id,omitempty" db:"to_agent_id,omitempty"`
	ToHandler string     `json:"to_handler,omitempty" db:"to_handler,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type AgentExtension struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Plug      string     `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	Agent     string     `json:"agent_id,omitempty" db:"agent_id,omitempty"`
	RefFile   string     `json:"ref_file,omitempty" db:"ref_file,omitempty"`
	BprintId  string     `json:"bprint_id,omitempty" db:"bprint_id,omitempty"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId  string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}
