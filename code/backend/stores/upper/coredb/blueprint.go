package coredb

import (
	"github.com/temphia/temphia/code/core/backend/libx/dbutils"
	"github.com/temphia/temphia/code/core/backend/libx/easyerr"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

var updateAllowed = []string{"slug", "name", "type", "sub_type", "inline_schema", "description", "icon", "source_id", "files", "tags", "extra_meta"}

func (d *DB) BprintNew(tenantId string, et *entities.BPrint) error {
	table := d.blueprintTable()

	_, err := table.Insert(et)
	return err
}

func (d *DB) BprintUpdate(tenantId, id string, data map[string]interface{}) error {
	table := d.blueprintTable()
	if !only(data, updateAllowed...) {
		return easyerr.Error("not allowed filed")
	}

	return table.Find(db.Cond{"id": id, "tenant_id": tenantId}).Update(data)
}

func (d *DB) BprintGet(tenantId, id string) (*entities.BPrint, error) {
	table := d.blueprintTable()

	data := &entities.BPrint{}

	err := table.Find(db.Cond{"id": id, "tenant_id": tenantId}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) BprintDel(tenantId, id string) error {
	table := d.blueprintTable()
	return table.Find(db.Cond{"id": id, "tenant_id": tenantId}).Delete()
}

func (d *DB) BprintList(tenantId, group string) ([]*entities.BPrint, error) {

	data := make([]*entities.BPrint, 0)

	table := d.blueprintTable()
	err := table.Find(db.Cond{"tenant_id": tenantId}).All(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// private

func (d *DB) blueprintTable() db.Collection {
	return dbutils.Table(d.session, "bprints")
}
