package distro

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
)

const (
	TemphiaStateFolder = ".temphia-data"
	TemphiaConfigFile  = "temphia.json"
)

func getConfig() []byte {
	randomId, err := xutils.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}

	return []byte(fmt.Sprintf(`{
		"server_port": ":4000",
		"tenant_id": "default0",
		"root_domain": "localhost",
		"runner_domain": "localhost",
		"master_key": "%s",
		"enable_local_door": true,
		"data_folder": "./%s",
		"database_config": {
			"name": "sqlite",
			"vendor": "sqlite",
			"provider": "sqlite",
			"target": "main.db"
		}
	}`, randomId, TemphiaStateFolder))

}

func (a *AppCLi) readConfig() (*config.Config, error) {

	if a.ConfigFile == "" {

		if a.ctx.Command() == "init-data" {
			os.Mkdir(TemphiaStateFolder, os.FileMode(0777))
			os.WriteFile(TemphiaConfigFile, getConfig(), os.FileMode(0666))
			a.ConfigFile = TemphiaConfigFile

		} else {
			if xutils.FileExists("./", TemphiaConfigFile) {
				a.ConfigFile = TemphiaConfigFile
			}
		}

	}

	return readConfig(a.ConfigFile)
}

func readConfig(file string) (*config.Config, error) {

	if file == "" {
		return nil, easyerr.Error("--config-file not passed")
	}

	out, err := os.ReadFile(file)
	if err != nil {
		return nil, easyerr.Wrap("err reading config file", err)
	}

	conf := &config.Config{}
	err = json.Unmarshal(out, &conf)
	if err != nil {
		return nil, easyerr.Wrap("err parsing config JSON", err)
	}

	err = conf.Init()
	if err != nil {
		return nil, err
	}

	return conf, nil
}
