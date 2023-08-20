package pacman

import (
	"sync"

	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xinstancer"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"

	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var (
	_ xpacman.Pacman = (*PacMan)(nil)
)

type PacMan struct {
	app     xtypes.App
	corehub store.CoreHub
	dynHub  dyndb.DataHub
	cabinet store.CabinetHub

	repoBuilders    map[string]xpacman.Builder
	activeRepo      map[string]map[int64]xpacman.Repository
	activeRepoMutex sync.Mutex
	bstore          xpacman.BStore
}

func New(_app xtypes.App) *PacMan {

	deps := _app.GetDeps()
	fhub := deps.Cabinet().(store.CabinetHub)

	pm := &PacMan{
		app:             _app,
		corehub:         deps.CoreHub().(store.CoreHub),
		dynHub:          deps.DataHub().(dyndb.DataHub),
		cabinet:         fhub,
		repoBuilders:    nil,
		activeRepoMutex: sync.Mutex{},
		activeRepo:      make(map[string]map[int64]xpacman.Repository),
	}

	return pm
}

func (p *PacMan) Start() error {

	reg := p.app.GetDeps().Registry().(*registry.Registry)
	p.repoBuilders = reg.GetRepoBuilders()

	return nil
}

func (p *PacMan) GetBprintFileStore() xpacman.BStore { return nil }
func (p *PacMan) GetInstancer() xinstancer.Instancer { return nil }
