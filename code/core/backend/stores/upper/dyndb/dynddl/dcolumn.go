package dynddl

import (
	"log"

	"github.com/temphia/temphia/code/core/backend/libx/dbutils"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DynDDL) AddColumn(tenantId, gslug, tslug string, model *bprints.NewColumn) error {
	qstr, err := d.dyngen.AddColumn(tenantId, gslug, tslug, model.Slug, model)
	if err != nil {
		return err
	}

	utok, err := d.sharedLock.GroupLock(tenantId, gslug)
	if err != nil {
		return err
	}

	defer d.sharedLock.GroupUnLock(tenantId, gslug, utok)

	err = d.AddColumnMeta(tenantId, gslug, tslug, model.To(tenantId, gslug, tslug))
	if err != nil {
		return err
	}

	err = dbutils.Execute(ucore.GetDriver(d.session), qstr)
	if err != nil {
		d.rollbackColumnMeta(tenantId, gslug, tslug, model.Slug)
	}
	return err
}

func (d *DynDDL) DeleteColumn(tenantId, gslug, tslug, cslug string) error {
	// fixme => actual column del not just meta table item
	return d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"slug":      cslug,
	}).Delete()
}

func (d *DynDDL) AddColumnMeta(tenantId, gslug, tslug string, model *entities.Column) (err error) {
	_, err = d.dataTableColumns().Insert(model)
	return
}

func (d *DynDDL) rollbackColumnMeta(tenantId, gslug, tslug, cslug string) error {
	log.Println("ROLLING BACK COLUMN....", tenantId, gslug, cslug)
	return d.DeleteColumn(tenantId, gslug, tslug, cslug)
}

func (d *DynDDL) ListColumnRef(tenantId, gslug, tslug string) ([]*entities.ColumnFKRef, error) {
	return nil, easyerr.NotImpl()
}
