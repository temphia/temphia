package distro

import (
	"encoding/json"
	"os"

	"github.com/alecthomas/kong"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/distro/common"
)

type AppCLi struct {
	InitData struct {
		SkipDBInit bool `help:"Skip database initilization."`
		DemoSeed   bool `help:"Demo seed database."`
	} `cmd:"" help:"Initilizes data folder and create db and run migrations."`

	InitDB struct {
		CustomMigration string `help:"Custom migration folder."`
	} `cmd:"" help:"Create db and run migrations."`

	Start struct{} `cmd:"" help:"Starts the application."`

	ActualStart struct{} `cmd:"" help:"Do not call, called internally."`

	ConfigFile string
}

func RunAppCLI(args []string) error {

	cli := &AppCLi{}
	ctx := kong.Parse(cli)

	switch ctx.Command() {
	case "init-data":
		return cli.initData()
	case "init-db":
		return cli.initDatabase()
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
		err = common.InitSQLiteDB(conf.DatabaseConfig.Target)
		if err != nil {
			return err
		}
	default:
		return easyerr.Error("db vendor not implemented")
	}

	return nil
}

func (a *AppCLi) initDatabase() error {

	conf, err := a.readConfig()
	if err != nil {
		return err
	}

	switch conf.DatabaseConfig.Vendor {
	case store.VendorSqlite:
		err = common.InitSQLiteDB(conf.DatabaseConfig.Target)
		if err != nil {
			return err
		}
	default:
		return easyerr.Error("db vendor not implemented")
	}

	return nil
}

func (a *AppCLi) start() error {

	return nil
}

func (a *AppCLi) actualStart() error {

	return nil
}

// private

func (a *AppCLi) readConfig() (*config.Config, error) {

	if a.ConfigFile == "" {
		return nil, easyerr.Error("--config-file not passed")
	}

	out, err := os.ReadFile(a.ConfigFile)
	if err != nil {
		return nil, easyerr.Wrap("err reading config file", err)
	}

	conf := &config.Config{}
	err = json.Unmarshal(out, &conf)
	if err != nil {
		return nil, easyerr.Wrap("err parsing config JSON", err)
	}

	return conf, nil
}
