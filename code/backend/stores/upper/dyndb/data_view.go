package dyndb

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DynDB) NewView(tenantId string, model *entities.DataView) error {
	_, err := d.viewTable().Insert(model)
	return err
}

// var allowedViewFields = []string{"count", "filter_conds", "selects", "main_column", "search_term", "ascending"}

func (d *DynDB) ModifyView(tenantId, gslug, tslug string, id int64, data map[string]interface{}) error {
	// fixme => disallow certain fields

	// fc, ok := data["filter_conds"]
	// if ok {
	// 	data["filter_conds"] = entities.FilterConds(fc.([]entities.FilterCond))
	// }

	// fixme => impl filter_conds

	return d.viewTable().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"id":        id,
	}).Update(data)
}

func (d *DynDB) ListView(tenantId, gslug, tslug string) ([]*entities.DataView, error) {
	views := make([]*entities.DataView, 0)
	err := d.viewTable().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
	}).All(&views)
	if err != nil {
		return nil, err
	}
	return views, nil
}

func (d *DynDB) DelView(tenantId, gslug, tslug string, id int64) error {
	return d.viewTable().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"id":        id,
	}).Delete()
}

func (d *DynDB) GetView(tenantId, gslug, tslug string, id int64) (*entities.DataView, error) {
	resp := &entities.DataView{}
	err := d.viewTable().Find(db.Cond{
		"tenant_id": tenantId,
		"group_id":  gslug,
		"table_id":  tslug,
		"id":        id,
	}).One(resp)

	pp.Println("@get_view", resp)

	if err != nil {
		return nil, err
	}
	return resp, nil

}
