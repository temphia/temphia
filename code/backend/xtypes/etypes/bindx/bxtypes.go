package bindx

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

type HttpRequest struct {
	Method  string            `json:"method,omitempty"`
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    []byte            `json:"body,omitempty"`
}

type HttpResponse struct {
	SatusCode int                 `json:"status_code,omitempty"`
	Headers   map[string][]string `json:"headers,omitempty"`
	Body      []byte              `json:"body,omitempty"`
}

type UserMessage struct {
	Title            string `json:"title,omitempty"`
	Contents         string `json:"contents,omitempty"`
	Encrypted        bool   `json:"encrypted,omitempty"`
	UsingCurrentUser bool   `json:"using_current_user,omitempty"`
}
