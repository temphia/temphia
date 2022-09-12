package bindx

type Value struct {
	Value    string `json:"value,omitempty"`
	Audience string `json:"audience,omitempty"`
	Version  int64  `json:"version,omitempty"`
	TTL      int64  `json:"ttl,omitempty"`
}

type Resource struct {
	Name    string            `json:"name,omitempty"`
	Type    string            `json:"type,omitempty"`
	Payload string            `json:"payload,omitempty"`
	Meta    map[string]string `json:"meta,omitempty"`
}

type Link struct {
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	PlugId  string `json:"plug_id,omitempty"`
	AgentId string `json:"agent_id,omitempty"`
}

type CabTicket struct {
	Prefix      string   `json:"prefix,omitempty"`
	PinnedFiles []string `json:"pinned_files,omitempty"`
	Operations  []string `json:"ops,omitempty"`
}

type HttpRequest struct {
	Method  string            `json:"method,omitempty"`
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    []byte            `json:"body,omitempty"`
}

type HttpResponse struct {
	SatusCode int                 `json:"status_code,omitempty"`
	Headers   map[string][]string `json:"headers,omitempty"`
	Json      bool                `json:"json,omitempty"`
	Body      []byte              `json:"body,omitempty"`
}
