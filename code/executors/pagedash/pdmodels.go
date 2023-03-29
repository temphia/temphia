package pagedash

type DashModel struct {
	Name       string            `json:"name,omitempty" yaml:"name,omitempty"`
	Sections   []Section         `json:"sections,omitempty" yaml:"sections,omitempty"`
	OnLoad     []Hook            `json:"on_load,omitempty" yaml:"on_load,omitempty"`
	OnBuild    []Hook            `json:"on_build,omitempty" yaml:"on_build,omitempty"`
	StaticData map[string]any    `json:"static_data,omitempty" yaml:"static_data,omitempty"`
	Sources    map[string]Source `json:"sources,omitempty" yaml:"sources,omitempty"`
}

type Hook struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Target  string         `json:"target,omitempty" yaml:"target,omitempty"`
	Type    string         `json:"type,omitempty" yaml:"type,omitempty"`
	Options map[string]any `json:"options,omitempty" yaml:"options,omitempty"`
}

type Section struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Layout  string         `json:"layout,omitempty" yaml:"layout,omitempty"`
	Panels  []Panel        `json:"panels,omitempty" yaml:"panels,omitempty"`
	Options map[string]any `json:"options,omitempty" yaml:"options,omitempty"`
}

type Panel struct {
	Name     string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type     string         `json:"type,omitempty" yaml:"type,omitempty"`
	Info     string         `json:"info,omitempty" yaml:"info,omitempty"`
	ViewOpts map[string]any `json:"view_opts,omitempty" yaml:"view_opts,omitempty"`
	DataOpts map[string]any `json:"data_opts,omitempty" yaml:"data_opts,omitempty"`
	Source   string         `json:"source,omitempty" yaml:"source,omitempty"`
}

type Source struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    string         `json:"type,omitempty" yaml:"type,omitempty"`
	Options map[string]any `json:"options,omitempty" yaml:"options,omitempty"`
}

// req/resp types

type LoadRequest struct {
	ExecData any    `json:"exec_data,omitempty"`
	Version  string `json:"version,omitempty"`
}

type LoadResponse struct {
	Name     string            `json:"name,omitempty" yaml:"name,omitempty"`
	Data     map[string]any    `json:"data,omitempty" yaml:"data,omitempty"`
	Sources  map[string]Source `json:"sources,omitempty" yaml:"sources,omitempty"`
	Sections []Section         `json:"sections,omitempty" yaml:"sections,omitempty"`
}

type BuildRequest struct {
	Options any `json:"options,omitempty"`
}

type BuildRespone struct {
	Id string `json:"options,omitempty"`
}
