package distro

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/distro/climux"
	"github.com/temphia/temphia/code/distro/common"
)

func init() {

	climux.Register(&climux.CliAction{
		Name: "app",
		Help: "run app related actions",
		Func: RunAppCLI,
	})

}

type AppCLi struct {
	InitData struct {
		SkipDBInit bool `help:"Skip database initilization."`
		DemoSeed   bool `help:"Demo seed database."`
	} `cmd:"" help:"Initilizes data folder, create db and run migrations."`

	Start struct{} `cmd:"" help:"Starts the application."`

	ActualStart struct{} `cmd:"" help:"Do not call, called internally."`

	ConfigFile string

	ctx *kong.Context
}

func RunAppCLI(args []string) error {

	os.Args = args

	cli := &AppCLi{}
	ctx := kong.Parse(cli)

	cli.ctx = ctx

	switch ctx.Command() {
	case "init-data":
		return cli.initData()
	case "start":
		return cli.start()
	case "actual-start":
		return cli.actualStart()
	default:
		panic("Not implemened command:" + ctx.Command())
	}

}

func (a *AppCLi) initData() error {
	conf, err := a.readConfig()
	if err != nil {
		return err
	}

	confd := config.New(conf)
	err = confd.InitDataFolder()
	if err != nil {
		return err
	}

	if a.InitData.SkipDBInit {
		return nil
	}

	switch conf.DatabaseConfig.Vendor {
	case store.VendorSqlite:
		_, err = common.InitSQLiteDB(conf.DatabaseConfig.Target)
		if err != nil {
			return err
		}
	default:
		return easyerr.Error("db vendor not implemented")
	}

	return nil
}

func (a *AppCLi) start() error {

	conf, err := a.readConfig()
	if err != nil {
		return err
	}
	// fixme => start log injester and set TEMPHIA_LOGD_SECRET TEMPHIA_LOGD_PORT
	// call actualStart in subprocess

	dapp, err := NewDistroApp(Options{
		Conf: conf,
		Dev:  false,
	})
	if err != nil {
		return easyerr.Wrap("could not build app", err)
	}

	return dapp.Run()
}

func (a *AppCLi) actualStart() error {

	return nil
}

// private
