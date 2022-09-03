package instance

import "github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

type Plug struct {
	NewPlugId   string `json:"new_plug_id,omitempty"`
	NewPlugName string `json:"new_plug_name,omitempty"`
	//Schema       string                        `json:"schema,omitempty"`
	AgentOptions map[string]*AgentOptions      `json:"agent_opts,omitempty"`
	Resources    map[string]*entities.Resource `json:"resources,omitempty"`
}

type PlugResponse struct {
	Agents       []string          `json:"agents,omitempty"`
	Resources    []string          `json:"resources,omitempty"`
	ErrAgents    map[string]string `json:"err_agents,omitempty"`
	ErrResources map[string]string `json:"err_resources,omitempty"`
}

type AgentOptions struct {
	Name      string            `json:"name,omitempty"`
	Resources map[string]string `json:"resources,omitempty"`
}
