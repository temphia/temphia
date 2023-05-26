package migreflect

import (
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"

	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

const (
	NewGroupRollBack          = "new_group_rollback"
	GlobalLockErr             = "global_lock_err"
	NewGroupMetadataCreated   = "new_group_metadata_created"
	NewGroupMetadataCreateErr = "new_group_metadata_created_err"
	NewGroupSchemaExecErr     = "new_group_schema_exec_err"
	ColumnsCleanupErr         = "columns_cleanup_err"
	TablesCleanupErr          = "tables_cleanup_err"
	TableCleanupErr           = "tables_cleanup_err"
	GroupCleanupErr           = "group_cleanup_err"
)

func (d *MigReflect) NewGroup(tenantId, stmt string, model *xbprint.NewTableGroup) error {

	utok, err := d.sharedLock.GlobalLock(tenantId)
	if err != nil {
		d.logger.Err(err).Send()
		return err
	}

	defer d.sharedLock.GlobalUnLock(tenantId, utok)

	err = d.newGroupRef(tenantId, model)
	if err != nil {
		d.logger.
			Err(err).
			Interface("model", model).
			Str("stmt", stmt).
			Caller().
			Msg(NewGroupMetadataCreateErr)

		return err
	}

	d.logger.Info().Msg(NewGroupMetadataCreated)
	err = dbutils.Execute(ucore.GetDriver(d.session), stmt)
	if err != nil {
		d.logger.
			Err(err).
			Interface("model", model).
			Str("stmt", stmt).
			Caller().
			Msg(NewGroupSchemaExecErr)

		d.rollbackGroupRef(tenantId, model.Slug)
	}

	return err
}

func (d *MigReflect) newGroupRef(tenantId string, model *xbprint.NewTableGroup) (err error) {

	clear := false
	defer func() {
		if clear {
			d.rollbackGroupRef(tenantId, model.Slug)
		}
	}()

	_, err = d.dataTableGroups().Insert(model.To(tenantId))
	if err != nil {
		return
	}

	clear = true

	for _, tbl := range model.Tables {
		err = d.AddTableRef(tenantId, model.Slug, tbl)
		if err != nil {
			pp.Println(err)
			return err
		}
	}
	clear = false

	return
}

func (d *MigReflect) rollbackGroupRef(tenantId string, gslug string) {

	err := d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
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
	}).Delete()
	if err != nil {
		d.logger.Err(err).
			Str("gslug", gslug).
			Caller().
			Msg(TablesCleanupErr)
	}

	err = d.dataTableGroups().Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      gslug,
	}).Delete()
	if err != nil {
		d.logger.Err(err).
			Str("gslug", gslug).
			Caller().
			Msg(GroupCleanupErr)
	}

}
