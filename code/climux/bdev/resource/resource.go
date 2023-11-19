package resource

type ResourceCLI struct {
	List      struct{} `cmd:"" help:"list resources inside this bprint/app."`
	LinkAgent struct{} `cmd:"" help:"link resource to agent."`
	Edit      struct{} `cmd:"" help:"edit resource."`
	Delete    struct{} `cmd:"" help:"delete resource."`
}
