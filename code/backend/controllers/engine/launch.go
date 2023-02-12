package engine

type ExecInstanceOptions struct {
	ApiBaseURL   string            `json:"api_base_url,omitempty"`
	Token        string            `json:"token,omitempty"`
	EntryName    string            `json:"entry,omitempty"`
	ExecLoader   string            `json:"exec_loader,omitempty"`
	JSPlugScript string            `json:"js_plug_script,omitempty"`
	StyleFile    string            `json:"style,omitempty"`
	ExtScripts   map[string]string `json:"ext_scripts,omitempty"`
	Plug         string            `json:"plug,omitempty"`
	Agent        string            `json:"agent,omitempty"`
}
