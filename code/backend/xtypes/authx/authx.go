package authx

type HandleOpts struct {
	Stage string
	Ctx   Context
	Opts  any
}

type FinishOpts struct {
	Ctx  Context
	Opts any
}

type AuthdProvider interface {
	Handle(opts HandleOpts) (any, error)
	Finish(opts FinishOpts) (any, error)
	GetJS(string) ([]byte, error)
}
