package pacman

import (
	"sync"

	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/services/pacman/instancerhub/hubv1"
	"github.com/temphia/temphia/code/backend/services/pacman/instancerhub/instancers"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xinstance"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"

	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var (
	_ repox.Pacman = (*PacMan)(nil)
)

type PacMan struct {
	app     xtypes.App
	corehub store.CoreHub
	dynHub  dyndb.DataHub
	cabinet store.CabinetHub

	instancers      map[string]xinstance.Instancer
	repoBuilders    map[string]repox.Builder
	activeRepo      map[string]map[int64]repox.Repository
	activeRepoMutex sync.Mutex

	ihubV1 repox.InstancerHubV1
	ihubV2 repox.InstancerHubV2
}

func New(_app xtypes.App) *PacMan {

	deps := _app.GetDeps()

	pm := &PacMan{
		app:             _app,
		corehub:         deps.CoreHub().(store.CoreHub),
		dynHub:          deps.DataHub().(dyndb.DataHub),
		cabinet:         deps.Cabinet().(store.CabinetHub),
		instancers:      nil,
		repoBuilders:    nil,
		activeRepoMutex: sync.Mutex{},
		activeRepo:      make(map[string]map[int64]repox.Repository),
	}

	return pm
}

func (p *PacMan) Start() error {

	reg := p.app.GetDeps().Registry().(*registry.Registry)
	p.repoBuilders = reg.GetRepoBuilders()

	p.ihubV1 = hubv1.New(instancers.GetInstancers(p.app), p)

	return nil
}

func (p *PacMan) GetInstancerHubV1() repox.InstancerHubV1 { return p.ihubV1 }

func (p *PacMan) GetInstancerHubV2() repox.InstancerHubV2 { return p.ihubV2 }
