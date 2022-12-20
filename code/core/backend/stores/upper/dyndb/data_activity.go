package dyndb

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

	"github.com/upper/db/v4"
)

func (d *DynDB) activityTable(tenantId, group, table string) db.Collection {
	return d.session.Collection(d.tns.ActivityTable(tenantId, group, table))
}

func (d *DynDB) QueryActivity(tenantId, group, table string, query *entities.ActivityQuery) ([]*entities.DynActivity, error) {
	tbl := d.activityTable(tenantId, group, table)

	cnd := db.Cond{
		"id >": query.Offset,
	}

	if len(query.Types) != 0 {
		cnd["types IN"] = query.Types
	}

	if query.UserId != "" {
		cnd["user_id"] = query.UserId
	}

	// fixme => impl between

	resp := make([]*entities.DynActivity, 0)
	err := tbl.Find(cnd).Limit(int(query.Count)).OrderBy("id").All(&resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
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
