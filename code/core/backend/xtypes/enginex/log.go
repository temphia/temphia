package enginex

type DebugMessage struct {
	EventId  string   `json:"event_id,omitempty"`
	PlugId   string   `json:"plug_id,omitempty"`
	AgentId  string   `json:"agent_id,omitempty"`
	Messages []string `json:"messages,omitempty"`
}
