package dev

import (
	"github.com/temphia/temphia/code/core/backend/libx/xutils"
	"github.com/temphia/temphia/code/distro"
	_ "github.com/temphia/temphia/code/distro/common"
)

func init() {
	handle(xutils.CreateIfNotExits("tmp/logs"))
}

func RunDev() {
	app := distro.NewApp(
		conf.AsConfig(), true, true,
	)

	err := app.Run()
	if err != nil {
		panic(err)
	}
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
