package demo

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/data"
	"github.com/temphia/temphia/code/core/backend/xtypes"
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

	dapp := distro.New(Conf.AsConfig(), true, true)

	ok, err := dapp.IsTenantSeeded(xtypes.DefaultTenant)
	if err != nil {
		return err
	}
	if !ok {
		pp.Println("Looks like new fresh run lets seed tenant and user ")
		err = dapp.TenantSeed(xtypes.DefaultTenant)
		if err != nil {
			return err
		}

		err = dapp.SeedWildcardDomain(xtypes.DefaultTenant)
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

func Reset() error {
	return nil
}

func ClearLock() error {
	return nil
}

// private

func getPort() (int, error) {
	port := os.Getenv("TEMPHIA_DEMO_PG_PORT")
	if port == "" {
		// fixme check postgres file
		return getFreePort()
	}

	pport, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(pport), nil
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
