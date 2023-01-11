package coredb

import (
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DB) AddTargetHook(data *entities.TargetHook) error {
	_, err := d.targetHookTable().Insert(data)
	return err
}

func (d *DB) UpdateTargetHook(tenantId, ttype string, id int64, data map[string]any) error {
	return d.targetHookTable().Find(db.Cond{
		"id": id, "tenant_id": tenantId,
		"target_type": ttype}).Update(data)
}

func (d *DB) ListTargetHook(tenantId string) ([]*entities.TargetHook, error) {
	ws := make([]*entities.TargetHook, 0)

	err := d.targetHookTable().Find(db.Cond{"tenant_id": tenantId}).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (d *DB) ListTargetHookByType(tenantId, ttype, target string) ([]*entities.TargetHook, error) {
	ws := make([]*entities.TargetHook, 0)

	cond := db.Cond{
		"tenant_id":   tenantId,
		"target_type": ttype,
	}

	if target != "" {
		cond["target"] = target
	}

	err := d.targetHookTable().Find(cond).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (d *DB) GetTargetHook(tenantId, ttype string, id int64) (*entities.TargetHook, error) {
	w := &entities.TargetHook{}
	err := d.targetHookTable().Find(db.Cond{
		"id": id, "tenant_id": tenantId,
		"target_type": ttype,
	}).One(w)

	if err != nil {
		return nil, err
	}
	return w, nil
}
func (d *DB) RemoveTargetHook(tenantId, ttype string, id int64) error {
	return d.targetHookTable().Find(db.Cond{
		"id":          id,
		"tenant_id":   tenantId,
		"target_type": ttype,
	}).Delete()
}

// target app

func (d *DB) AddTargetApp(data *entities.TargetApp) error {
	_, err := d.targetAppTable().Insert(data)
	return err
}

func (d *DB) UpdateTargetApp(tenantId, ttype string, id int64, data map[string]any) error {
	return d.targetAppTable().Find(db.Cond{
		"id": id, "tenant_id": tenantId,
		"target_type": ttype}).Update(data)
}

func (d *DB) ListTargetApp(tenantId string) ([]*entities.TargetApp, error) {
	ws := make([]*entities.TargetApp, 0)

	err := d.targetAppTable().Find(db.Cond{"tenant_id": tenantId}).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (d *DB) ListTargetAppByType(tenantId, ttype, target string) ([]*entities.TargetApp, error) {
	ws := make([]*entities.TargetApp, 0)

	cond := db.Cond{
		"tenant_id":   tenantId,
		"target_type": ttype,
	}

	if target != "" {
		cond["target"] = target
	}

	err := d.targetAppTable().Find(cond).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (d *DB) GetTargetApp(tenantId, ttype string, id int64) (*entities.TargetApp, error) {
	w := &entities.TargetApp{}
	err := d.targetAppTable().Find(db.Cond{
		"id": id, "tenant_id": tenantId,
		"target_type": ttype,
	}).One(w)

	if err != nil {
		return nil, err
	}
	return w, nil
}
func (d *DB) RemoveTargetApp(tenantId, ttype string, id int64) error {
	return d.targetAppTable().Find(db.Cond{
		"id":          id,
		"tenant_id":   tenantId,
		"target_type": ttype,
	}).Delete()
}

// private

func (d *DB) targetHookTable() db.Collection {
	return d.table("target_hooks")
}

func (d *DB) targetAppTable() db.Collection {
	return d.table("target_apps")
}

func (d *DB) ListTargetAppByUgroup(tenantId, ugroup string) ([]*entities.TargetApp, error) {
	ws := make([]*entities.TargetApp, 0)

	err := d.targetAppTable().Find(
		db.Cond{
			"tenant_id":   tenantId,
			"target_type": entities.TargetAppTypeUserGroupApp,
			"target":      ugroup,
		},
	).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (d *DB) ListTargetAppByPlug(tenantId, plug string) ([]*entities.TargetApp, error) {
	ws := make([]*entities.TargetApp, 0)

	err := d.targetAppTable().Find(
		db.Cond{
			"tenant_id": tenantId,
			"plug_id":   plug,
		},
	).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil

}

func (d *DB) ListTargetHookByPlug(tenantId, plug string) ([]*entities.TargetHook, error) {
	ws := make([]*entities.TargetHook, 0)

	err := d.targetHookTable().Find(
		db.Cond{
			"tenant_id": tenantId,
			"plug_id":   plug,
		},
	).All(&ws)
	if err != nil {
		return nil, err
	}

	return ws, nil

}
