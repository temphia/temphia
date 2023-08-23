package pacman

import (
	"sync"

	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/services/pacman/bstore"
	instancers "github.com/temphia/temphia/code/backend/services/pacman/instancer"
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
	instancer       xinstancer.Instancer
}

func New(_app xtypes.App) *PacMan {

	deps := _app.GetDeps()
	fhub := deps.Cabinet().(store.CabinetHub)
	corehub := deps.CoreHub().(store.CoreHub)
	dynHub := deps.DataHub().(dyndb.DataHub)
	bstore := bstore.New(fhub)

	pm := &PacMan{
		app:             _app,
		corehub:         corehub,
		dynHub:          dynHub,
		cabinet:         fhub,
		repoBuilders:    nil,
		activeRepoMutex: sync.Mutex{},
		activeRepo:      make(map[string]map[int64]xpacman.Repository),
		bstore:          bstore,
		instancer:       instancers.New(corehub, bstore, dynHub),
	}

	return pm
}

func (p *PacMan) Start() error {

	reg := p.app.GetDeps().Registry().(*registry.Registry)
	p.repoBuilders = reg.GetRepoBuilders()

	return nil
}

func (p *PacMan) GetBprintFileStore() xpacman.BStore { return p.bstore }
func (p *PacMan) GetInstancer() xinstancer.Instancer { return p.instancer }
