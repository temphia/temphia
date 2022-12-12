package dashmodels

type Dashboard struct {
	Name     string             `json:"name,omitempty" yaml:"name,omitempty"`
	Sources  map[string]*Source `json:"sources,omitempty" yaml:"sources,omitempty"`
	Sections []*Section         `json:"sections,omitempty" yaml:"sections,omitempty"`
}

type Source struct {
	Type       string                 `json:"type,omitempty" yaml:"type,omitempty"`
	Options    map[string]interface{} `json:"options,omitempty" yaml:"options,omitempty"`
	Depends    []string               `json:"depends,omitempty" yaml:"depends,omitempty"`
	StaticData interface{}            `json:"static_data,omitempty" yaml:"static_data,omitempty"`
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
	Interval string                 `json:"interval,omitempty" yaml:"interval,omitempty"`
	Type     string                 `json:"type,omitempty" yaml:"type,omitempty"`
	Source   string                 `json:"source,omitempty" yaml:"source,omitempty"`
	Options  map[string]interface{} `json:"options,omitempty" yaml:"options,omitempty"`
}
