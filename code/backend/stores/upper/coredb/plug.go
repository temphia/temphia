package coredb

import (
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

func (d *DB) PlugNew(tenantId string, pg *entities.Plug) error {
	_, err := d.plugTable().Insert(pg)
	return err
}

var plugField = []string{"name", "executor", "live", "dev", "owner", "bprint_id", "handlers", "extra_meta"}

func (d *DB) PlugUpdate(tenantId string, id string, data map[string]interface{}) error {

	if !only(data, plugField...) {
		return easyerr.Error("not allowed filed")
	}

	return d.plugTable().Find(db.Cond{"id": id, "tenant_id": tenantId}).Update(data)
}

func (d *DB) PlugGet(tenantId, pid string) (*entities.Plug, error) {

	data := &entities.Plug{}

	err := d.plugTable().Find(db.Cond{"id": pid, "tenant_id": tenantId}).One(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DB) PlugDel(tenantId, pid string) error {
	d.agentTable().Find(
		db.Cond{
			"tenant_id": tenantId,
			"plug_id":   pid,
		}).Delete()

	d.resTable().Find(
		db.Cond{
			"tenant_id": tenantId,
			"plug_id":   pid,
		}).Delete()

	return d.plugTable().Find(db.Cond{"id": pid, "tenant_id": tenantId}).Delete()
}

func (d *DB) PlugList(tenantId string) ([]*entities.Plug, error) {
	datas := make([]*entities.Plug, 0)

	err := d.plugTable().Find(db.Cond{"tenant_id": tenantId}).All(&datas)
	if err != nil {
		return nil, nil
	}

	return datas, nil
}

func (d *DB) PlugListByBprint(tenantId, bprint string) ([]*entities.Plug, error) {
	datas := make([]*entities.Plug, 0)

	err := d.plugTable().Find(db.Cond{
		"tenant_id": tenantId,
		"bprint_id": bprint,
	}).All(&datas)
	if err != nil {
		return nil, nil
	}

	return datas, nil
}

// agent

func (d *DB) AgentNew(tenantId string, data *entities.Agent) error {
	_, err := d.agentTable().Insert(data)
	return err
}

var agentFields = []string{
	"name",
	"type",
	"executor",
	"iface_file",
	"entry_file",
	"web_entry",
	"web_script",
	"web_style",
	"web_loader",
	"web_files",
	"extra_meta",
}

func (d *DB) AgentUpdate(tenantId, pid, id string, data map[string]interface{}) error {
	if !only(data, agentFields...) {
		return easyerr.Error("not allowed filed")
	}
	return d.agentTable().Find(db.Cond{"tenant_id": tenantId, "plug_id": pid, "id": id}).Update(data)
}

func (d *DB) AgentGet(tenantId, pid, id string) (*entities.Agent, error) {
	data := &entities.Agent{}

	err := d.agentTable().Find(db.Cond{"id": id, "tenant_id": tenantId, "plug_id": pid}).One(data)
	if err != nil {
		return nil, err
	}
	return data, nil

}
func (d *DB) AgentDel(tenantId, pid, agentId string) error {
	return d.agentTable().Find(db.Cond{"id": pid, "tenant_id": tenantId, "plug_id": pid}).Delete()
}

func (d *DB) AgentList(tenantId, pid string) ([]*entities.Agent, error) {
	datas := make([]*entities.Agent, 0)
	err := d.agentTable().Find(db.Cond{"tenant_id": tenantId, "plug_id": pid}).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

// private

func (d *DB) plugTable() db.Collection {
	return dbutils.Table(d.session, "plugs")
}

func (d *DB) agentTable() db.Collection {
	return dbutils.Table(d.session, "agents")
}
