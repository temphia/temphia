package demo

import (
	"fmt"
	"os"
	"syscall"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/seeder"
	"github.com/temphia/temphia/code/backend/data"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/distro"
	"github.com/temphia/temphia/code/distro/embedpg"
)

func RunDemo() error {

	port, err := getPort()
	if err != nil {
		return nil
	}

	dpg := embedpg.New("temphia-data/pgdata", port)

	Conf.Database.Port = fmt.Sprintf("%d", dpg.GetPort())

	pp.Println("PORT ", int(dpg.GetPort()))

	err = dpg.Start()
	if err != nil {
		return err
	}

	pp.Println("POSTGRES STARTED")

	defer func() {
		pp.Println("STOPPING POSTGRES", dpg.Stop())
	}()

	pp.Println("database started port ", int(dpg.GetPort()))

	out, err := data.DataDir.ReadFile("schema/postgres.sql")
	if err != nil {
		return err
	}

	err = dpg.RunSchema(string(out))
	if err != nil {
		return err
	}

	dapp := distro.NewDistroApp(Conf.AsConfig(), true, true)

	go setUpHandler(func(signal os.Signal) {
		if signal == syscall.SIGTERM {
			fmt.Println("Got kill signal. ")
			fmt.Println("Program will terminate now.")
			fmt.Println(dpg.Stop())
			os.Exit(0)
		} else if signal == syscall.SIGINT {
			fmt.Println("Got CTRL+C signal")
			fmt.Println("Closing.")
			fmt.Println(dpg.Stop())
			os.Exit(0)
		} else {
			fmt.Println("Ignoring signal: ", signal)
		}
	})

	err = runSeed(dapp)
	if err != nil {
		return err
	}

	err = dapp.Run()
	if err != nil {
		return err
	}

	return nil
}

func Reset() error {
	return os.Remove("temphia-data")
}

func ClearLock() error {
	return os.Remove("temphia-data/pgdata/data/postmaster.pid")
}

func runSeed(app xtypes.App) error {

	seedapp := seeder.NewAppSeeder(app)
	ok, err := seedapp.IsTenantSeeded()
	if err != nil {
		return err
	}
	if !ok {

		err = seedapp.TenantSeed()
		if err != nil {
			return err
		}

		err = seedapp.SeedWildcardDomain()
		if err != nil {
			return err
		}

		err = seedapp.SeedRepo()
		if err != nil {
			return err
		}

		err := seedapp.SendWelcome()
		if err != nil {
			return err
		}

	}

	return nil
}
