package demo

import "github.com/temphia/temphia/code/backend/app/config/simple"

var sqliteConf = simple.Config{
	AppName:          "Demo App",
	MasterKey:        "__secret_demo__11",
	ServerPort:       ":5000",
	OperatorName:     "ops",
	OperatorPassword: "ops123",
	Database: simple.DatabaseOptions{
		Provider: "sqlite",
		Vendor:   "sqlite",
		HostPath: "temphia-data/sqlitedata/demo.db",
		Options:  map[string]interface{}{},
	},

	ExecutorOptions: make(map[string]any),
	ModulesOptions:  make(map[string]any),
	FilesFolder:     "./temphia-data/files",
	LogFolder:       "./temphia-data/logs",
}

var postgresConf = simple.Config{
	AppName:          "Demo App",
	MasterKey:        "__secret_demo__11",
	ServerPort:       ":5000",
	OperatorName:     "ops",
	OperatorPassword: "ops123",
	Database: simple.DatabaseOptions{
		Provider: "postgres",
		Vendor:   "postgres",
		HostPath: "localhost",
		Target:   "demo",
		User:     "demo",
		Password: "demo123",
		Port:     "7032",
		Options:  map[string]interface{}{},
	},

	ExecutorOptions: make(map[string]any),
	ModulesOptions:  make(map[string]any),
	FilesFolder:     "./temphia-data/files",
	LogFolder:       "./temphia-data/logs",
}

var Conf simple.Config
