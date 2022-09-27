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
)

func NewApp(conf *config.Config, dev, singleTenantMode bool) xtypes.App {

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

	return builder.GetApp()
}
