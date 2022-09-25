package config

type Simple struct {
	AppName          string         `json:"app_name,omitempty"`
	MasterKey        string         `json:"master_key,omitempty"`
	ServerPort       string         `json:"http_port,omitempty"`
	OperatorName     string         `json:"op_name,omitempty"`
	OperatorPassword string         `json:"op_password,omitempty"`
	Database         *StoreSource   `json:"database,omitempty"`
	ExecutorOptions  map[string]any `json:"executors,omitempty"`
	ModulesOptions   map[string]any `json:"modules,omitempty"`
	FilesFolder      string         `json:"files_folder,omitempty"`
}

func (s *Simple) AsConfig() *Config {

	conf := &Config{
		AppName:          s.AppName,
		MasterKey:        s.MasterKey,
		OperatorName:     s.OperatorName,
		OperatorPassword: s.OperatorPassword,
		Stores: map[string]*StoreSource{
			"default": s.Database,
		},
		NodeOptions: &NodeOptions{
			ServerPort: s.ServerPort,
			Tags:       []string{},
		},
		ExecutorOptions: s.ExecutorOptions,
		ModulesOptions:  s.ModulesOptions,
		Coredb:          "default",
		DefaultCabinet:  "default",
		DefaultDyndb:    "default",
	}

	if s.FilesFolder != "" {
		conf.Stores["default_fs"] = &StoreSource{
			Name:     "default_fs",
			Provider: "local_fs",
			HostPath: s.FilesFolder,
		}

		conf.DefaultCabinet = "default_fs"
	}

	return conf
}
