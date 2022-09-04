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
}

func (s *Simple) AsConfig() *Config {

	return &Config{
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
	}
}
