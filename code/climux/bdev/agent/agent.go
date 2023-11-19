package agent

type BDevCLI struct {
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
