package dyndb

import (
	"github.com/temphia/temphia/code/core/backend/libx/dbutils"
	"github.com/temphia/temphia/code/core/backend/stores/upper/dyndb/dyncore"
	"github.com/temphia/temphia/code/core/backend/stores/upper/dyndb/tns"
	"github.com/temphia/temphia/code/core/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
	"github.com/upper/db/v4"
)

var _ store.DynDB = (*DynDB)(nil)

type DynDB struct {
	session    db.Session
	sharedLock service.DyndbLock
	dyngen     ucore.Zenerator
	txn        dbutils.TxManager
	tns        tns.TNS
	vendor     string
	cache      store.DCache
}

func New(opts ucore.DynDBOptions) *DynDB {
	d := &DynDB{
		session:    opts.Session,
		sharedLock: opts.SharedLock,
		txn:        opts.TxnManager,
		dyngen:     opts.DynGen,
		tns:        opts.TNS,
		vendor:     store.VendorPostgres, // fixme =>  from config
		cache:      nil,
	}

	d.cache = dyncore.NewCache(d.ListColumns)
	return d
}

func (d *DynDB) txOr(txid uint32, fn func(sess db.Session) error) error {
	return d.txn.TxOr(txid, d.session, fn)
}

func (d *DynDB) dataTableGroups() db.Collection {
	return dyncore.GroupTable(d.session)
}

func (d *DynDB) dataTables() db.Collection {
	return dyncore.Table(d.session)
}

func (d *DynDB) dataTableColumns() db.Collection {
	return dyncore.TableColumn(d.session)
}

func (d *DynDB) viewTable() db.Collection {
	return d.session.Collection("data_views")
}

func (d *DynDB) hookTable() db.Collection {
	return d.session.Collection("data_hooks")
}

func (d *DynDB) NewRow(txid uint32, req store.NewRowReq) (int64, error) {
	return d.newRow(txid, req)
}

func (d *DynDB) GetRow(txid uint32, req store.GetRowReq) (map[string]interface{}, error) {
	return d.getRow(txid, req)
}

func (d *DynDB) UpdateRow(txid uint32, req store.UpdateRowReq) (map[string]interface{}, error) {
	return d.updateRow(txid, req)
}

func (d *DynDB) DeleteRows(txid uint32, req store.DeleteRowReq) error {
	return d.deleteRows(txid, req)
}

func (d *DynDB) SimpleQuery(txid uint32, req store.SimpleQueryReq) (*store.QueryResult, error) {
	return d.simpleQuery(txid, req)
}

func (d *DynDB) FTSQuery(txid uint32, req store.FTSQueryReq) (*store.QueryResult, error) {
	return d._FTSQuery(txid, req)
}

func (d *DynDB) TemplateQuery(txid uint32, req store.TemplateQueryReq) (*store.QueryResult, error) {
	return d.templateQuery(txid, req)
}

func (d *DynDB) RawQuery(txid uint32, req store.RawQueryReq) (interface{}, error) {
	return d.rawQuery(txid, req)
}

func (d *DynDB) RefLoad(txid uint32, tenantId, gslug string, req *store.RefLoadReq) (*store.QueryResult, error) {
	return d.refLoad(txid, tenantId, gslug, req)
}

func (d *DynDB) ReverseRefLoad(txid uint32, tenantId, gslug string, req *store.RevRefLoadReq) (*store.QueryResult, error) {
	return d.reverseRefLoad(txid, tenantId, gslug, req)
}

func (d *DynDB) GetCache() store.DCache {

	return nil
}
