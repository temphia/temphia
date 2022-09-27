package dev

import (
	"github.com/temphia/temphia/code/distro"
	_ "github.com/temphia/temphia/code/distro/common"
)

func RunDev() {
	app := distro.NewApp(
		conf.AsConfig(), true, true,
	)

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
