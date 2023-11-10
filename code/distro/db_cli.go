package distro

import (
	"os"
	"path"

	"github.com/alecthomas/kong"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/distro/climux"
	"github.com/temphia/temphia/code/distro/common"
)

func init() {

	climux.Register(&climux.CliAction{
		Name: "db",
		Help: "run db related actions",
		Func: RunDatabaseCLI,
	})

}

type DatabaseCLI struct {
	Init struct {
		CustomMigration string `help:"Custom migration folder."`
	} `cmd:"" help:"Create db and run migrations. you should be using [temphia app init-data]."`

	Reset struct {
		CustomMigration string `help:"Custom migration folder."`
	} `cmd:"" help:"delete old db and init new db"`

	Migrate struct {
		CustomMigration string `help:"Custom migration folder."`
	} `cmd:"" help:"Run new migrations."`

	Delete struct {
		CustomMigration string `help:"Custom migration folder."`
	} `cmd:"" help:"Delete database."`

	ConfigFile string

	ctx *kong.Context
}

func RunDatabaseCLI(args []string) error {

	os.Args = args

	cli := &DatabaseCLI{}
	ctx := kong.Parse(cli)

	cli.ctx = ctx

	switch ctx.Command() {

	case "init":
		return cli.initDatabase()
	case "reset", "migrate", "rollback", "delete":
		panic("Not implemented currently :" + ctx.Command())
	default:
		panic("Not implemened :" + ctx.Command())
	}

}

func (a *DatabaseCLI) initDatabase() error {

	conf, err := readConfig(a.ConfigFile)
	if err != nil {
		return err
	}

	confd := config.New(conf)

	switch conf.DatabaseConfig.Vendor {
	case store.VendorSqlite:
		_, err = common.InitSQLiteDB(path.Join(confd.DBFolder(), conf.DatabaseConfig.Target))
		if err != nil {
			return err
		}
	default:
		return easyerr.Error("db vendor not implemented")
	}

	return nil
}
