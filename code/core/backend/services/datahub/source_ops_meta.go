package datahub

import (
	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

// group

var _ store.DynSource = (*dynSource)(nil)

type dynSource struct {
	hub      *DataHub
	source   string
	tenantId string
}

func (d *dynSource) dynDB() store.DynDB {

	pp.Println(d.source, d.tenantId)

	return d.hub.dyndbs[d.source]
}

func (d *dynSource) Name() string {
	return d.source
}

func (d *dynSource) NewGroup(model *bprints.NewTableGroup) error {
	ddb := d.dynDB()
	return ddb.NewGroup(d.tenantId, model)
}

func (d *dynSource) EditGroup(gslug string, model *entities.TableGroupPartial) error {
	ddb := d.dynDB()
	return ddb.EditGroup(d.tenantId, gslug, model)
}

func (d *dynSource) ListGroup() ([]*entities.TableGroup, error) {
	ddb := d.dynDB()
	return ddb.ListGroup(d.tenantId)
}

func (d *dynSource) GetGroup(gslug string) (*entities.TableGroup, error) {
	ddb := d.dynDB()
	return ddb.GetGroup(d.tenantId, gslug)
}

func (d *dynSource) DeleteGroup(gslug string) error {
	ddb := d.dynDB()
	return ddb.DeleteGroup(d.tenantId, gslug)
}

// table

func (d *dynSource) AddTable(gslug string, model *bprints.NewTable) error {
	ddb := d.dynDB()
	return ddb.AddTable(d.tenantId, gslug, model)
}
func (d *dynSource) EditTable(gslug, tslug string, model *entities.TablePartial) error {
	ddb := d.dynDB()
	return ddb.EditTable(d.tenantId, gslug, tslug, model)
}
func (d *dynSource) ListTables(gslug string) ([]*entities.Table, error) {
	ddb := d.dynDB()
	return ddb.ListTables(d.tenantId, gslug)
}
func (d *dynSource) DeleteTable(gslug, tslug string) error {
	ddb := d.dynDB()
	return ddb.DeleteTable(d.tenantId, gslug, tslug)
}

func (d *dynSource) GetTable(gslug, tslug string) (*entities.Table, error) {
	ddb := d.dynDB()
	return ddb.GetTable(d.tenantId, gslug, tslug)
}

func (d *dynSource) RefResolve(txid uint32, gslug string, req *store.RefResolveReq) (*store.QueryResult, error) {
	ddb := d.dynDB()
	return ddb.RefResolve(txid, d.tenantId, gslug, req)
}

func (d *dynSource) RefLoad(txid uint32, gslug string, req *store.RefLoadReq) (*store.QueryResult, error) {
	ddb := d.dynDB()

	return ddb.RefLoad(txid, d.tenantId, gslug, req)
}

func (d *dynSource) ReverseRefLoad(txid uint32, gslug string, req *store.RevRefLoadReq) (*store.QueryResult, error) {
	ddb := d.dynDB()
	return ddb.ReverseRefLoad(txid, d.tenantId, gslug, req)
}

// column

func (d *dynSource) AddColumn(gslug, tslug string, model *bprints.NewColumn) error {
	ddb := d.dynDB()
	return ddb.AddColumn(d.tenantId, gslug, tslug, model)
}

func (d *dynSource) GetColumn(gslug, tslug, cslug string) (*entities.Column, error) {
	ddb := d.dynDB()
	return ddb.GetColumn(d.tenantId, gslug, tslug, cslug)
}

func (d *dynSource) EditColumn(gslug, tslug, cslug string, model *entities.ColumnPartial) error {
	ddb := d.dynDB()
	return ddb.EditColumn(d.tenantId, gslug, tslug, cslug, model)
}
func (d *dynSource) ListColumns(gslug, tslug string) ([]*entities.Column, error) {
	ddb := d.dynDB()
	return ddb.ListColumns(d.tenantId, gslug, tslug)
}
func (d *dynSource) DeleteColumn(gslug, tslug, cslug string) error {
	ddb := d.dynDB()
	return ddb.DeleteColumn(d.tenantId, gslug, tslug, cslug)
}

// index

func (d *dynSource) AddIndex(gslug, tslug string, model *entities.Index) error {
	ddb := d.dynDB()
	return ddb.AddIndex(d.tenantId, gslug, tslug, model)
}

func (d *dynSource) AddUniqueIndex(gslug, tslug string, model *entities.Index) error {
	ddb := d.dynDB()
	return ddb.AddUniqueIndex(d.tenantId, gslug, tslug, model)
}

func (d *dynSource) AddFTSIndex(gslug, tslug string, model *entities.FTSIndex) error {
	ddb := d.dynDB()
	return ddb.AddFTSIndex(d.tenantId, gslug, tslug, model)
}
func (d *dynSource) AddColumnFRef(gslug, tslug string, model *entities.ColumnFKRef) error {
	ddb := d.dynDB()
	return ddb.AddColumnFRef(d.tenantId, gslug, tslug, model)
}
func (d *dynSource) ListIndex(gslug, tslug string) ([]*entities.Index, error) {
	ddb := d.dynDB()
	return ddb.ListIndex(d.tenantId, gslug, tslug)
}
func (d *dynSource) ListColumnRef(gslug, tslug string) ([]*entities.ColumnFKRef, error) {
	ddb := d.dynDB()
	return ddb.ListColumnRef(d.tenantId, gslug, tslug)
}
func (d *dynSource) RemoveIndex(gslug, tslug, slug string) error {
	ddb := d.dynDB()
	return ddb.RemoveIndex(d.tenantId, gslug, tslug, slug)
}

// view

func (d *dynSource) NewView(model *entities.DataView) error {
	ddb := d.dynDB()
	return ddb.NewView(model)
}

func (d *dynSource) ModifyView(gslug, tslug string, id int64, data map[string]any) error {
	ddb := d.dynDB()
	return ddb.ModifyView(d.tenantId, gslug, tslug, id, data)
}

func (d *dynSource) ListView(gslug, tslug string) ([]*entities.DataView, error) {
	ddb := d.dynDB()
	return ddb.ListView(d.tenantId, gslug, tslug)
}

func (d *dynSource) DelView(gslug, tslug string, id int64) error {
	ddb := d.dynDB()
	return ddb.DelView(d.tenantId, gslug, tslug, id)
}

func (d *dynSource) GetView(gslug, tslug string, id int64) (*entities.DataView, error) {
	ddb := d.dynDB()
	return ddb.GetView(d.tenantId, gslug, tslug, id)
}

func (d *dynSource) QueryActivity(group, table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error) {
	ddb := d.dynDB()
	return ddb.QueryActivity(d.tenantId, group, table, query)
}

func (d *dynSource) ListActivity(group, table string, rowId int) ([]*entities.DynActivity, error) {
	ddb := d.dynDB()
	return ddb.ListActivity(d.tenantId, group, table, rowId)
}

func (d *dynSource) NewActivity(group, table string, record *entities.DynActivity) error {
	ddb := d.dynDB()

	id, err := ddb.NewActivity(d.tenantId, group, table, record)

	// fixme => here check for mentions

	pp.Println(id)

	return err
}
