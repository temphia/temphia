package xpackage

const (
	TypeBundle     = "bundle"
	TypeDataGroup  = "data_group"
	TypeDataSheet  = "data_sheet"
	TypePlug       = "plug"
	TypeResource   = "resource"
	TypeTargetApp  = "target_app"
	TypeTargetHook = "target_hook"
)

type Manifest struct {
	Name        string            `yaml:"name,omitempty"`
	Slug        string            `yaml:"slug,omitempty"`
	Type        string            `yaml:"type,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Icon        string            `yaml:"icon,omitempty"`
	Screenshots []string          `yaml:"screenshots,omitempty"`
	Version     string            `yaml:"version,omitempty"`
	Tags        []string          `yaml:"tags,omitempty"`
	Files       map[string]string `yaml:"files,omitempty"`
	ExtraMeta   map[string]any    `yaml:"extra_meta,omitempty"`
	EnvFile     string            `yaml:"env_file,omitempty"`
}

type AppSchema struct {
	Name    string               `yaml:"name,omitempty"`
	Slug    string               `yaml:"slug,omitempty"`
	Objects map[string]AppObject `yaml:"objects,omitempty"`
	Steps   []AppStep            `yaml:"steps,omitempty"`
}

type AppStep struct {
	Name string `yaml:"name,omitempty"`
	Type string `yaml:"type,omitempty"`
	Data any    `yaml:"data,omitempty"`
}

type AppObject struct {
	Name string `yaml:"name,omitempty"`
	Type string `yaml:"type,omitempty"`
}
