package app

import (
	"github.com/temphia/temphia/code/core/backend/engine"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/services/courier"
	"github.com/temphia/temphia/code/core/backend/services/pacman"
	"github.com/temphia/temphia/code/core/backend/services/signer"
	"github.com/temphia/temphia/code/core/backend/services/sockd"
)

/*


	init, building order

	registry
	logger
	config
	xplane

	signer
	sockd

	stores
	-	coredb/hub
	-	dyndb/datahub
	-	engine


	services
		- pacman
		- nodecache
		- courier

	controller
	server


*/

func (b *Builder) preCheck() error {
	deps := &b.app.deps

	if deps.registry == nil {
		return easyerr.Error("Empty Registry")
	}

	if deps.logService == nil {
		return easyerr.Error("Empty LogService")
	}

	if b.config == nil {
		return easyerr.Error("Empty Config")
	}

	if deps.registry == nil {
		return easyerr.Error("Empty Registry")
	}

	return nil
}

func (b *Builder) build() error {

	err := b.preCheck()
	if err != nil {
		return err
	}

	deps := &b.app.deps

	deps.signer = signer.New([]byte(b.config.MasterKey), "temphia")

	err = b.buildServices()
	if err != nil {
		return err
	}

	return nil

}

func (b *Builder) buildServices() error {
	deps := &b.app.deps

	deps.engine = engine.New(b.app, *deps.logService.GetEngineLogger())

	deps.sockd = sockd.New(sockd.SockdOptions{
		ServerIdent: b.app.clusterId,
		Logger:      deps.logService.GetServiceLogger("sockd"),
		Syncer:      nil,
		SysHelper:   nil,
	})

	deps.pacman = pacman.New(b.app)
	deps.courier = courier.New()

	return nil
}
