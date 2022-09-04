package xtypes

type App interface {
	Run() error

	NodeId() string
	ClusterId() string
	DevMode() bool

	SingleTenant() bool
	TenantId() string

	HostAddrs(privatePriIp, privateSecIps, p2p bool) []string

	GetDeps() Deps
	GetServer() Server
	GetGlobalVar() GlobalVar
	Data() DataBox
}

type Deps interface {
	Registry() any
	RootController() any
	ControlPlane() any
	LogService() any
	CoreHub() any
	PlugKV() any
	Cabinet() any
	DynHub() any
	Fencer() any
	Engine() any
	Sockd() any
	Signer() any
	Pacman() any
	Courier() any
	NodeCache() any
}

type GlobalVar interface {
	Set(key string, val any)
	Get(key string) any
	Del(key string)
}

type Server interface {
	BindRoutes()
	Listen() error
	Close() error
}
