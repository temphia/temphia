package distro

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
)

const (
	TemphiaStateFolder = ".temphia-data"
	TemphiaConfigFile  = "temphia.json"
)

func GetConfig(port, tenantId, root_domain, runner_domain, master_key, statefolder string) []byte {

	if port == "" {
		l, err := net.Listen("tcp", ":4000")
		if err != nil {
			p, err := xutils.GetFreePort()
			if err != nil {
				panic("cannot alllocate port")
			}

			port = fmt.Sprintf(":%d", p)

		} else {
			l.Close()
			port = ":4000"

		}
	}

	if tenantId == "" {
		tenantId = "default0"
	}

	if root_domain == "" {
		root_domain = "localhost"
	}

	if runner_domain == "" {
		runner_domain = "localhost"
	}

	if master_key == "" {
		randomId, err := xutils.GenerateRandomString(32)
		if err != nil {
			panic(err)
		}
		master_key = randomId
	}

	if statefolder == "" {
		statefolder = TemphiaStateFolder
	}

	return []byte(fmt.Sprintf(`{
		"server_port": "%s",
		"tenant_id": "%s",
		"root_domain": "%s",
		"runner_domain": "%s",
		"master_key": "%s",
		"enable_local_door": true,
		"data_folder": "./%s",
		"database_config": {
			"name": "sqlite",
			"vendor": "sqlite",
			"provider": "sqlite"
		}
	}`, port, tenantId, root_domain, runner_domain, master_key, statefolder))

}

func ReadConfig(file string) (*config.Config, error) {

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
		pp.Println("@UNMARSHEL ERROR", string(out), err.Error())
		return nil, easyerr.Wrap("err parsing config JSON", err)
	}

	err = conf.Init()
	if err != nil {
		return nil, err
	}

	return conf, nil
}
