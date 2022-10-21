package dev

import (
	"os"

	"github.com/temphia/temphia/code/core/backend/libx/xutils"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/distro"
	_ "github.com/temphia/temphia/code/distro/common"
)

func init() {
	handle(xutils.CreateIfNotExits("tmp/logs"))
}

func RunDev() {
	dapp := distro.New(
		conf.AsConfig(), true, true,
	)

	applied, err := dapp.IsDbSchemaApplied()
	handle(err)

	if !applied {
		panic("Schema not applied")
	}

	if os.Getenv("TEMPHIA_SKIP_TENANT_SEED") != "1" {
		tseeded, err := dapp.IsTenantSeeded(xtypes.DefaultTenant)
		handle(err)
		if !tseeded {
			err = dapp.TenantSeed(xtypes.DefaultTenant)
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
