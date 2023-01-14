package coredb

import (
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

var resFields = []string{"name", "type", "sub_type", "owner", "schema", "policy", "plug_id", "invoker", "target", "extra_meta"}

func (d *DB) ResourceNew(tenantId string, obj *entities.Resource) error {
	_, err := d.resTable().Insert(obj)
	return err
}

func (d *DB) ResourceUpdate(tenantId string, id string, data map[string]interface{}) error {
	if !only(data, resFields...) {
		return easyerr.Error("not allowed filed")
	}

	return d.resTable().Find(db.Cond{"tenant_id": tenantId, "id": id}).Update(data)
}

func (d *DB) ResourceGet(tenantId, rid string) (*entities.Resource, error) {
	res := &entities.Resource{}

	err := d.resTable().Find(db.Cond{"tenant_id": tenantId, "id": rid}).One(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *DB) ResourceDel(tenantId, rid string) error {
	return d.resTable().Find(db.Cond{"tenant_id": tenantId, "id": rid}).Delete()
}

func (d *DB) ResourceList(tenantId string) ([]*entities.Resource, error) {
	ress := make([]*entities.Resource, 0)

	err := d.resTable().Find(db.Cond{"tenant_id": tenantId}).All(&ress)
	if err != nil {
		return nil, err
	}

	return ress, nil
}

func (d *DB) ResourcesMulti(tenantId string, rids ...string) ([]*entities.Resource, error) {
	ress := make([]*entities.Resource, 0)

	err := d.resTable().Find(db.Cond{"tenant_id": tenantId, "id IN": rids}).All(&ress)
	if err != nil {
		return nil, err
	}

	return ress, nil

}

func (d *DB) ResourcesByTarget(tenantId string, target string) ([]*entities.Resource, error) {
	ress := make([]*entities.Resource, 0)

	err := d.resTable().Find(db.Cond{"tenant_id": tenantId, "target": target}).All(&ress)
	if err != nil {
		return nil, err
	}

	return ress, nil
}

// private

func (d *DB) resTable() db.Collection {
	return dbutils.Table(d.session, "resources")
}
