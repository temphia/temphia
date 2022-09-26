package stores

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/app/config"
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/services/cabhub"
	"github.com/temphia/temphia/code/core/backend/services/corehub"
	"github.com/temphia/temphia/code/core/backend/services/datahub"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type Options struct {
	Registry *registry.Registry
	Config   *config.Config
}

type Builder struct {
	registry *registry.Registry
	config   *config.Config

	stores     map[string]store.Store
	cdb        store.CoreDB
	pkv        store.PlugStateKV
	cabSources map[string]store.CabinetSource

	cabhub  store.CabinetHub
	coreHub store.CoreHub
	dataHub store.DataHub
}

func NewBuilder(opts Options) *Builder {
	return &Builder{
		registry:   opts.Registry,
		config:     opts.Config,
		stores:     make(map[string]store.Store),
		cabSources: make(map[string]store.CabinetSource),
		cdb:        nil,
		pkv:        nil,
		cabhub:     nil,
		coreHub:    nil,
		dataHub:    nil,
	}
}

func (b *Builder) Build() error {

	b.registry.Freeze()

	storeBuilders := b.registry.GetStoreBuilders()

	for _, ss := range b.config.Stores {
		sBuilder := storeBuilders[ss.Provider]
		if sBuilder == nil {
			pp.Println(ss)
			fmt.Println(storeBuilders)

			return easyerr.Error(fmt.Sprintf("Provider %s not found", ss.Provider))
		}

		store, err := sBuilder(ss)
		if err != nil {
			return easyerr.Wrap("err while building store", err)
		}
		b.stores[ss.Name] = store
	}

	if _, ok := b.stores[b.config.DefaultCabinet]; !ok {
		return easyerr.Error(fmt.Sprintf("default cabinet not found %s", b.config.DefaultCabinet))
	}

	if sp, ok := b.stores[b.config.Coredb]; !ok || !sp.Supports(store.TypeCoreDB) {
		return easyerr.Error(fmt.Sprintf("default coredb not loaded %s", b.config.Coredb))
	} else {

		b.cdb = sp.CoreDB()
		b.pkv = sp.StateDB()
	}

	for storek, _store := range b.stores {
		if !_store.Supports(store.TypeBlobStore) {
			continue
		}

		b.cabSources[storek] = _store.CabinetSource()
	}

	_, ok := b.cabSources[b.config.DefaultCabinet]
	if !ok {
		return easyerr.Error(fmt.Sprintf("default cabinet not loaded %s", b.config.DefaultCabinet))
	}

	dyns := make(map[string]store.DynDB)
	for k, s := range b.stores {
		if !s.Supports(store.TypeDynDB) {
			continue
		}

		dyns[k] = s.DynDB()
	}

	b.dataHub = datahub.New(dyns)
	b.cabhub = cabhub.New(b.cabSources, b.config.DefaultCabinet)
	b.coreHub = corehub.New(b.cdb)

	return nil
}

func (b *Builder) CabHub() store.CabinetHub {
	return b.cabhub
}

func (b *Builder) CoreHub() store.CoreHub {
	return b.coreHub
}

func (b *Builder) DataHub() store.DataHub {
	return b.dataHub
}

func (b *Builder) Inject(app xtypes.App) {
	b.coreHub.Inject(app)
	b.dataHub.Inject(app)
}
