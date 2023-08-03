package config

import (
	"path"

	"github.com/temphia/temphia/code/backend/libx/xutils"
)

type Confd interface {
	GetConfig() *Config
	RootDataFolder() string
	LogFolder() string
	FileStoreFolder() string
	DBFolder() string
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
	return path.Join(c.config.DataFolder, "db", c.config.DatabaseConfig.Vendor)
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

	return xutils.CreateIfNotExits(c.DBFolder())
}
