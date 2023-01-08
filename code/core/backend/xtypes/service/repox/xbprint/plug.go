package xbprint

type NewPlug struct {
	Slug   string     `json:"slug,omitempty"`
	Name   string     `json:"name,omitempty"`
	Agents []NewAgent `json:"agents,omitempty"`
}

type NewAgent struct {
	Name      string            `json:"name,omitempty"`
	Type      string            `json:"type,omitempty"`
	Executor  string            `json:"executor,omitempty"`
	IfaceFile string            `json:"iface_file,omitempty"`
	WebEntry  string            `json:"web_entry,omitempty"`
	WebScript string            `json:"web_script,omitempty"`
	WebStyle  string            `json:"web_style,omitempty"`
	WebLoader string            `json:"web_loader,omitempty"`
	WebFiles  map[string]string `json:"web_files,omitempty"`

	// secondary objects
	Resources  map[string]string `json:"resources,omitempty"`
	Links      map[string]string `json:"links,omitempty"`
	Extensions map[string]string `json:"exts,omitempty"`
}
