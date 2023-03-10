package coredb

import (
	"github.com/temphia/temphia/code/backend/libx/dbutils"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/upper/db/v4"
)

// links

func (d *DB) AgentLinkNew(tenantId string, data *entities.AgentLink) error {
	_, err := d.agentLinkTable().Insert(data)
	return err
}

func (d *DB) AgentLinkUpdate(tenantId, pid, aid string, id int64, data map[string]interface{}) error {
	return d.agentLinkTable().Find(db.Cond{
		"id":            id,
		"tenant_id":     tenantId,
		"from_plug_id":  pid,
		"from_agent_id": aid,
	}).Update(data)
}

func (d *DB) AgentLinkGet(tenantId, pid, aid string, id int64) (*entities.AgentLink, error) {
	data := &entities.AgentLink{}

	err := d.agentLinkTable().Find(db.Cond{
		"id":            id,
		"tenant_id":     tenantId,
		"from_plug_id":  pid,
		"from_agent_id": aid,
	}).One(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DB) AgentLinkDel(tenantId, pid, aid string, id int64) error {
	return d.agentLinkTable().Find(db.Cond{
		"id":            id,
		"tenant_id":     tenantId,
		"from_plug_id":  pid,
		"from_agent_id": aid,
	}).Delete()
}

func (d *DB) AgentLinkList(tenantId, pid, aid string) ([]*entities.AgentLink, error) {

	data := make([]*entities.AgentLink, 0)

	err := d.agentLinkTable().Find(db.Cond{
		"tenant_id":     tenantId,
		"from_plug_id":  pid,
		"from_agent_id": aid,
	}).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DB) AgentLinkListReverse(tenantId, pid, aid string) ([]*entities.AgentLink, error) {
	data := make([]*entities.AgentLink, 0)

	err := d.agentLinkTable().Find(db.Cond{
		"tenant_id":   tenantId,
		"to_plug_id":  pid,
		"to_agent_id": aid,
	}).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// extension

func (d *DB) AgentExtensionNew(tenantId string, data *entities.AgentExtension) error {
	_, err := d.agentExtensionTable().Insert(data)
	return err
}

func (d *DB) AgentExtensionUpdate(tenantId, pid, aid string, id int64, data map[string]interface{}) error {
	return d.agentExtensionTable().Find(db.Cond{
		"id":        id,
		"tenant_id": tenantId,
		"plug_id":   pid,
		"agent_id":  aid,
	}).Update(data)
}

func (d *DB) AgentExtensionGet(tenantId, pid, aid string, id int64) (*entities.AgentExtension, error) {
	data := &entities.AgentExtension{}

	err := d.agentExtensionTable().Find(db.Cond{
		"id":        id,
		"tenant_id": tenantId,
		"plug_id":   pid,
		"agent_id":  aid,
	}).One(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DB) AgentExtensionDel(tenantId, pid, aid string, id int64) error {
	return d.agentExtensionTable().Find(db.Cond{
		"id":        id,
		"tenant_id": tenantId,
		"plug_id":   pid,
		"agent_id":  aid,
	}).Delete()
}

func (d *DB) AgentExtensionList(tenantId, pid, aid string) ([]*entities.AgentExtension, error) {
	data := make([]*entities.AgentExtension, 0)

	err := d.agentExtensionTable().Find(db.Cond{
		"tenant_id": tenantId,
		"plug_id":   pid,
		"agent_id":  aid,
	}).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// resources

func (d *DB) AgentResourceNew(tenantId string, data *entities.AgentResource) error {
	_, err := d.agentResourceTable().Insert(data)
	return err
}

func (d *DB) AgentResourceUpdate(tenantId, pid, aid, slug string, data map[string]interface{}) error {
	return d.agentResourceTable().Find(db.Cond{
		"slug":      slug,
		"plug_id":   pid,
		"agent_id":  aid,
		"tenant_id": tenantId,
	}).One(data)
}

func (d *DB) AgentResourceGet(tenantId, pid, aid, slug string) (*entities.AgentResource, error) {
	data := &entities.AgentResource{}

	err := d.agentResourceTable().Find(db.Cond{
		"slug":      slug,
		"plug_id":   pid,
		"agent_id":  aid,
		"tenant_id": tenantId,
	}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) AgentResourceDel(tenantId, pid, aid, slug string) error {
	return d.agentResourceTable().Find(db.Cond{
		"slug":      slug,
		"plug_id":   pid,
		"agent_id":  aid,
		"tenant_id": tenantId,
	}).Delete()
}

func (d *DB) AgentResourceList(tenantId, pid, aid string) ([]*entities.AgentResource, error) {
	data := make([]*entities.AgentResource, 0)
	err := d.agentResourceTable().Find(db.Cond{
		"plug_id":   pid,
		"agent_id":  aid,
		"tenant_id": tenantId,
	}).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DB) ListResourcePairs(tenantId string, pid, aid string) ([]entities.ResourcePair, error) {
	// fixme => use join ?

	rrids := make([]*entities.AgentResource, 0)
	err := d.agentResourceTable().Find(db.Cond{
		"plug_id":   pid,
		"agent_id":  aid,
		"tenant_id": tenantId,
	}).All(&rrids)
	if err != nil {
		return nil, err
	}

	rids := make([]string, 0, len(rrids))
	for _, r := range rrids {
		rids = append(rids, r.ResourceId)
	}

	ress := make([]*entities.Resource, 0)
	err = d.resTable().Find(db.Cond{
		"tenant_id": tenantId,
		"id IN":     rids,
	}).All(&ress)
	if err != nil {
		return nil, err
	}

	resp := make([]entities.ResourcePair, len(rrids))

	for _, ar := range rrids {
		var res *entities.Resource

		for _, r := range ress {
			if ar.ResourceId == r.Id {
				res = r
				break
			}
		}

		resp = append(resp, entities.ResourcePair{
			AgentResource: ar,
			Resource:      res,
		})
	}

	return resp, nil

}

func (d *DB) AgentExtensionListByPlug(tenantId, pid string) ([]*entities.AgentExtension, error) {
	data := make([]*entities.AgentExtension, 0)

	err := d.agentExtensionTable().Find(db.Cond{
		"tenant_id": tenantId,
		"plug_id":   pid,
	}).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DB) AgentLinkListByPlug(tenantId, pid string) ([]*entities.AgentLink, error) {
	data := make([]*entities.AgentLink, 0)

	err := d.agentLinkTable().Find(db.Cond{
		"tenant_id": tenantId,
		"plug_id":   pid,
	}).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// private

func (d *DB) agentLinkTable() db.Collection {
	return dbutils.Table(d.session, "agent_links")
}

func (d *DB) agentExtensionTable() db.Collection {
	return dbutils.Table(d.session, "agent_extensions")
}

func (d *DB) agentResourceTable() db.Collection {
	return dbutils.Table(d.session, "agent_resources")
}
