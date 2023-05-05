package app

import (
	"github.com/temphia/temphia/code/backend/app/server"
	"github.com/temphia/temphia/code/backend/controllers"
	"github.com/temphia/temphia/code/backend/engine"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/services/courierhub/courier"
	"github.com/temphia/temphia/code/backend/services/repohub"
	"github.com/temphia/temphia/code/backend/services/shared/nodecache"
	"github.com/temphia/temphia/code/backend/services/shared/signer"
	"github.com/temphia/temphia/code/backend/services/sockdhub"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
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

	staticHosts := b.config.NodeOptions.StaticHosts
	if staticHosts == nil {
		staticHosts = map[string]string{}
	}

	svc := server.New(server.Options{
		App:               b.app,
		GinEngine:         b.ginEngine,
		StaticHosts:       staticHosts,
		ResolveHostTenant: nil,
		RootHost:          b.config.NodeOptions.RootHost,
		TenantHostBase:    b.config.NodeOptions.TenantHostBase,
		Port:              b.config.NodeOptions.ServerPort,
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

	deps.nodeCache = nodecache.New(b.config.NodeOptions.NodeCache)
	deps.repoHub = repohub.New(b.app)
	deps.courier = courier.New()
	deps.plugKV = b.sbuilder.PlugKV()
	return deps.repoHub.Start()
}
