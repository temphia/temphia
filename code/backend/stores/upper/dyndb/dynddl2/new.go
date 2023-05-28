package dynddl2

import (
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/dyndb/dyncore"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/upper/db/v4"
)

func (d *DynDDL) runNew(tenantId string, migctx MigrateContext) error {
	err := d.newGroup(tenantId, migctx.StmtString, migctx.BaseSchema)
	if err != nil {
		return err
	}

	err = dyncore.GroupTable(d.session).Find(db.Cond{
		"tenant_id": tenantId,
		"slug":      migctx.BaseSchema.Slug,
	}).Update(db.Cond{
		"bprint_id":          migctx.Options.BprintId,
		"bprint_item_id":     migctx.Options.BprintItemId,
		"bprint_instance_id": migctx.Options.BprintInstanceId,
		"bprint_step_head":   migctx.NextMigHead,
		"active":             true,
	})

	if err != nil {
		d.logger.
			Err(err).
			Interface("migctx", migctx).
			Msg(logid.DyndbSetMigHeadErr)
	}

	return nil

}

func (d *DynDDL) newGroup(tenantId, stmt string, model *xbprint.NewTableGroup) error {

	utok, err := d.sharedLock.GlobalLock(tenantId)
	if err != nil {
		d.logger.Err(err).Msg(logid.DyndbGlobalLockErr)
		return err
	}

	defer d.sharedLock.GlobalUnLock(tenantId, utok)

	err = d.MetaNewGroup(tenantId, model)
	if err != nil {
		d.logger.
			Err(err).
			Interface("model", model).
			Str("stmt", stmt).
			Caller().
			Msg(logid.DyndbNewGroupMetadataCreateErr)

		return err
	}

	d.logger.Info().Msg(logid.DyndbNewGroupMetadataCreated)
	err = dbutils.Execute(ucore.GetDriver(d.session), stmt)
	if err != nil {
		d.logger.
			Err(err).
			Interface("model", model).
			Str("stmt", stmt).
			Caller().
			Msg(logid.DyndbNewGroupSchemaExecErr)

		d.MetaRollbackGroup(tenantId, model.Slug)
	}

	return nil

}
