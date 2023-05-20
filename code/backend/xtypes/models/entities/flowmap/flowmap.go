package flowmap

type Data struct {
	Plugs          []*Plug          `json:"plugs,omitempty"`
	Agents         []*Agent         `json:"agents,omitempty"`
	AgentLinks     []*AgentLink     `json:"agent_links,omitempty"`
	AgentExts      []*AgentExt      `json:"agent_exts,omitempty"`
	TargetApps     []*App           `json:"target_apps,omitempty"`
	TargetHooks    []*Hook          `json:"target_hooks,omitempty"`
	AgentResources []*AgentResource `json:"agent_resources,omitempty"`
}

type Plug struct {
	Id       string `json:"id,omitempty" db:"id,omitempty"`
	Name     string `json:"name,omitempty" db:"name,omitempty"`
	BprintId string `json:"bprint_id,omitempty"  db:"bprint_id,omitempty"`
}

type Agent struct {
	Id     string `json:"id,omitempty" db:"id,omitempty"`
	Name   string `json:"name,omitempty" db:"name,omitempty"`
	Type   string `json:"type,omitempty" db:"type,omitempty"`
	PlugId string `json:"plug_id,omitempty" db:"plug_id,omitempty"`
}

type AgentLink struct {
	Id        int64  `json:"id,omitempty" db:"id,omitempty"`
	Name      string `json:"name,omitempty" db:"name,omitempty"`
	FromPlug  string `json:"from_plug_id,omitempty" db:"from_plug_id,omitempty"`
	FromAgent string `json:"from_agent_id,omitempty" db:"from_agent_id,omitempty"`
	ToPlug    string `json:"to_plug_id,omitempty" db:"to_plug_id,omitempty"`
	ToAgent   string `json:"to_agent_id,omitempty" db:"to_agent_id,omitempty"`
}

type AgentExt struct {
	Id    int64  `json:"id,omitempty" db:"id,omitempty"`
	Name  string `json:"name,omitempty" db:"name,omitempty"`
	Plug  string `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	Agent string `json:"agent_id,omitempty" db:"agent_id,omitempty"`
}

type App struct {
	Id         int64  `json:"id,omitempty" db:"id,omitempty"`
	Name       string `json:"name,omitempty" db:"name,omitempty"`
	TargetType string `json:"target_type,omitempty" db:"target_type,omitempty"`
	Target     string `json:"target,omitempty" db:"target,omitempty"`
	PlugId     string `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	AgentId    string `json:"agent_id,omitempty" db:"agent_id,omitempty"`
}

type Hook struct {
	Id         int64  `json:"id,omitempty" db:"id,omitempty"`
	Name       string `json:"name,omitempty" db:"name,omitempty"`
	TargetType string `json:"target_type,omitempty" db:"target_type,omitempty"`
	Target     string `json:"target,omitempty" db:"target,omitempty"`
	PlugId     string `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	AgentId    string `json:"agent_id,omitempty" db:"agent_id,omitempty"`
}

type AgentResource struct {
	Slug       string `json:"slug,omitempty" db:"slug,omitempty"`
	PlugId     string `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	AgentId    string `json:"agent_id,omitempty" db:"agent_id,omitempty"`
	ResourceId string `json:"resource_id,omitempty" db:"resource_id,omitempty"`
}
