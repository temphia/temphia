package dyndb

import (
	"strings"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/dynddl"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"

	"github.com/upper/db/v4"
)

func (d *DynDB) ddlSess() *dynddl.DynDDL {
	return dynddl.New(d.session, d.sharedLock, d.dyngen)
}

func (d *DynDB) AddTable(tenantId, gslug string, model *xbprint.NewTable) error {
	ddl := d.ddlSess()
	return ddl.AddTable(tenantId, gslug, model)
}

func (d *DynDB) EditTable(tenantId, gslug, tslug string, model *entities.TablePartial) error {

	return d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"slug":      tslug,
	}).Update(model)
}

func (d *DynDB) GetTable(tenantId, gslug, tslug string) (*entities.Table, error) {
	model := &entities.Table{}

	err := d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"slug":      tslug,
	}).One(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (d *DynDB) ListTables(tenantId, gslug string) ([]*entities.Table, error) {
	ts := make([]*entities.Table, 0)
	err := d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).All(&ts)
	return ts, err
}

func (d *DynDB) DeleteTable(tenantId, gslug, tslug string) error {
	ddl := d.ddlSess()
	return ddl.DeleteTable(tenantId, gslug, tslug)
}

// columns stuff

func (d *DynDB) AddColumn(tenantId, gslug, tslug string, model *xbprint.NewColumn) error {
	ddl := d.ddlSess()
	return ddl.AddColumn(tenantId, gslug, tslug, model)
}

func (d *DynDB) GetColumn(tenantId, gslug, tslug, cslug string) (*entities.Column, error) {
	resp := &entities.Column{}
	err := d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"slug":      cslug,
	}).One(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DynDB) EditColumn(tenantId, gslug, tslug, cslug string, data *entities.ColumnPartial) error {
	return d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"slug":      cslug,
	}).Update(data)
}

func (d *DynDB) ListColumns(tenantId, gslug, tslug string) ([]*entities.Column, error) {
	cols := make([]*entities.Column, 0)

	pp.Println(tenantId, gslug, tslug)

	err := d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
	}).All(&cols)
	return cols, err
}

func (d *DynDB) ListReverseColumnRef(tenantId, gslug, tslug string) ([]*entities.Column, error) {
	cols := make([]*entities.Column, 0)

	err := d.dataTableColumns().Find(db.Cond{
		"tenant_id":  tenantId,
		"group_id":   gslug,
		"ref_target": tslug,
	}).All(&cols)

	return cols, err
}

func (d *DynDB) DeleteColumn(tenantId, gslug, tslug, cslug string) error {
	ddl := d.ddlSess()
	return ddl.DeleteColumn(tenantId, gslug, tslug, cslug)
}

func (d *DynDB) ListColumnRef(tenantId, gslug, tslug string) ([]*entities.ColumnFKRef, error) {
	ddl := d.ddlSess()
	return ddl.ListColumnRef(tenantId, gslug, tslug)
}

// group

func (d *DynDB) NewGroup(tenantId string, model *xbprint.NewTableGroup) error {
	_, err := d.GetGroup(tenantId, model.Slug)
	if err == nil {
		return easyerr.Error("Group of same slug exists already")
	}
	ddl := d.ddlSess()
	return ddl.NewGroup(tenantId, model)
}

func (d *DynDB) EditGroup(tenantId string, gslug string, model *entities.TableGroupPartial) error {
	return d.dataTableGroups().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      gslug,
	}).Update(model)
}

func (d *DynDB) GetGroup(tenantId, gslug string) (*entities.TableGroup, error) {
	tg := &entities.TableGroup{}

	err := d.dataTableGroups().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      gslug,
	}).One(tg)

	return tg, err
}

func (d *DynDB) ListGroup(tenantId string) ([]*entities.TableGroup, error) {
	tgs := make([]*entities.TableGroup, 0, 10)
	err := d.dataTableGroups().Find("tenant_id", tenantId).All(&tgs)
	return tgs, err
}

func (d *DynDB) DeleteGroup(tenantId, gslug string) error {

	// soft delete vs hard delete

	datas := make([]*entities.Table, 0)
	d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).All(&datas)

	var buf strings.Builder
	for _, data := range datas {
		tstr, err := d.dyngen.DropTable(tenantId, gslug, data.Slug)
		if err != nil {
			continue
		}
		buf.Write([]byte(tstr))
	}

	d.viewTable().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).Delete()

	d.hookTable().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).Delete()

	dbutils.Execute(ucore.GetDriver(d.session), buf.String())

	d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).Delete()

	d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).Delete()

	return d.dataTableGroups().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      gslug,
	}).Delete()

}

func (d *DynDB) AddIndex(tenantId, gslug, tslug string, model *entities.Index) error {
	ddl := d.ddlSess()
	return ddl.AddIndex(tenantId, gslug, tslug, model)
}

func (d *DynDB) GroupExecute(tenantId, gslug, qstr string) error {
	utok, err := d.sharedLock.GlobalLock(tenantId)
	if err != nil {
		return err
	}
	defer d.sharedLock.GlobalUnLock(tenantId, utok)
	return dbutils.Execute(ucore.GetDriver(d.session), qstr)
}

func (d *DynDB) AddUniqueIndex(tenantId, gslug, tslug string, model *entities.Index) error {
	ddl := d.ddlSess()
	return ddl.AddUniqueIndex(tenantId, gslug, tslug, model)
}

func (d *DynDB) AddFTSIndex(tenantId, gslug, tslug string, model *entities.FTSIndex) error {
	ddl := d.ddlSess()
	return ddl.AddFTSIndex(tenantId, gslug, tslug, model)
}

func (d *DynDB) AddColumnFRef(tenantId, gslug, tslug string, model *entities.ColumnFKRef) error {
	ddl := d.ddlSess()
	return ddl.AddColumnFRef(tenantId, gslug, tslug, model)
}

func (d *DynDB) ListIndex(tenantId, gslug, tslug string) ([]*entities.Index, error) {
	ddl := d.ddlSess()
	return ddl.ListIndex(tenantId, gslug, tslug)
}

func (d *DynDB) ListFKRef(tenantId, gslug, tslug string) ([]*entities.Index, error) {
	ddl := d.ddlSess()
	return ddl.ListFKRef(tenantId, gslug, tslug)
}

func (d *DynDB) RemoveIndex(tenantId, gslug, tslug, slug string) error {
	ddl := d.ddlSess()
	return ddl.RemoveIndex(tenantId, gslug, tslug, slug)
}

func (d *DynDB) RemoveFKRef(tenantId, gslug, tslug, refslug string) error {
	ddl := d.ddlSess()
	return ddl.RemoveFKRef(tenantId, gslug, tslug, refslug)
}
