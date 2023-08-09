package distro

import (
	"github.com/temphia/temphia/code/backend/app/config"

	"github.com/temphia/temphia/code/backend/xtypes"
)

func (d *DistroApp) Run() error {
	return d.app.Run()
}

func (d *DistroApp) GetApp() xtypes.App {
	return d.app
}

func (d *DistroApp) Configd() config.Confd {
	return d.confd
}
