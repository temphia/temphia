package config

import "path"

type Config struct {
	ServerPort      string        `json:"server_port,omitempty" toml:"server_port"`
	TenantId        string        `json:"tenant_id,omitempty" toml:"tenant_id"`
	RootDomain      string        `json:"root_domain,omitempty" toml:"root_domain"`
	RunnerDomain    string        `json:"runner_domain,omitempty" toml:"runner_domain"`
	MasterKey       string        `json:"master_key,omitempty" toml:"master_key"`
	EnableLocalDoor bool          `json:"enable_local_door,omitempty" toml:"enable_local_door"`
	DataFolder      string        `json:"data_folder,omitempty" toml:"data_folder"`
	LogFolder       string        `json:"log_folder,omitempty" toml:"log_folder"`
	DatabaseConfig  *StoreConfig  `json:"database_config,omitempty" toml:"database_config"`
	FileStoreConfig *StoreConfig  `json:"filestore_config,omitempty" toml:"filestore_config"`
	LogIngestConfig *IngestConfig `json:"log_ingest_config,omitempty" toml:"log_ingest_config"`
}

type StoreConfig struct {
	Name     string         `json:"name,omitempty" toml:"name"`
	Vendor   string         `json:"vendor,omitempty" toml:"vendor"`
	Provider string         `json:"provider,omitempty" toml:"provider"`
	HostPath string         `json:"host_path,omitempty" toml:"host_path"`
	User     string         `json:"user,omitempty" toml:"user"`
	Password string         `json:"password,omitempty" toml:"password"`
	Port     string         `json:"port,omitempty" toml:"port"`
	Options  map[string]any `json:"options,omitempty" toml:"options"`
	Target   string         `json:"-,omitempty"` // computed
}

func (c *Config) Init() error {

	c.DatabaseConfig.Target = path.Join(c.DataFolder, "db", c.DatabaseConfig.Vendor)

	return nil

}

type IngestConfig struct {
	Upstream string `json:"upstream,omitempty" toml:"upstream"`
	User     string `json:"user,omitempty" toml:"user"`
	Password string `json:"password,omitempty" toml:"password"`
}
