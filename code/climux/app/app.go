package app

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alecthomas/kong"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/climux"
	"github.com/temphia/temphia/code/distro"
)

func init() {

	climux.Register(&climux.Action{
		Name: "app",
		Help: "run app related actions",
		Func: RunAppCLI,
	})

}

type AppCLi struct {
	Init struct {
		SkipDBInit   bool   `help:"Skip database initilization."`
		OrgName      string `help:"Org/Tenant name slug."`
		HttpPort     string `help:"HTTP Port for service."`
		RootDomain   string `help:"Root domain to run main temphia app. like example.com"`
		RunnerDomain string `help:"Runner domain to run apps under like myapp.example.com. so it should support wildcard subdomain in DNS. Pass as example.com here."`
		DemoSeed     bool   `help:"Demo seed database."`
	} `cmd:"" help:"Initilizes data folder, create db and run migrations."`

	Start       struct{} `cmd:"" help:"Starts the application."`
	ActualStart struct{} `cmd:"" help:"Do not call, called internally."`
	Status      struct{} `cmd:"" help:"Status of App [running/stopped]"`

	ConfigFile string

	ctx *kong.Context
}

func RunAppCLI(cctx climux.Context) error {
	os.Args = cctx.Args

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
		pp.Println("read config error")
		return err
	}

	confd := config.New(conf)
	err = confd.InitDataFolder()
	if err != nil {
		pp.Println("init folder error")
		return err
	}

	if a.Init.SkipDBInit {
		return nil
	}

	ran := false

	switch conf.DatabaseConfig.Vendor {
	case store.VendorSqlite:
		_ran, err := distro.InitSQLiteDB(conf.DatabaseConfig.Target)
		if err != nil {
			return err
		}

		ran = _ran

	default:
		return easyerr.Error("db vendor not implemented")
	}

	dapp, err := distro.NewDistroApp(distro.Options{
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
		fmt.Sprintf("%s=%s", xtypes.EnvLogdSecret, logdsecret),
		fmt.Sprintf("%s=%s", xtypes.EnvLogdSecret, "5555"), // fixme => real logd
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	pp.Println("@starting_actual_process", cmd.String())

	/*

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)
		go func() {
			pp.Println("@WAITING_SIGNAL")

			s := <-sigchan
			pp.Println("@GOT_SIGNAL", s.String())

			if cmd.Process == nil {
				return
			}

			pp.Println("@killing pid:", cmd.Process.Pid)

			cmd.Process.Signal(s)
			cmd.Process.Kill()
		}()

	*/

	err = cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()

}

func (a *AppCLi) actualStart() error {

	conf, err := a.readConfig()
	if err != nil {
		return err
	}

	// fixme use TEMPHIA_LOGD_SECRET TEMPHIA_LOGD_PORT

	dapp, err := distro.NewDistroApp(distro.Options{
		Conf: conf,
		Dev:  false,
	})
	if err != nil {
		return easyerr.Wrap("could not build app", err)
	}

	return dapp.Run()
}

func (a *AppCLi) readConfig() (*config.Config, error) {

	if a.ConfigFile == "" {

		if a.ctx.Command() == "init" {
			os.Mkdir(distro.TemphiaStateFolder, os.FileMode(0777))
			os.WriteFile(distro.TemphiaConfigFile, distro.GetConfig(
				a.Init.HttpPort,
				a.Init.OrgName,
				a.Init.RootDomain,
				a.Init.RunnerDomain,
				os.Getenv(xtypes.EnvAppInitSecret),
				"",
			), os.FileMode(0666))
			a.ConfigFile = distro.TemphiaConfigFile

		} else {
			if xutils.FileExists("./", distro.TemphiaConfigFile) {
				a.ConfigFile = distro.TemphiaConfigFile
			}
		}

	}

	return distro.ReadConfig(a.ConfigFile)
}
