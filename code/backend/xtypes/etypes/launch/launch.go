package launch

type Response struct {
	ApiBaseURL string `json:"api_base_url,omitempty"`
	Token      string `json:"token,omitempty"`
	Domain     string `json:"domain,omitempty"`
	PlugId     string `json:"plug_id,omitempty"`
	AgentId    string `json:"agent_id,omitempty"`
	TenantId   string `json:"tenant_id,omitempty"`
	AuthType   string `json:"auth_type,omitempty"`
	StartPage  string `json:"start_page,omitempty"`
}

type EraOptions struct {
	EntryName  string            `json:"entry,omitempty"`
	ExecLoader string            `json:"exec_loader,omitempty"`
	ScriptFile string            `json:"script_file,omitempty"`
	StyleFile  string            `json:"style_file,omitempty"`
	ExtScripts map[string]string `json:"ext_scripts,omitempty"`
	BootData   string            `json:"boot_data,omitempty"`
}
