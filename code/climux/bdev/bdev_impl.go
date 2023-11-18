package bdev

import (
	"encoding/json"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/tidwall/pretty"
)

func (c *CLI) reset() error {
	return c.devClient.Reset(c.PlugId, c.AgentId)
}

func (c *CLI) watch() {
	c.devClient.Watch(c.PlugId, c.AgentId)
}

func (c *CLI) zipit() error {
	if c.Zip.OutFile == "" {
		c.Zip.OutFile = fmt.Sprintf("build/%s.zip", c.bp.Slug)
	}

	return (ZipIt(c.bp, c.Zip.OutFile))
}

func (c *CLI) push() error {
	return (c.devClient.PushFile(c.Push.Name, c.bp.Files[c.Push.Name]))
}

func (c *CLI) execute() error {
	var data any
	err := json.Unmarshal([]byte(c.Exec.Data), &data)
	if err != nil {
		pp.Println(err)
		return err
	}

	resp, err := c.devClient.ExecRun(c.PlugId, c.AgentId, c.Exec.Action, data)
	if err != nil {
		pp.Println(err)
		return err
	}

	fmt.Print(string(pretty.Color(pretty.Pretty(resp), nil)))

	return nil
}
