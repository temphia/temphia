package dynddl2

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/upper/db/v4"
)

func (d *DynDDL) update(tenantId string, migctx MigrateContext) error {
	nextHead := ""

	utok, err := d.sharedLock.GlobalLock(tenantId)
	if err != nil {
		d.logger.Err(err).Msg(logid.DyndbGlobalLockErr)
		return err
	}

	defer func() {
		if nextHead != "" {
			err = d.dataTableGroups().Find(db.Cond{
				"tenant_id": tenantId,
				"slug":      migctx.BaseSchema.Slug,
			}).Update(db.Cond{
				"migration_head": nextHead,
				"active":         true,
			})
		}

		d.sharedLock.GlobalUnLock(tenantId, utok)
	}()

	for _, pd := range migctx.PostItems {
		pp.Println("@pd", pd)

		switch pd.Mtype {

		case step.MigTypeAddTable:

			err = dbutils.Execute(ucore.GetDriver(d.session), pd.Stmt)
			if err != nil {
				return err
			}

			err = d.MetaNewTable(
				tenantId,
				migctx.Options.Gslug,
				pd.Data.(*xbprint.NewTable),
			)
			if err != nil {
				return err
			}

		case step.MigTypeRemoveTable:
			schema := pd.Data.(*xbprint.RemoveTable)
			d.MetaRollbackTable(tenantId, migctx.Options.Gslug, schema.Slug)
			ok, err := d.dataTables().Find(db.Cond{
				"tenant_id": tenantId,
				"slug":      schema.Slug,
			}).Exists()
			if err != nil {
				return err
			}

			if ok {
				return easyerr.Error("could not drop table meta")
			}

			err = dbutils.Execute(ucore.GetDriver(d.session), pd.Stmt)
			if err != nil {
				return err
			}

		case step.MigTypeAddColumn:
			schema := pd.Data.(*xbprint.NewColumn)

			err = dbutils.Execute(ucore.GetDriver(d.session), pd.Stmt)
			if err != nil {
				return err
			}

			err = d.MetaNewColumn(
				tenantId,
				migctx.Options.Gslug,
				schema.Table,
				schema.To(tenantId, migctx.Options.Gslug, schema.Table),
			)

			if err != nil {
				return err
			}

		case step.MigTypeRemoveColumn:
			schema := pd.Data.(*xbprint.RemoveColumn)
			d.MetaRollbackColumn(tenantId, migctx.Options.Gslug, schema.Table, schema.Slug)

			ok, err := d.dataTableColumns().Find(db.Cond{
				"tenant_id": tenantId,
				"table_id":  schema.Table,
				"slug":      schema.Slug,
			}).Exists()

			if err != nil {
				return err
			}

			if ok {
				return easyerr.Error("could not drop column meta")
			}

			err = dbutils.Execute(ucore.GetDriver(d.session), pd.Stmt)
			if err != nil {
				return err
			}
		}

		nextHead = pd.Stmt

	}

	return nil
}
