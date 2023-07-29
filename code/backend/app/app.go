package app

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
)

var _ xtypes.App = (*App)(nil)

type App struct {
	nodeId    string
	clusterId string
	tenantId  string
	devmode   bool
	deps      AppDeps
	global    Global
}

func (a *App) Run() error { return a.run() }

func (a *App) NodeId() string                 { return a.nodeId }
func (a *App) ClusterId() string              { return a.clusterId }
func (a *App) DevMode() bool                  { return a.devmode }
func (a *App) TenantId() string               { return a.tenantId }
func (a *App) GetDeps() xtypes.Deps           { return &a.deps }
func (a *App) GetServer() xtypes.Server       { return nil }
func (a *App) GetGlobalVar() xtypes.GlobalVar { return &a.global }

func (a *App) run() error {

	err := a.deps.controlPlane.Start()
	if err != nil {
		return err
	}

	err = a.deps.engine.Start()
	if err != nil {
		return err
	}

	// ectrl := a.deps.croot.EngineController()
	// ectrl.RunStartupHooks(a.tenantIds, time.Minute*2)

	pp.Println(a.
		deps.
		cabinetHub.
		Start(a.deps.controlPlane.GetMsgBus()))

	return a.deps.server.Listen()
}
