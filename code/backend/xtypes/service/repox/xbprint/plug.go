package xbprint

import "encoding/json"

type PlugSchemaV1 struct {
	NewPlug
}

type NewPlug struct {
	Slug   string     `json:"slug,omitempty"`
	Name   string     `json:"name,omitempty"`
	Agents []NewAgent `json:"agents,omitempty"`
}

type NewAgent struct {
	Name       string            `json:"name,omitempty"`
	Type       string            `json:"type,omitempty"`
	Executor   string            `json:"executor,omitempty"`
	IfaceFile  string            `json:"iface_file,omitempty"`
	EntryFile  string            `json:"entry_file,omitempty"`
	WebOptions map[string]string `json:"web_options,omitempty"`
	WebFiles   map[string]string `json:"web_files,omitempty"`

	// secondary objects
	Resources []NewAgentResource `json:"resources,omitempty"`
}

type NewAgentResource struct {
	Name    string       `json:"name,omitempty"`
	Type    string       `json:"type,omitempty"`
	RefName string       `json:"ref_name,omitempty"`
	RefData *NewResource `json:"ref_data,omitempty"`
}

type PlugSchemaV2 struct {
	Steps []Step `json:"steps,omitempty"`
}

type NewInnerLink struct {
	Slug       string            `json:"slug,omitempty"`
	From       string            `json:"from,omitempty"`
	To         string            `json:"to,omitempty"`
	HandlerMap map[string]string `json:"handler_map,omitempty"`
}

type Step struct {
	Name string          `json:"name,omitempty" yaml:"name,omitempty"`
	Type string          `json:"type,omitempty" yaml:"type,omitempty"`
	Data json.RawMessage `json:"data,omitempty" yaml:"data,omitempty"`
}
