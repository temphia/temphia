package config

type Config struct {
	ServerPort      string       `json:"server_port,omitempty"`
	TenantId        string       `json:"tenant_id,omitempty"`
	RootDomain      string       `json:"root_domain,omitempty"`
	RunnerDomain    string       `json:"runner_domain,omitempty"`
	MasterKey       string       `json:"master_key,omitempty"`
	EnableLocalDoor bool         `json:"enable_local_door,omitempty"`
	DataFolder      string       `json:"data_folder,omitempty"`
	DatabaseConfig  *StoreConfig `json:"database_config,omitempty"`
	FileStoreConfig *StoreConfig `json:"filestore_config,omitempty"`
}

type StoreConfig struct {
	Name     string         `json:"name,omitempty"`
	Vendor   string         `json:"vendor,omitempty"`
	Provider string         `json:"provider,omitempty"`
	HostPath string         `json:"host_path,omitempty"`
	User     string         `json:"user,omitempty"`
	Password string         `json:"password,omitempty"`
	Port     string         `json:"port,omitempty"`
	Options  map[string]any `json:"options,omitempty"`
	Target   string         `json:"target,omitempty"`
}
