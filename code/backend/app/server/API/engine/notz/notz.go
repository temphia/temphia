package notz

import (
	"sync"

	"github.com/temphia/temphia/code/backend/app/server/API/engine/router"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz"
)

var _ xnotz.Notz = (*Notz)(nil)

type routeCache struct {
	bprintId string
	config   *router.RouteConfig
}

type Notz struct {
	ehub    etypes.EngineHub
	corehub store.CoreHub
	cabinet store.CabinetHub
	ecache  etypes.Ecache
	pacman  xpacman.Pacman
	engine  etypes.Engine

	routesCaches map[string]*routeCache
	rMutext      sync.Mutex
}

func New(app xtypes.App) *Notz {
	deps := app.GetDeps()

	ehub := deps.EngineHub().(etypes.EngineHub)
	corehub := deps.CoreHub().(store.CoreHub)
	cabinet := deps.Cabinet().(store.CabinetHub)
	pacman := deps.RepoHub().(xpacman.Pacman)

	return &Notz{
		ehub:         ehub,
		corehub:      corehub,
		cabinet:      cabinet,
		ecache:       nil,
		pacman:       pacman,
		routesCaches: make(map[string]*routeCache),
		rMutext:      sync.Mutex{},
	}
}

func (n *Notz) Start() error {
	ecahe := n.ehub.GetCache()
	if ecahe == nil {
		return easyerr.Error("ecache not found")
	}

	n.engine = n.ehub.GetEngine()
	n.ecache = ecahe

	return nil
}
