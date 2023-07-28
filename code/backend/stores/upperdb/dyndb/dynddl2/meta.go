package dynddl2

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/stores/upperdb/dyndb/dyncore"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/upper/db/v4"
)

func (d *DynDDL) MetaNewGroup(tenantId string, model *xbprint.NewTableGroup) (err error) {

	clear := false
	defer func() {
		if clear {
			d.MetaRollbackGroup(tenantId, model.Slug)
		}
	}()

	_, err = d.dataTableGroups().Insert(model.To(tenantId))
	if err != nil {
		return
	}

	clear = true

	for _, tbl := range model.Tables {
		err = d.MetaNewTable(tenantId, model.Slug, tbl)
		if err != nil {
			pp.Println(err)
			return err
		}
	}
	clear = false

	return
}

func (d *DynDDL) MetaRollbackGroup(tenantId, gslug string) {

	err := d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).Delete()
	if err != nil {
		d.logger.Err(err).
			Str("gslug", gslug).
			Caller().
			Msg(logid.DyndbColumnsCleanupErr)
	}

	err = d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).Delete()
	if err != nil {
		d.logger.Err(err).
			Str("gslug", gslug).
			Caller().
			Msg(logid.DyndbTablesCleanupErr)
	}

	err = d.dataTableGroups().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      gslug,
	}).Delete()
	if err != nil {
		d.logger.Err(err).
			Str("gslug", gslug).
			Caller().
			Msg(logid.DyndbGroupCleanupErr)
	}
}

func (d *DynDDL) MetaNewTable(tenantId, gslug string, model *xbprint.NewTable) error {
	_, err := d.dataTables().Insert(model.To(tenantId, gslug))
	if err != nil {
		return err
	}

	columns := dyndb.ExtractColumns(model, tenantId, gslug)
	for _, col := range columns {
		err = d.MetaNewColumn(tenantId, gslug, model.Slug, col)
		if err != nil {

			d.dataTableColumns().Find(db.Cond{
				"tenant_id": tenantId,
				"group_id":  gslug,
				"table_id":  model.Slug,
			}).Delete()

			return err
		}
	}

	return nil
}

func (d *DynDDL) MetaRollbackTable(tenantId, gslug, tslug string) {

	d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"slug":      tslug,
	}).Delete()

	d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
	}).Delete()
}

func (d *DynDDL) MetaRollbackColumn(tenantId, gslug, tslug, cslug string) {
	d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"slug":      cslug,
	}).Delete()
}

func (d *DynDDL) MetaNewColumn(tenantId, gslug, tslug string, model *entities.Column) (err error) {
	_, err = d.dataTableColumns().Insert(model)
	return
}

func (d *DynDDL) dataTableGroups() db.Collection {
	return dyncore.GroupTable(d.session)
}

func (d *DynDDL) dataTables() db.Collection {
	return dyncore.Table(d.session)
}

func (d *DynDDL) dataTableColumns() db.Collection {
	return dyncore.TableColumn(d.session)
}
