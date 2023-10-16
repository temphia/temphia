package config

import (
	"fmt"
	"path"

	"github.com/temphia/temphia/code/backend/libx/xutils"
)

type Confd interface {
	GetConfig() *Config
	RootDataFolder() string
	LogFolder() string
	FileStoreFolder() string
	DBFolder() string
	InitDataFolder() error
	LocalSocket() string
	GetRemoteExecEnvs(plug, agent, bprint, token string) []string
}

type confd struct {
	config *Config
}

func New(conf *Config) *confd {
	return &confd{
		config: conf,
	}
}

func (c *confd) GetConfig() *Config { return c.config }

func (c *confd) RootDataFolder() string {
	return c.config.DataFolder
}

func (c *confd) LogFolder() string {
	logfolder := c.config.LogFolder
	if logfolder == "" {
		logfolder = path.Join(c.config.DataFolder, "logs")
	}

	return logfolder
}

func (c *confd) FileStoreFolder() string {
	return path.Join(c.config.DataFolder, "files")
}

func (c *confd) DBFolder() string {
	return path.Join(c.config.DataFolder, "db")
}

func (c *confd) LocalSocket() string {
	return path.Join(c.config.DataFolder, "./local.sock")
}

func (c *confd) InitDataFolder() error {
	err := xutils.CreateIfNotExits(c.config.DataFolder)
	if err != nil {
		return err
	}

	err = xutils.CreateIfNotExits(c.LogFolder())
	if err != nil {
		return err
	}

	err = xutils.CreateIfNotExits(c.FileStoreFolder())
	if err != nil {
		return err
	}

	err = xutils.CreateIfNotExits(c.DBFolder())
	if err != nil {
		return err
	}

	spath := path.Join(c.DBFolder(), c.config.DatabaseConfig.Vendor)

	return xutils.CreateIfNotExits(spath)
}

func (c *confd) GetRemoteExecEnvs(plug, agent, bprint, token string) []string {

	return []string{
		"TEMPHIA_HOST=localhost",
		fmt.Sprintf("TEMPHIA_PORT=%s", c.config.ServerPort),
		fmt.Sprintf("TEMPHIA_RE_PATH=/z/api/%s/v2/engine/remote", c.config.TenantId),
		fmt.Sprintf("TEMPHIA_RE_TOKEN=%s", token),
		fmt.Sprintf("TEMPHIA_TENANT_ID=%s", c.config.TenantId),
		fmt.Sprintf("TEMPHIA_PLUG_ID=%s", plug),
		fmt.Sprintf("TEMPHIA_AGENT_ID=%s", agent),
		fmt.Sprintf("TEMPHIA_BPRINT_ID=%s", bprint),
	}

}
