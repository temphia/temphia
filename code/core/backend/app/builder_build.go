package app

import (
	"github.com/temphia/temphia/code/core/backend/app/server"
	"github.com/temphia/temphia/code/core/backend/controllers"
	"github.com/temphia/temphia/code/core/backend/engine"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/services/courierhub/courier"
	"github.com/temphia/temphia/code/core/backend/services/repohub"
	"github.com/temphia/temphia/code/core/backend/services/shared/nodecache"
	"github.com/temphia/temphia/code/core/backend/services/shared/signer"
	"github.com/temphia/temphia/code/core/backend/services/sockdhub"
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

	if b.sbuilder == nil {
		return easyerr.Error("Empty store builder")
	}

	return nil
}

func (b *Builder) buildServer() error {

	b.app.deps.croot = controllers.New(controllers.Options{
		App:              b.app,
		OperatorUser:     b.config.OperatorName,
		OperatorPassword: b.config.OperatorPassword,
	})

	svc := server.New(server.Options{
		App:               b.app,
		GinEngine:         b.ginEngine,
		StaticHosts:       make(map[string]string),
		ResolveHostTenant: nil,
		RootHost:          "", // fixme => from config
		TenantHostBase:    "",
		Port:              ":4000",
		RootController:    b.app.deps.croot,
	})
	svc.BuildRoutes()

	b.app.deps.server = svc

	return nil
}

func (b *Builder) buildServices() error {

	err := b.preCheck()
	if err != nil {
		return err
	}

	deps := &b.app.deps

	deps.coreHub = b.sbuilder.CoreHub()
	deps.cabinetHub = b.sbuilder.CabHub()
	deps.dataHub = b.sbuilder.DataHub()

	deps.signer = signer.New([]byte(b.config.MasterKey), "temphia")
	deps.engine = engine.New(b.app, *deps.logService.GetEngineLogger())

	deps.sockdhub = sockdhub.New(sockdx.Options{
		ServerIdent: b.app.clusterId,
		Logger:      deps.logService.GetServiceLogger("sockd"),
		Syncer:      nil,
		SysHelper:   nil,
	})

	deps.nodeCache = nodecache.New("/tmp/mem1")
	deps.repoHub = repohub.New(b.app)
	deps.courier = courier.New()
	deps.plugKV = b.sbuilder.PlugKV()

	return deps.repoHub.Start()
}
