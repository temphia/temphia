package dyndb

import (
	"fmt"
	"regexp"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/libx/dbutils/hsql"
	"github.com/temphia/temphia/code/backend/stores/upperdb/dyndb/dyncore"
	"github.com/temphia/temphia/code/backend/stores/upperdb/dyndb/tns"
	"github.com/temphia/temphia/code/backend/stores/upperdb/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/xpacman/xpackage"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/upper/db/v4"
)

var _ dyndb.DynDB = (*DynDB)(nil)

type DynDB struct {
	session    db.Session
	sharedLock service.DyndbLock
	dyngen     ucore.Zenerator
	txn        dbutils.TxManager
	tns        tns.TNS
	vendor     string
	cache      dyndb.DCache
	hsql       *hsql.Hsql

	loggerBuilder func() zerolog.Logger
}

func New(opts ucore.DynDBOptions) *DynDB {

	d := &DynDB{
		session:       opts.Session,
		sharedLock:    opts.SharedLock,
		txn:           opts.TxnManager,
		dyngen:        opts.DynGen,
		tns:           opts.TNS,
		vendor:        opts.Vendor,
		cache:         nil,
		hsql:          hsql.New(opts.TNS),
		loggerBuilder: opts.LoggerBuilder,
	}

	d.cache = dyncore.NewCache(d.ListColumns)
	return d
}

func (d *DynDB) NewRow(txid uint32, req dyndb.NewRowReq) (int64, error) {
	return d.newRow(txid, req)
}

func (d *DynDB) GetRow(txid uint32, req dyndb.GetRowReq) (map[string]interface{}, error) {
	return d.getRow(txid, req)
}

func (d *DynDB) UpdateRow(txid uint32, req dyndb.UpdateRowReq) (map[string]interface{}, error) {
	return d.updateRow(txid, req)
}

func (d *DynDB) DeleteRowBatch(txid uint32, req dyndb.DeleteRowBatchReq) ([]int64, error) {
	return d.deleteRowBatch(txid, req)
}

func (d *DynDB) DeleteRowMulti(txid uint32, req dyndb.DeleteRowMultiReq) error {
	return d.deleteRowMulti(txid, req)
}

func (d *DynDB) DeleteRow(txid uint32, req dyndb.DeleteRowReq) error {
	return d.deleteRow(txid, req)
}

func (d *DynDB) SimpleQuery(txid uint32, req dyndb.SimpleQueryReq) (*dyndb.QueryResult, error) {
	return d.simpleQuery(txid, req)
}

func (d *DynDB) FTSQuery(txid uint32, req dyndb.FTSQueryReq) (*dyndb.QueryResult, error) {
	return d.ftsQuery(txid, req)
}

func (d *DynDB) JoinQuery(txid uint32, req dyndb.JoinReq) (*dyndb.JoinResult, error) {
	return d.joinQuery(txid, req)
}

func (d *DynDB) MultiJoinQuery(txid uint32, req dyndb.MultiJoinReq) (*dyndb.MultiJoinResult, error) {
	return d.multiJoinQuery(txid, req)
}

func (d *DynDB) TemplateQuery(txid uint32, req dyndb.TemplateQueryReq) (*dyndb.QueryResult, error) {
	return d.templateQuery(txid, req)
}

func (d *DynDB) MigrateSchema(tenantId string, opts xpackage.MigrateOptions) error {
	return d.migrateSchema(tenantId, opts)
}

func (d *DynDB) SqlQueryRaw(txid uint32, tenantId, group, qstr string) (any, error) {
	return d.sqlQueryRaw(txid, tenantId, group, qstr)
}

func (d *DynDB) SqlQueryScopped(txid uint32, tenantId, group, qstr string) (any, error) {
	return d.sqlQueryScopped(txid, tenantId, group, qstr)
}

func (d *DynDB) RefLoad(txid uint32, tenantId, gslug string, req *dyndb.RefLoadReq) (*dyndb.QueryResult, error) {
	return d.refLoad(txid, tenantId, gslug, req)
}

func (d *DynDB) ReverseRefLoad(txid uint32, tenantId, gslug string, req *dyndb.RevRefLoadReq) (*dyndb.QueryResult, error) {
	return d.reverseRefLoad(txid, tenantId, gslug, req)
}

func (d *DynDB) GetCache() dyndb.DCache {
	return d.cache
}

func (d *DynDB) GetDriver() any {
	return d.session.Driver()
}

// private

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

func removeSQLComments(query string) string {
	commentRegex := regexp.MustCompile(`(--|#|/\*).*?(\*/|\n)`)
	return commentRegex.ReplaceAllString(query, "")
}

func extractTableName(input string) (string, error) {
	regex := regexp.MustCompile(`'([^']*)'`)

	match := regex.FindStringSubmatch(input)

	if len(match) != 2 {
		return "", fmt.Errorf("Unable to extract table name from input string: %s", input)
	}

	return match[1], nil
}
