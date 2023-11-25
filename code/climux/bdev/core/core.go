package core

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
	"github.com/temphia/temphia/code/goclient/devc"
)

type BdevContext struct {
	ConfigFile string
	KongCtx    *kong.Context
	Vars       EnvVars
}

func (b *BdevContext) MustGetConfig() *xpackage.Manifest {

	conf, err := b.readConfig()
	if err != nil {
		panic(err)
	}

	return conf
}

func (b *BdevContext) readConfig() (*xpackage.Manifest, error) {

	if b.ConfigFile == "" {
		b.ConfigFile = os.Getenv(xtypes.EnvBdevBprintConfig)
		if b.ConfigFile == "" {
			for _, pfile := range []string{".bprint.toml", ".meta/.bprint.toml"} {
				if xutils.FileExists("./", pfile) {
					b.ConfigFile = pfile
				}
			}

			if b.ConfigFile == "" {
				return nil, easyerr.NotFound("bprint.toml")
			}
		}

	}

	out, err := os.ReadFile(b.ConfigFile)
	if err != nil {
		return nil, easyerr.Wrap("bprint file not found", err)
	}

	bprint := &xpackage.Manifest{}
	err = toml.Unmarshal(out, bprint)
	if err != nil {
		return nil, easyerr.Wrap("err unmarsheling .bprint file", err)
	}

	godotenv.Load(bprint.EnvFile)

	cctx := ReadEnvVars()
	b.Vars = *cctx

	return bprint, nil
}

func (b *BdevContext) GetDevClient() (*devc.DevClient, error) {

	_, err := b.readConfig()
	if err != nil {
		return nil, err
	}

	return devc.New(b.Vars.APIURL, b.Vars.Token), nil

}

type EnvVars struct {
	Token   string `json:"token,omitempty"`
	APIURL  string `json:"api_url,omitempty"`
	PlugId  string `json:"plug_id,omitempty"`
	AgentId string `json:"agent_id,omitempty"`
}

func ReadEnvVars() *EnvVars {

	return &EnvVars{
		Token:   os.Getenv(xtypes.EnvBdevToken),
		APIURL:  os.Getenv(xtypes.EnvBdevApiURL),
		PlugId:  os.Getenv(xtypes.EnvBdevPlugId),
		AgentId: os.Getenv(xtypes.EnvBdevAgentId),
	}
}
