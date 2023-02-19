package bdev

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
	client "github.com/temphia/temphia/code/goclient"
	"github.com/temphia/temphia/code/goclient/devc"
	"github.com/tidwall/pretty"
)

type CLI struct {
	Push struct {
		Name string `arg:"" help:"File name to push"`
		File string `arg:"" help:"File to push"`
	} `cmd:"" help:"Push File in bprint."`

	Execute struct {
		Action string `arg:"" help:"File name to push"`
		Data   string `arg:"" help:"File name to push"`
	} `cmd:"" help:"Execute action on agent."`

	Watch struct {
	} `cmd:"" help:"Watch agent."`

	Reset struct {
	} `cmd:"" help:"Reset agent."`

	ctx       *kong.Context
	devClient *devc.DevClient
	PlugId    string
	AgentId   string
}

type UpperScope struct {
	BprintFile string
	Ctx        *kong.Context
}

func (c *CLI) preRun(bfile string) error {

	_, err := godotenv.Read(bfile)
	if err != nil {
		return err
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

func (c *CLI) Run(uscope *UpperScope) error {

	if uscope.BprintFile == "" {
		panic(".bprint.yaml not specified")
	}

	c.ctx = uscope.Ctx
	c.preRun(uscope.BprintFile)

	switch c.ctx.Command() {
	case "bdev push <name> <file>":
		c.push()
	case "bdev execute <action> <data>":
		c.execute()
	case "bdev  reset":
		c.reset()
	case "bdev watch":
		c.watch()
	default:
		panic("Command not found |> " + c.ctx.Command())
	}

	return nil

}

func (c *CLI) reset() {
	c.devClient.Reset(c.PlugId, c.AgentId)
}

func (c *CLI) watch() {
	c.devClient.Watch(c.PlugId, c.AgentId)
}

func (c *CLI) push() {
	c.devClient.PushFile(c.Push.Name, c.Push.File)
}

func (c *CLI) execute() {
	var data any
	err := json.Unmarshal([]byte(c.Execute.Data), &data)
	if err != nil {
		pp.Println(err)
		return
	}

	resp, err := c.devClient.ExecRun(c.PlugId, c.AgentId, c.Execute.Action, data)
	if err != nil {
		pp.Println(err)
		return
	}

	fmt.Print(string(pretty.Color(pretty.Pretty(resp), nil)))

}
