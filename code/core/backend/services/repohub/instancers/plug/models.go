package plug

type PlugResponse struct {
	Agents       []string          `json:"agents,omitempty"`
	Resources    []string          `json:"resources,omitempty"`
	ErrAgents    map[string]string `json:"err_agents,omitempty"`
	ErrResources map[string]string `json:"err_resources,omitempty"`
}
