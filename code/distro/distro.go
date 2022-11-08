package distro

import (
	"github.com/temphia/temphia/code/core/backend/app"
	"github.com/temphia/temphia/code/core/backend/app/config"
	"github.com/temphia/temphia/code/core/backend/app/log"
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/data"
	"github.com/temphia/temphia/code/core/backend/plane"
	"github.com/temphia/temphia/code/core/backend/stores"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

type App struct {
	App     xtypes.App
	CoreHub store.CoreHub
}

func New(conf *config.Config, dev, singleTenantMode bool) App {

	reg := registry.New(true)
	sbuilder := stores.NewBuilder(stores.Options{
		Registry: reg,
		Config:   conf,
	})

	err := sbuilder.Build()
	if err != nil {
		panic(err)
	}

	lite := plane.NewLite(sbuilder.CoreHub())

	builder := app.NewBuilder()
	builder.SetConfig(conf)
	builder.SetLogger(log.Default(lite))
	builder.SetRegistry(reg)
	builder.SetXplane(lite)
	builder.SetStoreBuilder(sbuilder)
	builder.SetDataBox(data.DefaultNew())
	builder.SetMode(dev)
	if singleTenantMode {
		builder.SetSingleTenant(xtypes.DefaultTenant)
	}

	err = builder.Build()
	if err != nil {
		panic(err)
	}

	app := builder.GetApp()

	sbuilder.Inject(app)

	return App{
		App:     app,
		CoreHub: app.GetDeps().CoreHub().(store.CoreHub),
	}

}

func (da *App) Run() error {
	return da.App.Run()
}
