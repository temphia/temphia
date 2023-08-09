package dev

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/distro"
)

func Run() error {
	pp.Println("@i_am_dev")

	dapp, err := distro.NewDistroApp(&config.Config{
		ServerPort:      "4000",
		TenantId:        xtypes.DefaultTenant,
		EnableLocalDoor: true,
		DataFolder:      "./tmp",
		MasterKey:       "test123",

		DatabaseConfig: &config.StoreConfig{
			Name:     "sqlite",
			Vendor:   store.VendorSqlite,
			Provider: "sqlite",
			Target:   "main.db",
		},
	}, true)
	if err != nil {
		panic(err)
	}

	confd := dapp.Configd()

	err = confd.InitDataFolder()
	if err != nil {
		return err
	}

	err = xutils.CreateIfNotExits(confd.RootDataFolder())
	if err != nil {
		return err
	}

	err = dapp.Run()
	if err != nil {
		return err
	}

	return nil
}
