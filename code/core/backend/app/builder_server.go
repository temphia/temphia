package app

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/app/server"
)

func (b *Builder) buildServer() error {

	svr := server.New(server.Options{
		App:               b.app,
		GinEngine:         b.ginEngine,
		StaticHosts:       make(map[string]string),
		ResolveHostTenant: nil,
		RootHost:          "",
		TenantHostBase:    "",
		Port:              "",
		OperatorUser:      "",
		OperatorPassword:  "",
	})

	pp.Println(svr)

	return nil
}
