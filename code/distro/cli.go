package distro

import (
	"github.com/alecthomas/kong"
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

	return nil
}

func (a *AppCLi) initDatabase() error {

	return nil
}

func (a *AppCLi) start() error {

	return nil
}

func (a *AppCLi) actualStart() error {

	return nil
}
