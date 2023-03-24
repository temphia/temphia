package pagedash

type Dash struct {
	Name     string             `json:"name,omitempty" yaml:"name,omitempty"`
	Sources  map[string]*Source `json:"sources,omitempty" yaml:"sources,omitempty"`
	Sections []Section          `json:"sections,omitempty" yaml:"sections,omitempty"`
	OnLoad   string             `json:"on_load,omitempty" yaml:"on_load,omitempty"`
	OnBuild  string             `json:"on_build,omitempty" yaml:"on_build,omitempty"`
}

type Section struct {
	Name   string   `json:"name,omitempty" yaml:"name,omitempty"`
	Layout string   `json:"layout,omitempty" yaml:"layout,omitempty"`
	Panels []*Panel `json:"panels,omitempty" yaml:"panels,omitempty"`
}

type Panel struct {
	Name     string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Width    uint8                  `json:"width,omitempty" yaml:"width,omitempty"`
	Height   uint8                  `json:"height,omitempty" yaml:"height,omitempty"`
	Interval uint8                  `json:"interval,omitempty" yaml:"interval,omitempty"`
	Type     string                 `json:"type,omitempty" yaml:"type,omitempty"`
	Source   string                 `json:"source,omitempty" yaml:"source,omitempty"`
	Options  map[string]interface{} `json:"options,omitempty" yaml:"options,omitempty"`
}

type Source struct {
	Type    string         `json:"type,omitempty" yaml:"type,omitempty"`
	Data    any            `json:"data,omitempty" yaml:"data,omitempty"`
	Options map[string]any `json:"options,omitempty" yaml:"options,omitempty"`
	Handler string         `json:"handler,omitempty" yaml:"handler,omitempty"`
}
