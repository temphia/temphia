package dev

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/app"
	"github.com/temphia/temphia/code/core/backend/app/config"
	"github.com/temphia/temphia/code/core/backend/app/log"
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/data"
	"github.com/temphia/temphia/code/core/backend/plane"
	"github.com/temphia/temphia/code/core/backend/stores"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	_ "github.com/temphia/temphia/code/distro/common"
)

func RunDev() {
	app := NewApp(conf.AsConfig())
	pp.Println(app.Run())
}

func NewApp(conf *config.Config) xtypes.App {

	fmt.Println("lets begin the madness")

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

	err = builder.Build()
	if err != nil {
		panic(err)
	}

	return builder.GetApp()
}
