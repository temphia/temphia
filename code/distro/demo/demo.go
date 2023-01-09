package demo

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/data"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/distro"
	"github.com/temphia/temphia/code/distro/embedpg"
)

func Main() error {

	dpg := embedpg.New("tmp/pgdata")

	Conf.Database.Port = fmt.Sprintf("%d", dpg.GetPort())

	pp.Println("PORT ", dpg.GetPort())

	err := dpg.Start()
	if err != nil {
		return err
	}

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
		err = dapp.TenantSeed(xtypes.DefaultTenant)
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
