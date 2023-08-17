package xpackage

type NewResource struct {
	Name      string            `json:"name,omitempty"`
	Type      string            `json:"type,omitempty"`
	SubType   string            `json:"sub_type,omitempty"`
	Payload   string            `json:"payload,omitempty"`
	Policy    string            `json:"policy,omitempty"`
	TargetRef string            `json:"target_ref,omitempty"`
	Meta      map[string]string `json:"meta,omitempty"`
}
