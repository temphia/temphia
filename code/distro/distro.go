package distro

import (
	"io/fs"
	"os"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/app"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/log"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/plane"
	"github.com/temphia/temphia/code/backend/stores"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/frontend/ui"
)

type DistroApp struct {
	app   xtypes.App
	confd config.Confd
}

type Options struct {
	Conf        *config.Config
	Dev         bool
	BuildFolder fs.FS
}

func NewDistroApp(opts Options) (*DistroApp, error) {
	if opts.BuildFolder == nil {
		opts.BuildFolder = ui.BuildProd
	}

	confd := config.New(opts.Conf)

	var lservice *log.LogService

	reg := registry.New(true)
	sbuilder := stores.NewBuilder(stores.Options{
		Registry: reg,
		Config:   opts.Conf,
		LogBuilder: func() zerolog.Logger {
			return lservice.GetServiceLogger("store")
		},
	})

	err := sbuilder.Build()
	if err != nil {
		return nil, err
	}

	lite := plane.NewLite(sbuilder.CoreHub())

	logdSecret := os.Getenv("TEMPHIA_LOGD_SECRET")
	logdPort := os.Getenv("TEMPHIA_LOGD_PORT")

	lservice = log.New(log.LogOptions{
		LogdSecret: logdSecret,
		Folder:     confd.LogFolder(),
		LogdPort:   logdPort,
		NodeId:     lite.GetNodeId(),
	})

	builder := app.NewBuilder()
	builder.SetConfigd(confd)
	builder.SetLogger(lservice)
	builder.SetRegistry(reg)
	builder.SetXplane(lite)
	builder.SetStoreBuilder(sbuilder)
	builder.SetMode(opts.Dev)
	builder.SetBuildFolder(opts.BuildFolder)
	builder.SetTenantId(xtypes.DefaultTenant)

	err = builder.Build()
	if err != nil {
		return nil, err
	}

	app := builder.GetApp()

	sbuilder.Inject(app)

	return &DistroApp{
		app:   app,
		confd: confd,
	}, nil

}
