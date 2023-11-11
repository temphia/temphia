package distro

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alecthomas/kong"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
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
	Init struct {
		SkipDBInit bool   `help:"Skip database initilization."`
		OrgName    string `help:"Org/Tenant name slug."`
		HttpPort   string `help:"HTTP Port for service."`
		DemoSeed   bool   `help:"Demo seed database."`
	} `cmd:"" help:"Initilizes data folder, create db and run migrations."`

	Start       struct{} `cmd:"" help:"Starts the application."`
	ActualStart struct{} `cmd:"" help:"Do not call, called internally."`
	Status      struct{} `cmd:"" help:"Status of App [running/stopped]"`

	ConfigFile string

	ctx *kong.Context
}

func RunAppCLI(args []string) error {

	os.Args = args

	cli := &AppCLi{}
	ctx := kong.Parse(cli)

	cli.ctx = ctx

	switch ctx.Command() {
	case "init":
		return cli.init()
	case "start":
		return cli.start()
	case "actual-start":
		return cli.actualStart()
	case "status":
		panic("not impl yet")
	default:
		panic("Not implemened command:" + ctx.Command())
	}

}

func (a *AppCLi) init() error {
	conf, err := a.readConfig()
	if err != nil {
		return err
	}

	confd := config.New(conf)
	err = confd.InitDataFolder()
	if err != nil {
		return err
	}

	if a.Init.SkipDBInit {
		return nil
	}

	ran := false

	switch conf.DatabaseConfig.Vendor {
	case store.VendorSqlite:
		_ran, err := common.InitSQLiteDB(conf.DatabaseConfig.Target)
		if err != nil {
			return err
		}

		ran = _ran

	default:
		return easyerr.Error("db vendor not implemented")
	}

	dapp, err := NewDistroApp(Options{
		Conf:        conf,
		Dev:         true,
		BuildFolder: nil,
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

	return nil
}

func (a *AppCLi) start() error {

	_, err := a.readConfig()
	if err != nil {
		return err
	}

	selfbinary, err := os.Executable()
	if err != nil {
		return err
	}

	logdsecret, err := xutils.GenerateRandomString(32)
	if err != nil {
		return err
	}

	// fixme => start logd service / injester

	cmd := exec.Command(selfbinary, "app", fmt.Sprintf("--config-file=%s", a.ConfigFile), "actual-start")
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Env = append(cmd.Env,
		fmt.Sprintf("TEMPHIA_LOGD_SECRET=%s", logdsecret),
		fmt.Sprintf("TEMPHIA_LOGD_PORT=%s", "5555"), // fixme => real logd
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	pp.Println("@starting_actual_process", cmd.String())

	return cmd.Run()

}

func (a *AppCLi) actualStart() error {

	conf, err := a.readConfig()
	if err != nil {
		return err
	}

	// fixme use TEMPHIA_LOGD_SECRET TEMPHIA_LOGD_PORT

	dapp, err := NewDistroApp(Options{
		Conf: conf,
		Dev:  false,
	})
	if err != nil {
		return easyerr.Wrap("could not build app", err)
	}

	return dapp.Run()
}
