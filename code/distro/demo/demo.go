package demo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/seeder"
	"github.com/temphia/temphia/code/backend/data"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/distro"
	"github.com/temphia/temphia/code/distro/embedpg"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

func RunDemo() error {

	switch Conf.Database.Vendor {
	case store.VendorPostgres:
		err := initPg()
		if err != nil {
			return err
		}
	case store.VendorSqlite:
		err := initSqlite()
		if err != nil {
			return err
		}
	default:
		panic("Not supported vendor " + Conf.Database.Vendor)
	}

	dapp := distro.NewDistroApp(Conf.AsConfig(), true, true)

	err := runSeed(dapp)
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

		err = seedExtraUser(seedapp)
		if err != nil {
			return err
		}
	}

	return nil
}

func seedExtraUser(sapp *seeder.AppSeeder) error {

	err := sapp.AddNewUserGroup("LabRats", "laprats")
	if err != nil {
		return err
	}

	err = sapp.AddNewUser("toprat1", "Top Rat1", "top1@lrat.com", "cheesy123", "laprats")
	if err != nil {
		return err
	}

	err = sapp.AddNewUser("toprat2", "Top Rat2", "top2@lrat.com", "cheesy123", "laprats")
	if err != nil {
		return err
	}

	err = sapp.AddNewUser("toprat3", "Top Rat3", "top3@lrat.com", "cheesy123", "laprats")
	if err != nil {
		return err
	}

	err = sapp.AddNewUser("toprat4", "Top Rat4", "top4@lrat.com", "cheesy123", "laprats")
	if err != nil {
		return err
	}

	return nil
}

func initSqlite() error {

	sess, err := sqlite.Open(sqlite.ConnectionURL{
		Database: Conf.Database.HostPath,
	})

	if err != nil {
		return err
	}

	ok, err := sess.Collection("tenants").Exists()
	if err != nil {
		if !errors.Is(err, db.ErrCollectionDoesNotExist) {
			return err
		}
	}

	if ok {
		return nil
	}

	conn := sess.Driver().(*sql.DB)

	out, err := data.DataDir.ReadFile("schema/sqlite.sql")
	if err != nil {
		return err
	}

	ctx, cfunc := context.WithTimeout(context.Background(), time.Minute*2)
	defer cfunc()
	_, err = conn.ExecContext(ctx, string(out))
	return err
}

func initPg() error {
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

	return nil
}
