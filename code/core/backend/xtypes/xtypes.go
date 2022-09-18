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
	GetMeshes() []Mesh
}

type Deps interface {
	Registry() any
	RootController() any
	ControlPlane() any
	LogService() any
	CoreHub() any
	PlugKV() any
	Cabinet() any
	DataHub() any
	Engine() any
	SockdHub() any
	RepoHub() any

	Courier() any
	NodeCache() any
	Signer() any
}

type GlobalVar interface {
	Set(key string, val any)
	Get(key string) any
	Del(key string)
}

type Server interface {
	Listen() error
	Addr() string
	Close() error
}
