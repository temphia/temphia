package xnotz

type Notz interface {
	HandleAgent(ctx Context)
	HandleDomain(ctx Context)
	RegisterLocalAddr(plug, agent, addr string)
	Start() error
}
