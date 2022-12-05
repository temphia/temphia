package entities

type TargetApp struct {
	Id          int64      `json:"id,omitempty" db:"id,omitempty"`
	Name        string     `json:"name,omitempty" db:"name,omitempty"`
	Icon        string     `json:"icon,omitempty" db:"icon,omitempty"`
	Policy      string     `json:"policy,omitempty" db:"policy,omitempty"`
	TargetType  string     `json:"target_type,omitempty" db:"target_type,omitempty"`
	Target      string     `json:"target,omitempty" db:"target,omitempty"`
	ContextType string     `json:"context_type,omitempty" db:"context_type,omitempty"`
	PlugId      string     `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	AgentId     string     `json:"agent_id,omitempty" db:"agent_id,omitempty"`
	ExecDomain  int64      `json:"exec_domain,omitempty" db:"exec_domain,omitempty"`
	ExecMeta    JsonStrMap `json:"exec_meta,omitempty" db:"exec_meta,omitempty"`
	ExtraMeta   JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId    string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

/*
	app targets

	- user_app
	- auth_app
	- domain_widget_app
	- domain_editor_app
*/

type TargetHook struct {
	Id         int64      `json:"id,omitempty" db:"id,omitempty"`
	Name       string     `json:"name,omitempty" db:"name,omitempty"`
	TargetType string     `json:"target_type,omitempty" db:"target_type,omitempty"`
	Target     string     `json:"target,omitempty" db:"target,omitempty"`
	EventType  string     `json:"event_type,omitempty" db:"event_type,omitempty"`
	Policy     string     `json:"policy,omitempty" db:"policy,omitempty"`
	PlugId     string     `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	AgentId    string     `json:"agent_id,omitempty" db:"agent_id,omitempty"`
	Handler    string     `json:"handler,omitempty" db:"handler,omitempty"`
	ExecMeta   JsonStrMap `json:"exec_meta,omitempty" db:"exec_meta,omitempty"`
	ExtraMeta  JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	TenantId   string     `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

/*
	hook_targets

	- before_auth_begin
	- before_auth_end
	- after_login
	- after_logout
	- before_password_reset
	- data_before_row_modify
	- data_after_row_modify
	- after_schema_change
	- domain_init_controller

*/
