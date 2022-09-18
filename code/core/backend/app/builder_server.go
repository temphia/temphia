package app

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/app/server"
	"github.com/temphia/temphia/code/core/backend/controllers"
)

func (b *Builder) buildServer() error {

	b.app.deps.croot = controllers.New(controllers.Options{
		App:              b.app,
		OperatorUser:     "",
		OperatorPassword: "",
	})

	svr := server.New(server.Options{
		App:               b.app,
		GinEngine:         b.ginEngine,
		StaticHosts:       make(map[string]string),
		ResolveHostTenant: nil,
		RootHost:          "",
		TenantHostBase:    "",
		Port:              "",
		RootController:    b.app.deps.croot,
	})

	pp.Println(svr)

	return nil
}
