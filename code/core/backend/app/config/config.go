package config

type Config struct {
	AppName          string                  `json:"app_name,omitempty"`
	MasterKey        string                  `json:"master_key,omitempty"`
	OperatorName     string                  `json:"op_name,omitempty"`
	OperatorPassword string                  `json:"op_password,omitempty"`
	Stores           map[string]*StoreSource `json:"stores,omitempty"`
	Coredb           string                  `json:"coredb,omitempty"`
	DefaultCabinet   string                  `json:"default_cabinet,omitempty"`
	DefaultDyndb     string                  `json:"default_dyndb,omitempty"`
	NodeOptions      *NodeOptions            `json:"node,omitempty"`
	ExecutorOptions  map[string]any          `json:"executors,omitempty"`
	ModulesOptions   map[string]any          `json:"modules,omitempty"`
}

type NodeOptions struct {
	TenantId   string   `json:"tenant_id,omitempty"`
	ServerPort string   `json:"http_port,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}

type StoreSource struct {
	Name     string         `json:"name,omitempty"`
	Vendor   string         `json:"vendor,omitempty"`
	Provider string         `json:"provider,omitempty"`
	HostPath string         `json:"host_path,omitempty"`
	User     string         `json:"user,omitempty"`
	Password string         `json:"password,omitempty"`
	Target   string         `json:"target,omitempty"`
	Port     string         `json:"port,omitempty"`
	Features []string       `json:"features,omitempty"`
	Options  map[string]any `json:"options,omitempty"`
}
