package app

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/app/config"
	"github.com/temphia/temphia/code/core/backend/app/registry"
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx"
	"github.com/temphia/temphia/code/core/backend/xtypes/xplane"
)

type Builder struct {
	app       *App
	config    *config.Config
	ginEngine *gin.Engine
}

func NewBuilder() *Builder {
	return &Builder{
		app: &App{
			global: Global{
				globalVars: make(map[string]any),
				gmutex:     sync.Mutex{},
			},
			deps:   AppDeps{},
			meshes: make([]xtypes.Mesh, 0, 2),
		},
	}
}

func (b *Builder) SetConfig(conf *config.Config) {
	b.config = conf
	if b.config.NodeOptions.TenantId != "" {
		b.app.singleTenantMode = true
		b.app.tenantId = b.config.NodeOptions.TenantId
	}
}

func (b *Builder) SetRegistry(reg *registry.Registry) {
	b.app.deps.registry = reg
}

func (b *Builder) SetLogger(log logx.Service) {
	b.app.deps.logService = log
}

func (b *Builder) Xplane(xp xplane.ControlPlane) {
	b.app.deps.controlPlane = xp
}

func (b *Builder) SetEngine(e *gin.Engine) {
	b.ginEngine = e
}

func (b *Builder) Build() error {
	return b.build()
}

func (b *Builder) BuildServer() error {
	return b.buildServer()
}

func (b *Builder) GetApp() xtypes.App {
	return b.app
}
