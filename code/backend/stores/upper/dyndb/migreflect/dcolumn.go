package migreflect

import (
	"log"

	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/xbprint"
	"github.com/upper/db/v4"
)

func (d *MigReflect) AddColumn(tenantId, gslug, tslug, stmt string, model *xbprint.NewColumn) error {

	utok, err := d.sharedLock.GroupLock(tenantId, gslug)
	if err != nil {
		return err
	}

	defer d.sharedLock.GroupUnLock(tenantId, gslug, utok)

	err = d.AddColumnMeta(tenantId, gslug, tslug, model.To(tenantId, gslug, tslug))
	if err != nil {
		return err
	}

	err = dbutils.Execute(ucore.GetDriver(d.session), stmt)
	if err != nil {
		d.rollbackColumnMeta(tenantId, gslug, tslug, model.Slug)
	}
	return err
}

func (d *MigReflect) DeleteColumn(tenantId, gslug, tslug, cslug string) error {
	return d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"slug":      cslug,
	}).Delete()
}

func (d *MigReflect) AddColumnMeta(tenantId, gslug, tslug string, model *entities.Column) (err error) {
	_, err = d.dataTableColumns().Insert(model)
	return
}

func (d *MigReflect) rollbackColumnMeta(tenantId, gslug, tslug, cslug string) error {
	log.Println("ROLLING BACK COLUMN....", tenantId, gslug, cslug)
	return d.DeleteColumn(tenantId, gslug, tslug, cslug)
}
