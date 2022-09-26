package dev

import (
	"github.com/temphia/temphia/code/core/backend/app/config/simple"
)

var conf = simple.Config{
	AppName:          "Dev Test",
	MasterKey:        "test123",
	ServerPort:       ":4000",
	OperatorName:     "ops",
	OperatorPassword: "ops123",
	Database: simple.DatabaseOptions{
		Provider: "postgres",
		Vendor:   "postgres",
		HostPath: "localhost",
		Target:   "temphia",
		User:     "temphia",
		Password: "temphia123",
		Port:     "7032",
		Options:  map[string]interface{}{},
	},

	ExecutorOptions: make(map[string]any),
	ModulesOptions:  make(map[string]any),
	FilesFolder:     "./tmp/files",
}
