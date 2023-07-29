package localfs

import (
	"github.com/temphia/temphia/code/backend/app/registry"
	"github.com/temphia/temphia/code/backend/xtypes/store"

	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
)

type LocalFSBuilder struct {
	impl *NativeBlob
}

func init() {
	registry.SetStoreBuilders("localfs", NewLocalFSBuilder)
}

func NewLocalFSBuilder(opts store.BuilderOptions) (store.Store, error) {

	return &LocalFSBuilder{
		impl: &NativeBlob{
			rootFilePath: opts.Config.Target,
		},
	}, nil
}

func (u *LocalFSBuilder) Name() string {
	return ""
}

func (u *LocalFSBuilder) Supports(dbType store.StoreType) bool {
	switch dbType {
	case store.TypeBlobStore:
		return true
	default:
		return false
	}
}

func (u *LocalFSBuilder) CoreDB() store.CoreDB {
	return nil
}

func (u *LocalFSBuilder) SyncDB() store.SyncDB {
	return nil
}

func (u *LocalFSBuilder) StateDB() store.PlugStateKV {
	return nil
}

func (u *LocalFSBuilder) DynDB() dyndb.DynDB {
	return nil
}

func (u *LocalFSBuilder) FileStore() store.FileStore {
	return u.impl
}
