package bprints

type PlugNew struct {
	Slug          string                  `json:"slug,omitempty"`
	Name          string                  `json:"name,omitempty"`
	ResourceHints map[string]ResourceHint `json:"resource_hints,omitempty"`
	AgentHints    map[string]AgentHint    `json:"agent_hints,omitempty"`
}

type AgentHint struct {
	Name      string            `json:"name,omitempty"`
	Type      string            `json:"type,omitempty"`
	Executor  string            `json:"executor,omitempty"`
	IfaceFile string            `json:"iface_file,omitempty"`
	WebEntry  string            `json:"web_entry,omitempty"`
	WebScript string            `json:"web_script,omitempty"`
	WebStyle  string            `json:"web_style,omitempty"`
	WebLoader string            `json:"web_loader,omitempty"`
	WebFiles  map[string]string `json:"web_files,omitempty"`
	Resources map[string]string `json:"resources,omitempty"`
}

type ResourceHint struct {
	Name    string            `json:"name,omitempty"`
	Type    string            `json:"type,omitempty"`
	SubType string            `json:"sub_type,omitempty"`
	Payload string            `json:"schema,omitempty"`
	Policy  string            `json:"policy,omitempty"`
	Meta    map[string]string `json:"meta,omitempty"`
}
