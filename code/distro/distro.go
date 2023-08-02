package distro

import (
	"os"
	"path"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/app"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/log"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/plane"
	"github.com/temphia/temphia/code/backend/stores"
	"github.com/temphia/temphia/code/backend/xtypes"
)

func NewDistroApp(conf *config.Config, dev bool) xtypes.App {

	var lservice *log.LogService

	reg := registry.New(true)
	sbuilder := stores.NewBuilder(stores.Options{
		Registry: reg,
		Config:   conf,
		LogBuilder: func() zerolog.Logger {
			return lservice.GetServiceLogger("store")
		},
	})

	err := sbuilder.Build()
	if err != nil {
		panic(err)
	}

	lite := plane.NewLite(sbuilder.CoreHub())

	logdSecret := os.Getenv("TEMPHIA_LOGD_SECRET")
	logdPort := os.Getenv("TEMPHIA_LOGD_PORT")

	logfolder := conf.LogFolder
	if logfolder == "" {
		logfolder = path.Join(conf.DataFolder, "logs")
	}

	lservice = log.New(log.LogOptions{
		LogdSecret: logdSecret,
		Folder:     logfolder,
		LogdPort:   logdPort,
		NodeId:     lite.GetNodeId(),
	})

	builder := app.NewBuilder()
	builder.SetConfig(conf)
	builder.SetLogger(lservice)
	builder.SetRegistry(reg)
	builder.SetXplane(lite)
	builder.SetStoreBuilder(sbuilder)
	builder.SetMode(dev)

	err = builder.Build()
	if err != nil {
		panic(err)
	}

	app := builder.GetApp()

	sbuilder.Inject(app)

	return app
}
