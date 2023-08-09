package distro

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/app"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/log"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/plane"
	"github.com/temphia/temphia/code/backend/stores"
	"github.com/temphia/temphia/code/backend/xtypes"
)

type DistroApp struct {
	app   xtypes.App
	confd config.Confd
}

func NewDistroApp(conf *config.Config, dev bool) (*DistroApp, error) {

	confd := config.New(conf)

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
	builder.SetMode(dev)

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
