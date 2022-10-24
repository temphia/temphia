package coredb

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DB) AddTargetHook(data *entities.TargetHook) error {
	_, err := d.targetHookTable().Insert(data)
	return err
}

func (d *DB) UpdateTargetHook(tenantId, targetType string, id int64, data map[string]any) error {
	return d.targetHookTable().Find(db.Cond{
		"id": id, "tenant_id": tenantId,
		"target_type": targetType}).Update(data)
}

func (d *DB) ListTargetHook(tenantId string) ([]*entities.TargetHook, error) {
	ws := make([]*entities.TargetHook, 0)

	err := d.targetHookTable().Find(db.Cond{"tenant_id": tenantId}).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (d *DB) ListTargetHookByType(tenantId, targetType string) ([]*entities.TargetHook, error) {
	ws := make([]*entities.TargetHook, 0)

	err := d.targetHookTable().Find(db.Cond{"tenant_id": tenantId, "target_type": targetType}).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (d *DB) GetTargetHook(tenantId, targetType string, id int64) (*entities.TargetHook, error) {
	w := &entities.TargetHook{}
	err := d.targetHookTable().Find(db.Cond{
		"id": id, "tenant_id": tenantId,
		"target_type": targetType,
	}).One(w)

	if err != nil {
		return nil, err
	}
	return w, nil
}
func (d *DB) RemoveTargetHook(tenantId, targetType string, id int64) error {
	return d.targetHookTable().Find(db.Cond{
		"id":          id,
		"tenant_id":   tenantId,
		"target_type": targetType,
	}).Delete()
}

// target app

func (d *DB) AddTargetApp(data *entities.TargetApp) error {
	_, err := d.targetAppTable().Insert(data)
	return err
}

func (d *DB) UpdateTargetApp(tenantId, targetType string, id int64, data map[string]any) error {
	return d.targetAppTable().Find(db.Cond{
		"id": id, "tenant_id": tenantId,
		"target_type": targetType}).Update(data)
}

func (d *DB) ListTargetApp(tenantId string) ([]*entities.TargetApp, error) {
	ws := make([]*entities.TargetApp, 0)

	err := d.targetAppTable().Find(db.Cond{"tenant_id": tenantId}).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (d *DB) ListTargetAppByType(tenantId, targetType string) ([]*entities.TargetApp, error) {
	ws := make([]*entities.TargetApp, 0)

	err := d.targetAppTable().Find(db.Cond{"tenant_id": tenantId, "target_type": targetType}).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (d *DB) GetTargetApp(tenantId, targetType string, id int64) (*entities.TargetApp, error) {
	w := &entities.TargetApp{}
	err := d.targetAppTable().Find(db.Cond{
		"id": id, "tenant_id": tenantId,
		"target_type": targetType,
	}).One(w)

	if err != nil {
		return nil, err
	}
	return w, nil
}
func (d *DB) RemoveTargetApp(tenantId, targetType string, id int64) error {
	return d.targetAppTable().Find(db.Cond{
		"id":          id,
		"tenant_id":   tenantId,
		"target_type": targetType,
	}).Delete()
}

// private

func (d *DB) targetHookTable() db.Collection {
	return d.table("target_hooks")
}

func (d *DB) targetAppTable() db.Collection {
	return d.table("target_apps")
}
