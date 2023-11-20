package core

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
)

type BdevContext struct {
	ConfigFile string
	KongCtx    *kong.Context
}

func (b *BdevContext) MustGetConfig() *xpackage.Manifest {

	return nil
}

type EnvVars struct {
	Token   string `json:"token,omitempty"`
	APIURL  string `json:"api_url,omitempty"`
	PlugId  string `json:"plug_id,omitempty"`
	AgentId string `json:"agent_id,omitempty"`
}

func ReadEnvVars() *EnvVars {

	return &EnvVars{
		Token:   os.Getenv("TEMPHIA_BDEV_TOKEN"),
		APIURL:  os.Getenv("TEMPHIA_BDEV_API_URL"),
		PlugId:  os.Getenv("TEMPHIA_BDEV_PLUG_ID"),
		AgentId: os.Getenv("TEMPHIA_BDEV_AGENT_ID"),
	}
}
