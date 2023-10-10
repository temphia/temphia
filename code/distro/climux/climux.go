package climux

import "sync"

type CLiFunc func(args []string) error

type CliAction struct {
	Name string
	Help string
	Func func(args []string) error
}

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
