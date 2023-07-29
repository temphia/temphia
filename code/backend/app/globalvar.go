package app

import "sync"

type Global struct {
	globalVars map[string]any
	gmutex     sync.Mutex
}

func (g *Global) Set(key string, val any) {
	g.gmutex.Lock()
	defer g.gmutex.Unlock()
	g.globalVars[key] = val
}

func (g *Global) Get(key string) any {
	g.gmutex.Lock()
	defer g.gmutex.Unlock()

	return g.globalVars[key]
}

func (g *Global) Del(key string) {
	g.gmutex.Lock()
	defer g.gmutex.Unlock()
	delete(g.globalVars, key)
}
