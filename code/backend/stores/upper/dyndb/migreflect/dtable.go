package migreflect

import (
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/upper/db/v4"
)

const (
	AddTableRollBack          = "add_table_rollback"
	AddTableMetadataCreated   = "new_group_metadata_created"
	AddTableMetadataCreateErr = "new_group_metadata_created_err"
	AddTableSchemaExecErr     = "new_group_schema_exec_err"
)

func (d *MigReflect) AddTable(tenantId, gslug, ddlstr string, model *xbprint.NewTable) error {

	utok, err := d.sharedLock.GlobalLock(tenantId)
	if err != nil {
		return err
	}
	defer d.sharedLock.GlobalUnLock(tenantId, utok)

	err = d.AddTableRef(tenantId, gslug, model)
	if err != nil {
		d.logger.
			Err(err).
			Interface("model", model).
			Str("stmt", ddlstr).
			Caller().
			Msg(AddTableMetadataCreateErr)

		return err
	}

	d.logger.Info().
		Str("gslug", gslug).
		Msg(AddTableMetadataCreated)

	err = dbutils.Execute(ucore.GetDriver(d.session), ddlstr)
	if err != nil {
		d.logger.
			Err(err).
			Interface("model", model).
			Str("stmt", ddlstr).
			Caller().
			Msg(AddTableSchemaExecErr)

		d.rollbackTableMeta(tenantId, gslug, model.Slug)
	}

	return err
}

func (d *MigReflect) AddTableRef(tenantId, gslug string, model *xbprint.NewTable) (err error) {
	clear := false
	defer func() {
		if clear {
			d.rollbackTableMeta(tenantId, gslug, model.Slug)
		}
	}()

	_, err = d.dataTables().Insert(model.To(tenantId, gslug))
	if err != nil {
		return
	}
	clear = true

	columns := dyndb.ExtractColumns(model, tenantId, gslug)
	for _, col := range columns {
		err = d.AddColumnMeta(tenantId, gslug, model.Slug, col)
		if err != nil {
			return
		}
	}

	clear = false

	return
}

func (d *MigReflect) rollbackTableMeta(tenantId, gslug, tslug string) {

	err := d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
	}).Delete()
	if err != nil {
		d.logger.Err(err).
			Str("gslug", gslug).
			Caller().
			Msg(ColumnsCleanupErr)

	}

	err = d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
	}).Delete()
	if err != nil {
		d.logger.Err(err).
			Str("gslug", gslug).
			Caller().
			Msg(TableCleanupErr)
	}

}
