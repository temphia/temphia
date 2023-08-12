package app

import (
	"io/fs"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/app/xtension"
	"github.com/temphia/temphia/code/backend/stores"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

type Builder struct {
	app         *App
	confd       config.Confd
	config      *config.Config
	ginEngine   *gin.Engine
	assetsBuild fs.FS
	sbuilder    *stores.Builder
}

func NewBuilder() *Builder {

	return &Builder{
		app: &App{
			global: Global{
				globalVars: make(map[string]any),
				gmutex:     sync.Mutex{},
			},
			deps: AppDeps{
				extensions: make(map[string]xtension.Xtension),
			},
		},
	}
}

func (b *Builder) SetRegistry(reg *registry.Registry) {
	b.app.deps.registry = reg
}

func (b *Builder) SetLogger(log logx.Service) {
	b.app.deps.logService = log
}

func (b *Builder) SetXplane(xp xplane.ControlPlane) {
	b.app.deps.controlPlane = xp
}

func (b *Builder) SetConfigd(confd config.Confd) {
	b.confd = confd
	b.config = confd.GetConfig()
}

func (b *Builder) SetEngine(e *gin.Engine) {
	b.ginEngine = e
}

func (b *Builder) SetStoreBuilder(sbuilder *stores.Builder) {
	b.sbuilder = sbuilder
}

func (b *Builder) Build() error {

	err := b.buildServices()
	if err != nil {
		return err
	}

	return b.buildServer()
}

func (b *Builder) GetApp() xtypes.App {
	return b.app
}

func (b *Builder) SetMode(dev bool) {
	b.app.devmode = dev
}

func (b *Builder) SetTenantId(tenantId string) {
	b.app.tenantId = tenantId
}

func (b *Builder) SetBuildFolder(bf fs.FS) {
	b.assetsBuild = bf
}
