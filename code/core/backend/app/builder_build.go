package app

import (
	"fmt"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/core/backend/engine"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/services/cabhub"
	"github.com/temphia/temphia/code/core/backend/services/corehub"
	"github.com/temphia/temphia/code/core/backend/services/courier"
	"github.com/temphia/temphia/code/core/backend/services/datahub"
	"github.com/temphia/temphia/code/core/backend/services/nodecache"
	"github.com/temphia/temphia/code/core/backend/services/pacman"
	"github.com/temphia/temphia/code/core/backend/services/signer"
	"github.com/temphia/temphia/code/core/backend/services/sockd"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

/*


	init, building order

	registry
	logger
	config
	xplane

	signer
	sockd

	stores
	-	coredb/hub
	-	dyndb/datahub
	-	engine


	services
		- pacman
		- nodecache
		- courier

	controller
	server


*/

func (b *Builder) preCheck() error {
	deps := &b.app.deps

	if deps.registry == nil {
		return easyerr.Error("Empty Registry")
	}

	if deps.logService == nil {
		return easyerr.Error("Empty LogService")
	}

	if b.config == nil {
		return easyerr.Error("Empty Config")
	}

	if deps.registry == nil {
		return easyerr.Error("Empty Registry")
	}

	return nil
}

func (b *Builder) build() error {

	err := b.preCheck()
	if err != nil {
		return err
	}

	deps := &b.app.deps

	deps.signer = signer.New([]byte(b.config.MasterKey), "temphia")

	err = b.buildStores()
	if err != nil {
		return err
	}

	err = b.buildServices()
	if err != nil {
		return err
	}

	return nil

}

func (b *Builder) buildStores() error {

	deps := &b.app.deps
	deps.registry.Freeze()

	deps.nodeCache = nodecache.New("tmp/store_files/ncache.db")

	storeBuilders := deps.registry.GetStoreBuilders()

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

	if _, ok := b.stores[b.config.Coredb]; !ok {
		return easyerr.Error(fmt.Sprintf("default coredb not found %s", b.config.Coredb))
	}

	if sp, ok := b.stores[b.config.Coredb]; !ok || !sp.Supports(store.TypeCoreDB) {
		return easyerr.Error(fmt.Sprintf("default coredb not loaded %s", b.config.Coredb))
	} else {

		b.cdb = sp.CoreDB()
		b.pkv = sp.StateDB()
	}

	csources := make(map[string]store.CabinetSource)
	for storek, _store := range b.stores {
		if !_store.Supports(store.TypeBlobStore) {
			continue
		}
		csources[storek] = _store.CabinetSource()
	}

	_, ok := csources[b.config.DefaultCabinet]
	if !ok {
		return easyerr.Error(fmt.Sprintf("default cabinet not loaded %s", b.config.DefaultCabinet))
	}

	fmt.Println("@=>", b.stores)

	deps.cabinetHub = cabhub.New(csources, b.config.DefaultCabinet)
	deps.coreHub = corehub.New(b.cdb, deps.sockd, deps.controlPlane)

	return nil
}

func (b *Builder) buildServices() error {
	deps := &b.app.deps

	deps.engine = engine.New(b.app, *deps.logService.GetEngineLogger())

	deps.sockd = sockd.New(sockd.SockdOptions{
		ServerIdent: b.app.clusterId,
		Logger:      deps.logService.GetServiceLogger("sockd"),
		Syncer:      nil,
		SysHelper:   nil,
	})

	{
		dyns := make(map[string]store.DynDB)
		for k, s := range b.stores {
			if !s.Supports(store.TypeDynDB) {
				continue
			}

			dyns[k] = s.DynDB()
		}

		deps.dataHub = datahub.New(b.app, dyns)
	}

	deps.coreHub = corehub.New(b.cdb, deps.sockd, deps.controlPlane)
	deps.pacman = pacman.New(b.app)
	deps.courier = courier.New()

	return nil
}
