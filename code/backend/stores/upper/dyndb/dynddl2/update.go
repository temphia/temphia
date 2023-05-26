package dynddl2

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
)

func (d *DynDDL) update(tenantId string, migctx MigrateContext) error {

	utok, err := d.sharedLock.GlobalLock(tenantId)
	if err != nil {
		d.logger.Err(err).Msg(logid.DyndbGlobalLockErr)
		return err
	}

	defer d.sharedLock.GlobalUnLock(tenantId, utok)

	err = dbutils.Execute(ucore.GetDriver(d.session), migctx.StmtString)
	if err != nil {
		return err
	}

	for idx, pd := range migctx.PostItems {
		pp.Println("@pd", pd)

		switch pd.Mtype {

		case step.MigTypeAddTable:
			err = d.meta.NewTableMeta(
				tenantId,
				migctx.Options.Gslug,
				pd.Data.(*xbprint.NewTable),
			)
			if err != nil {
				d.tryRollbackUpdate(tenantId, idx, migctx)
			}

		case step.MigTypeRemoveTable:
		case step.MigTypeAddColumn:
		case step.MigTypeRemoveColumn:

		}

	}

	return nil
}

func (d *DynDDL) tryRollbackUpdate(tenantId string, currIdx int, migctx MigrateContext) {

	items := migctx.PostItems[:currIdx]

	for i := 0; i < len(items)-1; i = i + 1 {
		item := items[i]

		pp.Println("@item", item)

	}

}
