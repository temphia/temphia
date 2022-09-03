package entities

type DataHook struct {
	Id        int64      `json:"id,omitempty" db:"id,omitempty"`
	Name      string     `json:"name,omitempty" db:"name,omitempty"`
	Type      string     `json:"type,omitempty" db:"type,omitempty"`
	SubType   string     `json:"sub_type,omitempty" db:"sub_type,omitempty"`
	PlugId    string     `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	AgentId   string     `json:"agent_id,omitempty" db:"agent_id,omitempty"`
	Handler   string     `json:"handler,omitempty" db:"handler,omitempty"`
	Payload   string     `json:"payload,omitempty" db:"payload,omitempty"`
	TableID   string     `json:"table_id,omitempty" db:"table_id"`
	GroupID   string     `json:"group_id,omitempty" db:"group_id"`
	TenantID  string     `json:"tenant_id,omitempty" db:"tenant_id"`
	ExtraMeta JsonStrMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

/*
	types of hooks
		- selected_rows
		- selected_row
		- before_modify
		- after_modify
		- row_calulate
*/
