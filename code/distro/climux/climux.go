package climux

import "sync"

type CLiFunc func(args []string) error

var (
	cliRegistery = map[string]CLiFunc{}
	cLock        sync.Mutex
)

func Register(name string, f CLiFunc) {
	cLock.Lock()
	cliRegistery[name] = f
	cLock.Unlock()
}
