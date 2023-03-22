package dev

import (
	"os"

	"github.com/temphia/temphia/code/backend/app/seeder"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/distro"
	_ "github.com/temphia/temphia/code/distro/common"
)

func init() {
	handle(xutils.CreateIfNotExits("tmp/logs"))
}

func RunDev() {
	dapp := distro.NewDistroApp(
		conf.AsConfig(), true, true,
	)

	seedApp := seeder.NewAppSeeder(dapp)

	applied, err := seedApp.IsDbSchemaApplied()
	handle(err)

	if !applied {
		panic("Schema not applied")
	}

	if os.Getenv("TEMPHIA_SKIP_TENANT_SEED") != "1" {
		tseeded, err := seedApp.IsTenantSeeded()
		handle(err)
		if !tseeded {
			err = seedApp.TenantSeed()
			handle(err)
		}
	}

	err = dapp.Run()
	handle(err)
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
