package dev

import (
	"fmt"
	"io"
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

func copyFilesFromBuildProd() error {
	prod := "../code/frontend/ui/build_prod/"
	target := "../code/frontend/ui/build_dev/"

	files, err := os.ReadDir(prod)
	if err != nil {
		return fmt.Errorf("error reading source directory: %w", err)
	}

	for _, file := range files {
		fmt.Println("|>", file.Name())

		if file.IsDir() || file.Name() == ".gitkeep" {
			continue
		}

		sourceFilePath := path.Join(prod, file.Name())
		targetFilePath := path.Join(target, file.Name())

		pfile, err := os.Open(sourceFilePath)
		if err != nil {
			return fmt.Errorf("error opening source file %s: %w", sourceFilePath, err)
		}
		defer pfile.Close()

		_, err = os.Stat(targetFilePath)
		if err == nil {
			fmt.Println("Already contains:", file.Name())
			continue
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("error checking target file %s: %w", targetFilePath, err)
		}

		tfile, err := os.Create(targetFilePath)
		if err != nil {
			return fmt.Errorf("error creating target file %s: %w", targetFilePath, err)
		}
		defer tfile.Close()

		_, err = io.Copy(tfile, pfile)
		if err != nil {
			return fmt.Errorf("error copying data to target file %s: %w", targetFilePath, err)
		}
	}

	return nil
}
