package bdev

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/tools/sharedcli"

	client "github.com/temphia/temphia/code/goclient"
	"github.com/temphia/temphia/code/goclient/devc"
	"gopkg.in/yaml.v2"
)

type CLI struct {
	Push struct {
		Name string `arg:"" help:"File name to push"`
	} `cmd:"" help:"Push File in bprint."`

	Exec struct {
		Action string `arg:"" help:"Action namme to run."`
		Data   string `arg:"" help:"Data to pass to action."`
	} `cmd:"" help:"Execute action on agent."`

	Watch struct {
	} `cmd:"" help:"Watch agent."`

	Reset struct {
	} `cmd:"" help:"Reset agent."`

	Zip struct {
		OutFile string
	} `cmd:"" help:"create zip from bprint.yaml"`

	ctx       *kong.Context
	devClient *devc.DevClient
	PlugId    string
	AgentId   string

	bp *xbprint.LocalBprint
}

func (c *CLI) preRun(bfile string) error {

	out, err := os.ReadFile(bfile)
	if err != nil {
		return easyerr.Wrap("bprint file not found", err)
	}

	bprint := xbprint.LocalBprint{}
	err = yaml.Unmarshal(out, &bprint)
	if err != nil {
		return easyerr.Wrap("err unmarsheling .bprint file", err)
	}

	c.bp = &bprint

	err = godotenv.Load(bprint.EnvFile)
	if err != nil {
		return easyerr.Wrap("env file load err", err)
	}

	cctx := client.ReadContext()

	c.devClient = devc.New(cctx.APIURL, cctx.Token)
	if c.AgentId == "" {
		c.AgentId = cctx.AgentId
	}

	if c.PlugId == "" {
		c.PlugId = cctx.PlugId
	}

	return nil
}

func (c *CLI) Run(scope *sharedcli.Context) error {

	if scope.BprintFile == "" {
		bconf := os.Getenv("TEMPHIA_BDEV_BPRINT_CONFIG")
		if bconf == "" {
			panic(".bprint.yaml not specified")
		}
		scope.BprintFile = bconf
	}

	c.ctx = scope.KongContext
	err := c.preRun(scope.BprintFile)
	if err != nil {
		return err
	}

	return c.doExecute("bdev ")
}

func (c *CLI) Execute() error {
	return c.doExecute("")
}

func (c *CLI) doExecute(prefix string) error {

	switch c.ctx.Command() {
	case prefix + "push <name>":
		return c.push()
	case prefix + "exec <action> <data>":
		return c.execute()
	case prefix + "reset":
		return c.reset()
	case prefix + "watch":
		c.watch()
	case prefix + "zip":
		return c.zipit()
	default:
		panic("Command not found |> " + c.ctx.Command())
	}

	return nil
}
