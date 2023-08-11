package app

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/server"
	"github.com/temphia/temphia/code/backend/app/xtension"
	"github.com/temphia/temphia/code/backend/controllers"
	enginehub "github.com/temphia/temphia/code/backend/hub/engine"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/services/pacman"
	"github.com/temphia/temphia/code/backend/services/signer"
	"github.com/temphia/temphia/code/backend/services/sockd"
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

	if b.config.DatabaseConfig == nil {
		return easyerr.Error("Empty Database Config")
	}

	return nil
}

func (b *Builder) buildServer() error {

	b.app.deps.croot = controllers.New(controllers.Options{
		App: b.app,
	})

	server := server.New(server.Options{
		RootDomain:     b.config.RootDomain,
		RunnerDomain:   b.config.RunnerDomain,
		App:            b.app,
		GinEngine:      b.ginEngine,
		RootController: b.app.deps.croot,
		Port:           b.config.ServerPort,
		LocalSocket:    b.confd.LocalSocket(),
		BuildFS:        b.assetsBuild,
	})

	b.app.deps.server = server

	return nil
}

func (b *Builder) buildServices() error {

	err := b.preCheck()
	if err != nil {
		return err
	}

	deps := &b.app.deps

	deps.confd = b.confd
	deps.coreHub = b.sbuilder.CoreHub()
	deps.cabinetHub = b.sbuilder.CabHub()
	deps.dataHub = b.sbuilder.DataHub()

	deps.signer = signer.New([]byte(b.config.MasterKey), "temphia")
	deps.engine = enginehub.New(b.app, deps.logService)

	deps.sockdhub = sockd.New(sockdx.Options{
		ServerIdent: b.app.clusterId,
		Logger:      deps.logService.GetServiceLogger("sockd"),
		Syncer:      nil,
		SysHelper:   nil,
	})

	deps.repoHub = pacman.New(b.app)

	deps.plugKV = b.sbuilder.PlugKV()
	err = deps.repoHub.Start()
	if err != nil {
		return err
	}

	Xtbuilders := deps.registry.GetXtensionBuilder()
	ctx := xtension.NewHandle()

	for ename, extb := range Xtbuilders {
		ext, err := extb(ctx)
		if err != nil {
			pp.Println("@extension_error", ename, err.Error())
			return err
		}
		if ext != nil {
			deps.extensions[ename] = ext
		}
	}

	xtdata := xtension.GetContextData(ctx)

	b.app.global.Set("executors", xtdata.Executors)
	b.app.global.Set("modules", xtdata.Modules)
	b.app.global.Set("adapters", xtdata.Adapters)
	b.app.global.Set("scripts", xtdata.Scripts)

	return nil
}
