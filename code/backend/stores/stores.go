package stores

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/app/registry"
	cabinethub "github.com/temphia/temphia/code/backend/hub/cabinet"
	corehub "github.com/temphia/temphia/code/backend/hub/coredb"
	datahub "github.com/temphia/temphia/code/backend/hub/dyndb"

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
	opts Options

	pkv     store.PlugStateKV
	coreHub store.CoreHub
	dataHub dyndb.DataHub

	cabhub store.CabinetHub
}

func NewBuilder(opts Options) *Builder {
	return &Builder{
		opts:    opts,
		pkv:     nil,
		cabhub:  nil,
		coreHub: nil,
		dataHub: nil,
	}
}

func (b *Builder) Build() error {

	b.opts.Registry.Freeze()

	storeBuilders := b.opts.Registry.GetStoreBuilders()

	dbconf := b.opts.Config.DatabaseConfig

	dbBuilder := storeBuilders[dbconf.Provider]

	dbstore, err := dbBuilder(store.BuilderOptions{
		Config:     dbconf,
		LogBuilder: b.opts.LogBuilder,
	})
	if err != nil {
		return err
	}

	b.coreHub = corehub.New(dbstore.CoreDB())
	b.pkv = dbstore.StateDB()
	b.dataHub = datahub.New(dbstore.DynDB())

	fsBuilder := storeBuilders[dbconf.Provider]
	fsstore, err := fsBuilder(store.BuilderOptions{
		Config:     b.opts.Config.FileStoreConfig,
		LogBuilder: b.opts.LogBuilder,
	})
	if err != nil {
		return err
	}

	cabinethub.New(fsstore.FileStore())

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
