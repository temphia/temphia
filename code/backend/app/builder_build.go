package app

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/controllers"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
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

func (b *Builder) buildServer() error {

	b.app.deps.croot = controllers.New(controllers.Options{
		App: b.app,
	})

	b.app.deps.server = nil

	return nil
}

func (b *Builder) buildServices() error {

	err := b.preCheck()
	if err != nil {
		return err
	}

	deps := &b.app.deps

	// deps.coreHub = b.sbuilder.CoreHub()
	// deps.cabinetHub = b.sbuilder.CabHub()
	// deps.dataHub = b.sbuilder.DataHub()

	// deps.signer = signer.New([]byte(b.config.MasterKey), "temphia")
	// deps.engine = enginehub.New(b.app, deps.logService)

	// deps.sockdhub = sockdhub.New(sockdx.Options{
	// 	ServerIdent: b.app.clusterId,
	// 	Logger:      deps.logService.GetServiceLogger("sockd"),
	// 	Syncer:      nil,
	// 	SysHelper:   nil,
	// })

	//	deps.repoHub = repohub.New(b.app)
	//	deps.plugKV = b.sbuilder.PlugKV()
	err = deps.repoHub.Start()
	if err != nil {
		return err
	}

	exts := deps.registry.GetExecutorBuilder()

	exthub := b.extHandle

	for ename, extb := range exts {
		ext, err := extb(b.app, exthub)
		if err != nil {
			pp.Println("@extension_error", ename, err.Error())
			return err
		}

		b.app.deps.extensions[ename] = ext
	}

	b.app.global.Set("executors", exthub.executors)
	b.app.global.Set("modules", exthub.modules)
	b.app.global.Set("adapters", exthub.adapters)
	b.app.global.Set("scripts", exthub.adapters)

	return nil
}
