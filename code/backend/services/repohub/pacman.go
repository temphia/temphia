package repohub

import (
	"sync"

	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/services/repohub/instancers"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox/xinstance"

	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

var (
	_ repox.Hub = (*PacMan)(nil)
)

type PacMan struct {
	app     xtypes.App
	corehub store.CoreHub
	dynHub  store.DataHub
	cabinet store.CabinetHub

	instancers      map[string]xinstance.Instancer
	repoBuilders    map[string]repox.Builder
	activeRepo      map[string]map[int64]repox.Repository
	activeRepoMutex sync.Mutex
}

func New(_app xtypes.App) *PacMan {

	deps := _app.GetDeps()

	return &PacMan{
		app:             _app,
		corehub:         deps.CoreHub().(store.CoreHub),
		dynHub:          deps.DataHub().(store.DataHub),
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
