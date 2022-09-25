package dev

import "github.com/temphia/temphia/code/core/backend/app/config"

var conf = config.Simple{
	AppName:          "Dev Test",
	MasterKey:        "test123",
	ServerPort:       ":4000",
	OperatorName:     "ops",
	OperatorPassword: "ops123",
	Database: &config.StoreSource{
		Name:     "default_db",
		Provider: "postgres",
		Vendor:   "postgres",
		HostPath: "localhost",
		Target:   "temphia",
		User:     "temphia",
		Password: "temphia123",
		Port:     "7032",
		Features: []string{"core_db", "state_db", "dyn_db"},
		Options:  map[string]interface{}{},
	},

	ExecutorOptions: make(map[string]any),
	ModulesOptions:  make(map[string]any),
	FilesFolder:     "./tmp/files",
}
