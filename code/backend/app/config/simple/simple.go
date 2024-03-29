package simple

import (
	"path"

	"github.com/temphia/temphia/code/backend/app/config"
)

type Config struct {
	AppName          string          `json:"app_name,omitempty"`
	MasterKey        string          `json:"master_key,omitempty"`
	ServerPort       string          `json:"http_port,omitempty"`
	OperatorName     string          `json:"op_name,omitempty"`
	OperatorPassword string          `json:"op_password,omitempty"`
	Database         DatabaseOptions `json:"database,omitempty"`
	ExecutorOptions  map[string]any  `json:"executors,omitempty"`
	ModulesOptions   map[string]any  `json:"modules,omitempty"`
	DataFolder       string          `json:"data_folder,omitempty"`
}

type DatabaseOptions struct {
	Vendor   string         `json:"vendor,omitempty"`
	Provider string         `json:"provider,omitempty"`
	HostPath string         `json:"host_path,omitempty"`
	User     string         `json:"user,omitempty"`
	Password string         `json:"password,omitempty"`
	Target   string         `json:"target,omitempty"`
	Port     string         `json:"port,omitempty"`
	Options  map[string]any `json:"options,omitempty"`
}

func (s *Config) AsConfig() *config.Config {

	conf := &config.Config{
		AppName:          s.AppName,
		MasterKey:        s.MasterKey,
		OperatorName:     s.OperatorName,
		OperatorPassword: s.OperatorPassword,
		Stores: map[string]*config.StoreSource{
			"default": {
				Name:     "default",
				Vendor:   s.Database.Vendor,
				Provider: s.Database.Provider,
				HostPath: s.Database.HostPath,
				User:     s.Database.User,
				Password: s.Database.Password,
				Target:   s.Database.Target,
				Port:     s.Database.Port,
				Features: []string{"core_db", "state_db", "dyn_db"},
				Options:  s.Database.Options,
			},
		},
		NodeOptions: &config.NodeOptions{
			ServerPort:    s.ServerPort,
			Tags:          []string{},
			LogFolder:     path.Join(s.DataFolder, "logs"),
			LogFilePrefix: "temphia_log.log",
			NodeCache:     path.Join(s.DataFolder, "nodecache"),
		},

		ExecutorOptions: s.ExecutorOptions,
		ModulesOptions:  s.ModulesOptions,
		Coredb:          "default",
		DefaultCabinet:  "default_fs",
		DefaultDyndb:    "default",
	}

	conf.Stores["default_fs"] = &config.StoreSource{
		Name:     "default_fs",
		Provider: "local_fs",
		HostPath: path.Join(s.DataFolder, "files"),
	}

	conf.DefaultCabinet = "default_fs"

	return conf
}
