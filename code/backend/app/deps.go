package app

import (
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/app/xtension"
	"github.com/temphia/temphia/code/backend/controllers"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
	"github.com/temphia/temphia/code/backend/xtypes/xserver"
)

type AppDeps struct {
	registry     *registry.Registry
	confd        config.Confd
	logService   logx.Service
	controlPlane xplane.ControlPlane
	server       xserver.Server

	signer   service.Signer
	engine   etypes.EngineHub
	sockdhub sockdx.Hub

	coreHub    store.CoreHub
	cabinetHub store.CabinetHub
	plugKV     store.PlugStateKV
	dataHub    dyndb.DataHub
	repoHub    xpacman.Pacman

	croot      *controllers.RootController
	extensions map[string]xtension.Xtension
}

func (d *AppDeps) start() error {

	err := d.controlPlane.Start()
	if err != nil {
		return err
	}

	err = d.engine.Start()
	if err != nil {
		return err
	}

	err = d.cabinetHub.
		Start(d.controlPlane.GetMsgBus())
	if err != nil {
		return err
	}

	for _, xt := range d.extensions {
		err = xt.Start()
		if err != nil {
			return err
		}
	}

	// ectrl := a.deps.croot.EngineController()
	// ectrl.RunStartupHooks(a.tenantIds, time.Minute*2)

	return nil
}

func (d *AppDeps) Confd() any          { return d.confd }
func (d *AppDeps) Registry() any       { return d.registry }
func (d *AppDeps) RootController() any { return d.croot }
func (d *AppDeps) ControlPlane() any   { return d.controlPlane }
func (d *AppDeps) LogService() any     { return d.logService }
func (d *AppDeps) CoreHub() any        { return d.coreHub }
func (d *AppDeps) PlugKV() any         { return d.plugKV }
func (d *AppDeps) Cabinet() any        { return d.cabinetHub }
func (d *AppDeps) DataHub() any        { return d.dataHub }
func (d *AppDeps) EngineHub() any      { return d.engine }
func (d *AppDeps) SockdHub() any       { return d.sockdhub }
func (d *AppDeps) Signer() any         { return d.signer }
func (d *AppDeps) RepoHub() any        { return d.repoHub }
func (d *AppDeps) Extensions() any     { return d.extensions }
