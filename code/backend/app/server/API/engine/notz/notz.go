package notz

import (
	"net/http/httputil"
	"sync"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz"
)

var _ xnotz.Notz = (*Notz)(nil)

type Notz struct {
	ehub    etypes.EngineHub
	corehub store.CoreHub
	cabinet store.CabinetHub
	ecache  etypes.Ecache

	laddrs map[string]string
	laLock sync.Mutex

	rawProxies map[string]*httputil.ReverseProxy
	rLock      sync.RWMutex
}

func New(app xtypes.App) *Notz {
	deps := app.GetDeps()

	ehub := deps.EngineHub().(etypes.EngineHub)
	corehub := deps.CoreHub().(store.CoreHub)
	cabinet := deps.Cabinet().(store.CabinetHub)

	return &Notz{
		ehub:       ehub,
		corehub:    corehub,
		cabinet:    cabinet,
		ecache:     nil,
		laddrs:     make(map[string]string),
		laLock:     sync.Mutex{},
		rawProxies: make(map[string]*httputil.ReverseProxy),
		rLock:      sync.RWMutex{},
	}
}

func (n *Notz) Start() error {
	ecahe := n.ehub.GetCache()
	if ecahe == nil {
		return easyerr.Error("ecache not found")
	}

	n.ecache = ecahe

	return nil
}

func (n *Notz) RegisterLocalAddr(plug, agent, addr string) {
	n.laLock.Lock()
	n.laddrs[plug+agent] = addr
	n.laLock.Unlock()
}
