package bunjs

import (
	"sync"

	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func init() {

	registry.SetExecutor("bunjs", func(app any) (etypes.ExecutorBuilder, error) {

		_app := app.(xtypes.App)
		deps := _app.GetDeps()

		conf := deps.Confd().(config.Confd)

		b := &Builder{
			chub:   deps.Cabinet().(store.CabinetHub),
			signer: deps.Signer().(service.Signer),
			waits:  make(map[string]chan *etypes.RemoteOptions),
			wlock:  sync.Mutex{},
			confd:  conf,
		}

		return b, nil
	})

}
