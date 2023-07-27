package agent

import "net/http"

// /z/agent_auth

type AgentServer interface {
	Handle(w http.ResponseWriter, req *http.Request)
}

type Agent struct {
	Id                string         `json:"id,omitempty" db:"id,omitempty"`
	Name              string         `json:"name,omitempty" db:"name,omitempty"`
	Type              string         `json:"type,omitempty" db:"type,omitempty"` // headless, spa, ssr, template
	Executor          string         `json:"executor,omitempty" db:"executor,omitempty"`
	IfaceFile         string         `json:"iface_file,omitempty" db:"iface_file,omitempty"`
	EntryFile         string         `json:"entry_file,omitempty" db:"entry_file,omitempty"`
	WebFiles          map[string]any `json:"web_files,omitempty" db:"web_files,omitempty"`
	WebOptions        map[string]any `json:"web_options,omitempty" db:"web_options,omitempty"`
	ServiceWorkerFile string         `json:"sw_file,omitempty" db:"sw_file,omitempty"`
	ExtraMeta         map[string]any `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
	ModVersion        int64          `json:"mod_version,omitempty" db:"mod_version,omitempty"`
	PlugId            string         `json:"plug_id,omitempty" db:"plug_id,omitempty"`
	TenantId          string         `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}

type SPAOptions struct {
	WebEntry  string `json:"web_entry,omitempty"`
	WebScript string `json:"web_script,omitempty"`
	WebStyle  string `json:"web_style,omitempty"`
	WebLoader string `json:"web_loader,omitempty"`
}

type SSROptions struct {
	AutoInject     string `json:"auto_inject,omitempty"`
	InjectKey      string `json:"inject_key,omitempty"`
	InjectFiles    string `json:"inject_files,omitempty"`
	ActionRedirect string `json:"action_redirect,omitempty"`
}
