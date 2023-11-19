package socket

type SocketCLI struct {
	List    struct{} `cmd:"" help:"list socket inside this bprint/app."`
	Refresh struct{} `cmd:"" help:"refresh socket."`
	Status  struct{} `cmd:"" help:"get status of socket."`
	Watch   struct{} `cmd:"" help:"watch socket."`
}
