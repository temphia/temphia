package coredb

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/upper/db/v4"
)

// systemevent

func (d *DB) AddSystemEvent(data *entities.SystemEvent) error {
	_, err := d.systemEventTable().Insert(data)
	return err
}

func (d *DB) RemoveSystemEvent(id int64) error {
	return d.systemEventTable().Find(db.Cond{"id": id}).Delete()
}

func (d *DB) ListSystemEvent(last int64) ([]*entities.SystemEvent, error) {
	resp := make([]*entities.SystemEvent, 0)

	cond := db.Cond{
		"id >": last,
	}

	err := d.systemEventTable().Find(cond).All(&resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DB) QuerySystemEvent(query store.EventQuery) ([]*entities.SystemEvent, error) {

	resp := make([]*entities.SystemEvent, 0)

	cond := db.Cond{
		"tenant_id": query.TenantId,
		"id >":      query.Last,
	}

	if query.Etype != "" {
		cond["etype"] = query.Etype
	}

	err := d.systemEventTable().Find(cond).All(&resp)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

// systemkv

func (d *DB) AddSystemKV(tenantId string, data *entities.SystemKV) error {
	_, err := d.systemKVTable().Insert(data)
	return err

}

func (d *DB) UpdateSystemKV(tenantId, key, ktype string, data map[string]any) error {

	cond := db.Cond{"key": key, "tenant_id": tenantId}

	if ktype != "" {
		cond["type"] = ktype
	}

	return d.systemKVTable().Find(cond).Update(data)
}

func (d *DB) GetSystemKV(tenantId, key, ktype string) (*entities.SystemKV, error) {

	data := &entities.SystemKV{}

	cond := db.Cond{"key": key, "tenant_id": tenantId}

	if ktype != "" {
		cond["type"] = ktype
	}

	err := d.systemKVTable().Find(cond).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) RemoveSystemKV(tenantId, key, ktype string) error {

	cond := db.Cond{"key": key, "tenant_id": tenantId}

	if ktype != "" {
		cond["type"] = ktype
	}

	return d.systemKVTable().Find(cond).Delete()
}

func (d *DB) ListSystemKV(tenantId, ktype, prefix string, last int64) ([]*entities.SystemKV, error) {

	resp := make([]*entities.SystemKV, 0)

	cond := db.Cond{"tenant_id": tenantId}
	if ktype != "" {
		cond["type"] = ktype
	}

	if prefix != "" {
		cond["key LIKE"] = prefix + "%"
	}

	if last != 0 {
		cond["id >"] = last
	}

	err := d.systemKVTable().Find(cond).OrderBy("id").All(&resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// private

func (d *DB) systemKVTable() db.Collection {
	return d.table("system_kv")
}

func (d *DB) systemEventTable() db.Collection {
	return d.table("system_events")
}
