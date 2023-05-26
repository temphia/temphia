package dynddl2

import (
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/upper/db/v4"
)

type DynDDL struct {
	session    db.Session
	sharedLock service.DyndbLock
	dyngen     ucore.Zenerator
	logger     zerolog.Logger
}

func (d *DynDDL) RunNew(tenantId string, migctx MigrateContext) error {
	return d.newGroup(tenantId, migctx.StmtString, migctx.BaseSchema)
}

func (d *DynDDL) newGroup(tenantId, stmt string, model *xbprint.NewTableGroup) error {

	// utok, err := d.sharedLock.GlobalLock(tenantId)
	// if err != nil {
	// 	d.logger.Err(err).Send()
	// 	return err
	// }

	// defer d.sharedLock.GlobalUnLock(tenantId, utok)

	// err = d.newGroupMeta(tenantId, model)
	// if err != nil {
	// 	d.logger.
	// 		Err(err).
	// 		Interface("model", model).
	// 		Str("stmt", stmt).
	// 		Caller().
	// 		Msg(NewGroupMetadataCreateErr)

	// 	return err
	// }

	// d.logger.Info().Msg(NewGroupMetadataCreated)
	// err = dbutils.Execute(ucore.GetDriver(d.session), stmt)
	// if err != nil {
	// 	d.logger.
	// 		Err(err).
	// 		Interface("model", model).
	// 		Str("stmt", stmt).
	// 		Caller().
	// 		Msg(NewGroupSchemaExecErr)

	// 	d.rollbackGroupRef(tenantId, model.Slug)
	// }

	return nil

}

func (d *DynDDL) RunUpdate(tenantId string, migctx MigrateContext) error {

	return nil
}
