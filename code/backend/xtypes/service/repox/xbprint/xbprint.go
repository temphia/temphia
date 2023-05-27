package xbprint

const (
	TypeBundle     = "bundle"
	TypeDataGroup  = "data_group"
	TypeDataSheet  = "data_sheet"
	TypePlug       = "plug"
	TypeResource   = "resource"
	TypeTargetApp  = "target_app"
	TypeTargetHook = "target_hook"
)

type LocalBprint struct {
	Name        string            `yaml:"name,omitempty"`
	Slug        string            `yaml:"slug,omitempty"`
	Type        string            `yaml:"type,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Icon        string            `yaml:"icon,omitempty"`
	Version     string            `yaml:"version,omitempty"`
	Tags        []string          `yaml:"tags,omitempty"`
	Files       map[string]string `yaml:"files,omitempty"`
	ExtraMeta   map[string]any    `yaml:"extra_meta,omitempty"`
	EnvFile     string            `yaml:"env_file,omitempty"`
}
