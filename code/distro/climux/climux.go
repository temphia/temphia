package climux

import "sync"

var (
	DefaultCLI = "help"
)

type CLiFunc func(args []string) error

type CliAction struct {
	Name string
	Help string
	Func func(ctx Context) error
}

type Registery map[string]*CliAction

var (
	cliRegistery = map[string]*CliAction{}
	cLock        sync.Mutex
)

func Register(a *CliAction) {
	cLock.Lock()
	cliRegistery[a.Name] = a
	cLock.Unlock()
}

func GetRegistry() map[string]*CliAction {

	cLock.Lock()
	resp := map[string]*CliAction{}
	cLock.Unlock()

	for k, v := range cliRegistery {
		resp[k] = v
	}

	return cliRegistery
}
