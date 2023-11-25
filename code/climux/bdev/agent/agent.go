package agent

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/climux/bdev/core"
)

type AgentCLI struct {
	List struct{} `cmd:"" help:"List agents."`

	Reset struct {
		AgentId string `arg:"" help:"Agent id to reset."`
	} `cmd:"" help:"Reset running agent."`
	Watch struct {
		AgentId string `arg:"" help:"Agent id to watch."`
	} `cmd:"" help:"Watch agent events."`

	Status struct {
		AgentId string `arg:"" help:"Agent id to check status."`
	} `cmd:"" help:"Get agent status."`

	RPXExec struct{} `cmd:"" help:"Execute action on agent."`
	WebExec struct{} `cmd:"" help:"Execute action on agent."`

	PlugId string
}

func (a *AgentCLI) Run(ctx core.BdevContext) error {
	pp.Println("@agent_list", ctx.KongCtx.Command())

	switch ctx.KongCtx.Command() {
	case "agent list":
		return a.list(ctx)
	case "agent reset <agent-id>":
		return a.reset(ctx)
	case "agent watch <agent-id>":
		return a.watch(ctx)
	case "agent status <agent-id>":
		return a.status(ctx)
	case "agent rpx-exec":
		return a.rpxExec(ctx)
	case "agent web-exec":
		return a.webExec(ctx)
	default:
		panic("not implemented" + ctx.KongCtx.Command())

	}

}

func (a *AgentCLI) list(ctx core.BdevContext) error {

	dc, err := ctx.GetDevClient()
	if err != nil {
		return err
	}

	pp.Println("fixme", dc)

	return nil
}

func (a *AgentCLI) reset(ctx core.BdevContext) error {
	dc, err := ctx.GetDevClient()
	if err != nil {
		return err
	}

	if a.PlugId == "" {
		a.PlugId = ctx.Vars.PlugId
	}

	if a.Reset.AgentId == "" {
		a.Reset.AgentId = ctx.Vars.AgentId
	}

	return dc.Reset(a.PlugId, a.Reset.AgentId)

}

func (a *AgentCLI) watch(ctx core.BdevContext) error {
	dc, err := ctx.GetDevClient()
	if err != nil {
		return err
	}

	if a.PlugId == "" {
		a.PlugId = ctx.Vars.PlugId
	}

	if a.Watch.AgentId == "" {
		a.Watch.AgentId = ctx.Vars.AgentId
	}

	dc.Watch(a.PlugId, a.Watch.AgentId)

	return nil
}

func (a *AgentCLI) status(ctx core.BdevContext) error {

	dc, err := ctx.GetDevClient()
	if err != nil {
		return err
	}

	if a.PlugId == "" {
		a.PlugId = ctx.Vars.PlugId
	}

	if a.Status.AgentId == "" {
		a.Status.AgentId = ctx.Vars.AgentId
	}

	pp.Println("fixme", dc)

	return nil
}

func (a *AgentCLI) rpxExec(ctx core.BdevContext) error {
	return nil
}

func (a *AgentCLI) webExec(ctx core.BdevContext) error {
	return nil
}
