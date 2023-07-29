package stores

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/registry"

	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type Options struct {
	Registry   *registry.Registry
	Config     *config.Config
	LogBuilder func() zerolog.Logger
}

type Builder struct {
	registry   *registry.Registry
	config     *config.Config
	logBuilder func() zerolog.Logger

	cdb store.CoreDB
	pkv store.PlugStateKV

	cabhub  store.CabinetHub
	coreHub store.CoreHub
	dataHub dyndb.DataHub
}

func NewBuilder(opts Options) *Builder {
	return &Builder{
		registry:   opts.Registry,
		config:     opts.Config,
		logBuilder: opts.LogBuilder,
		cdb:        nil,
		pkv:        nil,
		cabhub:     nil,
		coreHub:    nil,
		dataHub:    nil,
	}
}

func (b *Builder) Build() error {

	b.registry.Freeze()

	//	storeBuilders := b.registry.GetStoreBuilders()
	// fixme => impl

	return nil
}

func (b *Builder) CabHub() store.CabinetHub {
	return b.cabhub
}

func (b *Builder) CoreHub() store.CoreHub {
	return b.coreHub
}

func (b *Builder) DataHub() dyndb.DataHub {
	return b.dataHub
}

func (b *Builder) PlugKV() store.PlugStateKV {
	return b.pkv
}

func (b *Builder) Inject(app xtypes.App) {
	b.coreHub.Inject(app)
	b.dataHub.Inject(app)
}
