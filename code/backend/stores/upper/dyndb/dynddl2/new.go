package dynddl2

import (
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

func (d *DynDDL) newGroup(tenantId, stmt string, model *xbprint.NewTableGroup) error {

	utok, err := d.sharedLock.GlobalLock(tenantId)
	if err != nil {
		d.logger.Err(err).Send()
		return err
	}

	defer d.sharedLock.GlobalUnLock(tenantId, utok)

	err = d.meta.NewGroupMeta(tenantId, model)
	if err != nil {
		d.logger.
			Err(err).
			Interface("model", model).
			Str("stmt", stmt).
			Caller().
			Msg(logid.DyndbNewGroupMetadataCreateErr)

		return err
	}

	d.logger.Info().Msg("NewGroupMetadataCreated")
	err = dbutils.Execute(ucore.GetDriver(d.session), stmt)
	if err != nil {
		d.logger.
			Err(err).
			Interface("model", model).
			Str("stmt", stmt).
			Caller().
			Msg(logid.DyndbNewGroupSchemaExecErr)

		d.meta.RollbackGroupMeta(tenantId, model.Slug)
	}

	return nil

}
