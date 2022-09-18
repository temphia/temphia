package app

import (
	"github.com/temphia/temphia/code/core/backend/engine"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/services/courierhub/courier"
	"github.com/temphia/temphia/code/core/backend/services/repohub"
	"github.com/temphia/temphia/code/core/backend/services/sockdhub"
	"github.com/temphia/temphia/code/core/backend/shared/signer"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
)

func (b *Builder) preCheck() error {
	deps := &b.app.deps

	if deps.registry == nil {
		return easyerr.Error("Empty Registry")
	}

	if deps.logService == nil {
		return easyerr.Error("Empty LogService")
	}

	if deps.controlPlane == nil {
		return easyerr.Error("Empty control plane")
	}

	if b.config == nil {
		return easyerr.Error("Empty Config")
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
	deps.engine = engine.New(b.app, *deps.logService.GetEngineLogger())

	deps.sockdhub = sockdhub.New(sockdx.Options{
		ServerIdent: b.app.clusterId,
		Logger:      deps.logService.GetServiceLogger("sockd"),
		Syncer:      nil,
		SysHelper:   nil,
	})

	deps.repoHub = repohub.New(b.app)
	deps.courier = courier.New()

	return nil

}
