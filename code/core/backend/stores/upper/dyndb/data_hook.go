package dyndb

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DynDB) NewHook(model *entities.DataHook) error {
	_, err := d.hookTable().Insert(model)
	return err
}

func (d *DynDB) ModifyHook(tenantId, gslug, tslug string, id int64, data map[string]interface{}) error {
	// fixme => disallow certain fields
	return d.hookTable().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"id":        id,
	}).Update(data)
}

func (d *DynDB) ListHook(tenantId, gslug, tslug string) ([]*entities.DataHook, error) {
	hooks := make([]*entities.DataHook, 0)
	err := d.hookTable().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
	}).All(&hooks)
	if err != nil {
		return nil, err
	}
	return hooks, nil
}

func (d *DynDB) GetHook(tenantId, gslug, tslug string, id int64) (*entities.DataHook, error) {
	resp := &entities.DataHook{}
	err := d.hookTable().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
	}).One(resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DynDB) DelHook(tenantId, gslug, tslug string, id int64) error {
	return d.hookTable().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"id":        id,
	}).Delete()
}
