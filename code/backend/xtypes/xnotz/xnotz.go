package xnotz

type Notz interface {
	HandleAgent(ctx Context)
	HandleDomain(ctx Context)
}
