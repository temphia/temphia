package dyndb

import (
	"fmt"

	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

	"github.com/upper/db/v4"
)

func (d *DynDB) activityTable(tenantId, group, table string) db.Collection {
	dtable := d.tns.Table(tenantId, group, table)
	return d.session.Collection(fmt.Sprintf("dact_%s", dtable))
}

func (d *DynDB) QueryActivity(tenantId, group, table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error) {
	//	tbl := d.activityTable(tenantId, group, table)

	return nil, easyerr.NotImpl()
}

func (d *DynDB) ListActivity(tenantId, group, table string, rowId int) ([]*entities.DynActivity, error) {
	tbl := d.activityTable(tenantId, group, table)

	data := make([]*entities.DynActivity, 0)

	err := tbl.Find(db.Cond{
		"row_id": rowId,
	}).All(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DynDB) NewActivity(tenantId, group, table string, record *entities.DynActivity) (int64, error) {

	tbl := d.activityTable(tenantId, group, table)

	rid, err := tbl.Insert(record)
	if err != nil {
		return 0, err
	}

	return rid.ID().(int64), nil
}
