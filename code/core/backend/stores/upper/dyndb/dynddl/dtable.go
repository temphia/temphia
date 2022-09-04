package dynddl

import (
	"log"

	"github.com/temphia/temphia/code/core/backend/libx/dbutils"
	"github.com/temphia/temphia/code/core/backend/stores/upper/ucore"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/bprints"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"

	"github.com/upper/db/v4"
)

func (d *DynDDL) AddTable(tenantId, gslug string, model *bprints.NewTable) error {
	ddlstr, err := d.dyngen.NewTable(tenantId, gslug, model, []string{"fixme"})
	if err != nil {
		return err
	}

	utok, err := d.sharedLock.GlobalLock(tenantId)
	if err != nil {
		return err
	}
	defer d.sharedLock.GlobalUnLock(tenantId, utok)

	err = d.AddTableRef(tenantId, gslug, model)
	if err != nil {
		return err
	}

	err = dbutils.Execute(ucore.GetDriver(d.session), ddlstr.String())
	if err != nil {
		d.rollbackTableMeta(tenantId, gslug, model.Slug)
	}

	return err
}

func (d *DynDDL) ListTables(tenantId, gslug string) ([]*entities.Table, error) {
	ts := make([]*entities.Table, 0)
	err := d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).All(&ts)
	return ts, err
}

func (d *DynDDL) DeleteTable(tenantId, gslug, tslug string) error {
	d.dataTableColumns().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
	}).Delete()

	return d.dataTables().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"slug":      tslug,
	}).Delete()
}

func (d *DynDDL) AddTableRef(tenantId, gslug string, model *bprints.NewTable) (err error) {
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

	columns := store.ExtractColumns(model, tenantId, gslug)
	for _, col := range columns {
		err = d.AddColumnMeta(tenantId, gslug, model.Slug, col)
		if err != nil {
			return
		}
	}

	clear = false

	return
}

func (d *DynDDL) rollbackTableMeta(tenantId, gslug, tslug string) error {
	log.Println("ROLLING BACK TABLE....", tenantId, gslug, tslug)
	return d.DeleteTable(tenantId, gslug, tslug)
}
