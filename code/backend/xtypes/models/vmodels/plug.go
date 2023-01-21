package vmodels

import "encoding/json"

type PlugRaw struct {
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
	WebScript string            `json:"web_script,omitempty"` // file
	WebStyle  string            `json:"web_style,omitempty"`  // file
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

// used in template rendering

type SubOriginData struct {
	LoaderJS       string
	LoaderOptsJSON string
	ApiBaseURL     string
	Token          string
	Plug           string
	Agent          string
	EntryName      string
	ExecLoader     string
	JSPlugScript   string
	StyleFile      string
	ExtScripts     map[string]string
}

func (s *SubOriginData) BuildJSONOpts() error {
	opts := &ExecInstanceOptions{
		ApiBaseURL:   s.ApiBaseURL,
		Token:        s.Token,
		EntryName:    s.EntryName,
		ExecLoader:   s.ExecLoader,
		JSPlugScript: s.JSPlugScript,
		StyleFile:    s.StyleFile,
		ExtScripts:   s.ExtScripts,
		Plug:         s.Plug,
		Agent:        s.Agent,
	}

	out, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	s.LoaderOptsJSON = string(out)
	return nil
}

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
