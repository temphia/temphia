package xtypes

import (
	"fmt"
	"time"

	"github.com/flosch/go-humanize"
)

var Version string = "dev"
var Buildtime string = ""

func PrintBuildInfo() {
	fmt.Println("version ", Version)
	fmt.Println("Build Time ", Buildtime)
	if Buildtime != "" {
		t, _ := time.Parse(time.RFC3339, Buildtime)
		fmt.Println("Compiled ", humanize.Time(t))
	}
	fmt.Println("============")
}

type App interface {
	Run() error

	NodeId() string
	ClusterId() string
	DevMode() bool

	SingleTenant() bool
	StaticTenants() []string

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
	EngineHub() any
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
	Close() error
	BuildRoutes() error
	GetAdapterHub() any
}
