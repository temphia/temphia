package upper

import (
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/coredb"
	udyndb "github.com/temphia/temphia/code/backend/stores/upper/dyndb"

	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/dlock"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/tns"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/zenerator"
	"github.com/temphia/temphia/code/backend/stores/upper/plugkv"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/upper/db/v4"
)

// Adapter is adapter for upper bashed implementation
type Adapter struct {
	db           db.Session
	uvendor      ucore.UpperVendor
	innerCoreDB  store.CoreDB
	innerStateDb store.PlugStateKV
	innerDynDB   dyndb.DynDB
}

func NewAdapter(upvendor ucore.UpperVendor) func(opts store.BuilderOptions) (store.Store, error) {

	return func(opts store.BuilderOptions) (store.Store, error) {

		_tns := tns.New("shared")
		ztr := zenerator.New(opts.Config.Vendor, _tns)

		sess, err := upvendor.Db(opts.Config)
		if err != nil {
			return nil, err
		}

		return &Adapter{
			db:           sess,
			uvendor:      upvendor,
			innerCoreDB:  coredb.New(sess, opts.Config.Vendor),
			innerStateDb: plugkv.New(sess, dbutils.NewTxMgr(upvendor.NewTx), opts.Config.Vendor),
			innerDynDB: udyndb.New(ucore.DynDBOptions{
				Session:       sess,
				TxnManager:    dbutils.NewTxMgr(upvendor.NewTx),
				DynGen:        ztr,
				TNS:           _tns,
				SharedLock:    dlock.New(""),
				Vendor:        opts.Config.Vendor,
				LoggerBuilder: opts.LogBuilder,
			}),
		}, nil
	}

}

func (u *Adapter) Name() string {
	return ""
}

func (u *Adapter) Supports(dbType store.StoreType) bool {
	switch dbType {
	case store.TypeCoreDB, store.TypeStateDB:
		return true
	case store.TypeDynDB:
		return true
	case store.TypeBlobStore:
		return false
	default:
		return false
	}
}

func (u *Adapter) Migrate() error {
	return nil
}

func (u *Adapter) CoreDB() store.CoreDB {
	return u.innerCoreDB
}

func (u *Adapter) SyncDB() store.SyncDB {
	return nil
}

func (u *Adapter) StateDB() store.PlugStateKV {
	return u.innerStateDb
}

func (u *Adapter) DynDB() dyndb.DynDB {
	return u.innerDynDB
}

func (u *Adapter) CabinetSource() store.CabinetSource {
	return nil
}
