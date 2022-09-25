package dev

import (
	"fmt"

	"github.com/temphia/temphia/code/core/backend/app"
	"github.com/temphia/temphia/code/core/backend/app/log"
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/plane"
	"github.com/temphia/temphia/code/core/backend/stores"
	_ "github.com/temphia/temphia/code/distro"
)

func RunDev() {

	fmt.Println("lets begin the madness")

	reg := registry.New(true)

	sbuilder := stores.NewBuilder(stores.Options{
		Registry: reg,
		Config:   conf.AsConfig(),
	})

	err := sbuilder.Build()
	if err != nil {
		panic(err)
	}

	lite := plane.NewLite(plane.LiteOptions{
		CoreHub: sbuilder.CoreHub(),
	})

	builder := app.NewBuilder()
	builder.SetConfig(nil)
	builder.SetLogger(log.New(log.LogOptions{
		LogdSecret: "",
		LogdPort:   "",
		Folder:     "",
		FilePrefix: "",
		NodeId:     lite.GetNodeId(),
	}))

	builder.SetRegistry(reg)
	builder.Xplane(lite)

	err = builder.BuildServer()
	if err != nil {
		panic(err)
	}

}
