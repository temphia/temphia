package plug

type PlugResponse struct {
	Agents           []string                     `json:"agents,omitempty"`
	Resources        []string                     `json:"resources,omitempty"`
	ErrAgents        map[string]string            `json:"err_agents,omitempty"`
	ErrAgentResource map[string]map[string]string `json:"err_agent_resources,omitempty"`
}

func (pr *PlugResponse) AddAgentErr(agent string, err error) {
	if pr.ErrAgents == nil {
		pr.ErrAgents = make(map[string]string)
	}

	pr.ErrAgents[agent] = err.Error()
}

func (pr *PlugResponse) AddResourceErr(agent, res string, err error) {
	if pr.ErrAgentResource == nil {
		pr.ErrAgentResource = make(map[string]map[string]string)
	}

	ress, ok := pr.ErrAgentResource[agent]
	if !ok {
		ress = map[string]string{}
		pr.ErrAgentResource[agent] = ress
	}

	ress[res] = err.Error()
}

type PlugOptions struct {
	Id string `json:"id,omitempty"`
}
