package app

import (
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/controllers"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

type AppDeps struct {
	registry     *registry.Registry
	logService   logx.Service
	controlPlane xplane.ControlPlane
	server       xtypes.Server

	signer   service.Signer
	engine   etypes.Engine
	sockdhub sockdx.Hub

	coreHub    store.CoreHub
	cabinetHub store.CabinetHub
	plugKV     store.PlugStateKV
	dataHub    store.DataHub
	repoHub    repox.Hub

	courier   service.Courier
	nodeCache service.NodeCache
	croot     *controllers.RootController
}

func (d *AppDeps) Registry() any       { return d.registry }
func (d *AppDeps) RootController() any { return d.croot }
func (d *AppDeps) ControlPlane() any   { return d.controlPlane }
func (d *AppDeps) LogService() any     { return d.logService }
func (d *AppDeps) CoreHub() any        { return d.coreHub }
func (d *AppDeps) PlugKV() any         { return d.plugKV }
func (d *AppDeps) Cabinet() any        { return d.cabinetHub }
func (d *AppDeps) DataHub() any        { return d.dataHub }
func (d *AppDeps) Engine() any         { return d.engine }
func (d *AppDeps) SockdHub() any       { return d.sockdhub }
func (d *AppDeps) Signer() any         { return d.signer }
func (d *AppDeps) RepoHub() any        { return d.repoHub }
func (d *AppDeps) Courier() any        { return d.courier }
func (d *AppDeps) NodeCache() any      { return d.nodeCache }
