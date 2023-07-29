package app

import (
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/xutils"
	"github.com/temphia/temphia/code/backend/xtypes"
)

var _ xtypes.App = (*App)(nil)

type App struct {
	nodeId           string
	clusterId        string
	singleTenantMode bool
	tenantIds        []string
	devmode          bool
	port             string
	deps             AppDeps
	global           Global
	meshes           []xtypes.Mesh
}

func (a *App) Run() error { return a.run() }

func (a *App) NodeId() string                 { return a.nodeId }
func (a *App) ClusterId() string              { return a.clusterId }
func (a *App) DevMode() bool                  { return a.devmode }
func (a *App) SingleTenant() bool             { return a.singleTenantMode }
func (a *App) StaticTenants() []string        { return a.tenantIds }
func (a *App) GetDeps() xtypes.Deps           { return &a.deps }
func (a *App) GetServer() xtypes.Server       { return nil }
func (a *App) GetGlobalVar() xtypes.GlobalVar { return &a.global }
func (a *App) GetMeshes() []xtypes.Mesh       { return a.meshes }
func (a *App) HostAddrs(privatePriIp, privateSecIps, p2p bool) []string {
	return a.hostAddrs(privatePriIp, privateSecIps, p2p)
}

// private
func (a *App) hostAddrs(privatePriIp, privateSecIps, p2p bool) []string {
	hosts := make([]string, 0, 10)

	host, err := os.Hostname()
	if err == nil {
		hosts = append(hosts, xutils.BuildAddr(host, a.port))
	}

	// fixme p2p

	if !privatePriIp && !privateSecIps {
		return hosts
	}

	ips := xutils.GetLocalIPs()
	for _, ip := range ips {
		hosts = append(hosts, xutils.BuildAddr(ip, a.port))
	}

	for _, mesh := range a.meshes {
		addrs := mesh.GetAddress()
		for _, ip := range addrs {
			hosts = append(hosts, xutils.BuildAddr(ip, a.port))
		}
	}

	return hosts
}

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
