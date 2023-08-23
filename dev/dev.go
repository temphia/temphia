package dev

import (
	"os"
	"path"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/config"

	"github.com/temphia/temphia/code/distro/common"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/distro"

	// stores
	_ "github.com/temphia/temphia/code/backend/stores/localfs"
	_ "github.com/temphia/temphia/code/backend/stores/upperdb/vendors/sqlite"
)

func Run() error {
	err := copyFilesFromBuildProd()
	if err != nil {
		return err
	}

	pp.Println("@i_am_dev")

	conf := &config.Config{
		ServerPort:      ":4000",
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
	}
	confd := config.New(conf)

	err = confd.InitDataFolder()
	if err != nil {
		return err
	}

	ran, err := common.InitSQLiteDB(path.Join(confd.DBFolder(), conf.DatabaseConfig.Target))
	if err != nil {
		return err
	}

	devbuild := os.DirFS("../code/frontend/ui/build_dev/")
	dapp, err := distro.NewDistroApp(distro.Options{
		Conf:        conf,
		Dev:         true,
		BuildFolder: devbuild,
	})

	if err != nil {
		return err
	}

	if ran {
		err = dapp.SeedSuperUser()
		if err != nil {
			return err
		}

		err = dapp.SeedRepos()
		if err != nil {
			return err
		}

	}

	err = dapp.Run()
	if err != nil {
		return err
	}

	return nil
}
