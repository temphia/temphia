package meta

import (
	"github.com/k0kubun/pp"
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/dyncore"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/temphia/temphia/code/backend/xtypes/store/dyndb"
	"github.com/upper/db/v4"
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

type DynMeta interface {
	NewGroupMeta(tenantId string, model *xbprint.NewTableGroup) error
	RollbackGroupMeta(tenantId, gslug string)

	NewTableMeta(tenantId, gslug string, model *xbprint.NewTable) error
	RollbackTableMeta(tenantId, gslug, tslug string)

	NewColumnMeta(tenantId, gslug, tslug string, model *entities.Column) error
	RollbackColumnMeta(tenantId, gslug, tslug, cslug string)
}

type dynMeta struct {
	session db.Session
	logger  zerolog.Logger
}

func (d *dynMeta) NewGroupMeta(tenantId string, model *xbprint.NewTableGroup) (err error) {

	clear := false
	defer func() {
		if clear {
			d.RollbackGroupMeta(tenantId, model.Slug)
		}
	}()

	_, err = d.dataTableGroups().Insert(model.To(tenantId))
	if err != nil {
		return
	}

	clear = true

	for _, tbl := range model.Tables {
		err = d.NewTableMeta(tenantId, model.Slug, tbl)
		if err != nil {
			pp.Println(err)
			return err
		}
	}
	clear = false

	return
}

func (d *dynMeta) RollbackGroupMeta(tenantId, gslug string) {

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

func (d *dynMeta) NewTableMeta(tenantId, gslug string, model *xbprint.NewTable) error {
	_, err := d.dataTables().Insert(model.To(tenantId, gslug))
	if err != nil {
		return err
	}

	columns := dyndb.ExtractColumns(model, tenantId, gslug)
	for _, col := range columns {
		err = d.NewColumnMeta(tenantId, gslug, model.Slug, col)
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

func (d *dynMeta) RollbackTableMeta(tenantId, gslug, tslug string) {

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

func (d *dynMeta) RollbackColumnMeta(tenantId, gslug, tslug, cslug string) {
	d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"slug":      cslug,
	}).Delete()
}

func (d *dynMeta) NewColumnMeta(tenantId, gslug, tslug string, model *entities.Column) (err error) {
	_, err = d.dataTableColumns().Insert(model)
	return
}

func (d *dynMeta) dataTableGroups() db.Collection {
	return dyncore.GroupTable(d.session)
}

func (d *dynMeta) dataTables() db.Collection {
	return dyncore.Table(d.session)
}

func (d *dynMeta) dataTableColumns() db.Collection {
	return dyncore.TableColumn(d.session)
}
