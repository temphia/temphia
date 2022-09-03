package pacman

import (
	"sync"

	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/services/pacman/instancers"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints/instancer"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"

	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

var (
	_ service.Pacman = (*PacMan)(nil)
)

type PacMan struct {
	app     xtypes.App
	sockd   sockdx.SockdCore
	syncer  store.SyncDB
	dynHub  store.DynHub
	cabinet store.CabinetHub

	instancers      map[string]instancer.Instancer
	repoBuilders    map[string]repox.Builder
	activeRepo      map[string]map[int64]repox.Repository
	activeRepoMutex sync.Mutex
}

func New(_app xtypes.App) *PacMan {

	deps := _app.GetDeps()

	return &PacMan{
		app:             _app,
		sockd:           deps.Sockd().(sockdx.SockdCore),
		syncer:          deps.CoreHub().(service.Syncer),
		dynHub:          deps.DynHub().(store.DynHub),
		cabinet:         deps.Cabinet().(store.CabinetHub),
		instancers:      nil,
		repoBuilders:    nil,
		activeRepoMutex: sync.Mutex{},
		activeRepo:      make(map[string]map[int64]repox.Repository),
	}
}

func (p *PacMan) blobStore(tenantId string) store.CabinetSourced {
	return p.cabinet.Default(tenantId)
}

func (p *PacMan) Start() error {
	p.instancers = instancers.GetInstancers(p.app)

	reg := p.app.GetDeps().Registry().(*registry.Registry)
	p.repoBuilders = reg.GetRepoBuilders()

	return nil
}
