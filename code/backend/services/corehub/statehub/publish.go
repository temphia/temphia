package statehub

import (
	"encoding/json"
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

func (r *StateHub) OnTargetAppChange(tenantId string, id int64, data *entities.TargetApp) error {
	return r.publishIntId(tenantId, "target_app", id, data)
}

func (r *StateHub) OnTargetHookChange(tenantId string, id int64, data *entities.TargetHook) error {
	return r.publishIntId(tenantId, "target_hook", id, data)
}

func (r *StateHub) OnResourceChange(tenantId, id string, data *entities.Resource) error {
	return r.publish(tenantId, "resource", id, data)
}

func (r *StateHub) OnUserGroupChange(tenantId, id string, data *entities.UserGroup) error {
	return r.publish(tenantId, "user_group", id, data)
}

func (r *StateHub) OnDataGroupChange(tenantId, gid string, data *entities.TableGroup) error {
	return r.publish(tenantId, "user_group", gid, data)
}

func (r *StateHub) OnDataTableChange(tenantId, gid, tid string, data *entities.Table) error {

	return nil
}

func (r *StateHub) OnDataColumnChange(tenantId, gid, tid, cid string, data *entities.Column) error {

	return nil
}

func (r *StateHub) OnTenantChange(id string, data *entities.Tenant) error {

	msg := xplane.Message{
		Data:  "",
		Path:  "create",
		Topic: "tenant",
	}

	if data == nil {
		msg.Path = "delete"
	} else if id != "" {
		msg.Path = "update"
	}

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg.Data = string(out)
	_, err = r.msgbus.Submit("tenant", msg)

	return err
}

func (r *StateHub) OnDomainChange(tenantId string, id int64, data *entities.TenantDomain) error {
	return r.publishIntId(tenantId, "domain", id, data)
}

func (r *StateHub) publishIntId(tenantId, topic string, id int64, data any) error {
	sid := ""
	if id != 0 {
		sid = fmt.Sprint(id)
	}

	return r.publish(tenantId, topic, sid, data)
}

func (r *StateHub) publish(tenantId, topic, id string, data any) error {

	msg := xplane.Message{
		Data:  "",
		Path:  "create",
		Topic: topic,
	}

	if data == nil {
		msg.Path = "delete"
	} else if id != "" {
		msg.Path = "update"
	}

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg.Data = string(out)
	_, err = r.msgbus.Submit("tenant", msg)

	return err
}
