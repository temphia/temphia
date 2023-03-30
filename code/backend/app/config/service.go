package config

import (
	"path"
)

func NewConfigd(basePath string) *Configd {

	return &Configd{
		basePath:     basePath,
		logFolder:    path.Join(basePath, "logs"),
		cplaneFolder: path.Join(basePath, "cplane"),
		dataFolder:   path.Join(basePath, "data"),
	}
}

type Configd struct {
	basePath     string
	logFolder    string
	cplaneFolder string
	dataFolder   string
}

func (c *Configd) Root() string               { return c.basePath }
func (c *Configd) LogFolder() string          { return c.logFolder }
func (c *Configd) ControlPlaneFolder() string { return c.cplaneFolder }
func (c *Configd) DataFolder() string         { return c.cplaneFolder }
