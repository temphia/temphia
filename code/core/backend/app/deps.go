package app

import (
	"github.com/temphia/temphia/code/core/backend/controllers"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

type AppDeps struct {
	registry     interface{}
	logService   logx.Service
	controlPlane xplane.ControlPlane
	engine       etypes.Engine
	sockd        sockdx.SockdCore
	coreHub      store.CoreHub
	plugKV       store.PlugStateKV
	signer       service.Signer
	courier      service.Courier
	cabinetHub   store.CabinetHub
	pacman       service.Pacman
	dynHub       store.DynHub
	nodeCache    service.NodeCache
	croot        *controllers.RootController
}

func (d *AppDeps) Registry() any       { return d.registry }
func (d *AppDeps) RootController() any { return d.croot }
func (d *AppDeps) ControlPlane() any   { return d.controlPlane }
func (d *AppDeps) LogService() any     { return d.logService }
func (d *AppDeps) CoreHub() any        { return d.coreHub }
func (d *AppDeps) PlugKV() any         { return d.plugKV }
func (d *AppDeps) Cabinet() any        { return d.cabinetHub }
func (d *AppDeps) DynHub() any         { return d.dynHub }
func (d *AppDeps) Engine() any         { return d.engine }
func (d *AppDeps) Sockd() any          { return d.sockd }
func (d *AppDeps) Signer() any         { return d.signer }
func (d *AppDeps) Pacman() any         { return d.pacman }
func (d *AppDeps) Courier() any        { return d.courier }
func (d *AppDeps) NodeCache() any      { return d.nodeCache }
