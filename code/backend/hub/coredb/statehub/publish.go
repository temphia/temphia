package statehub

import (
	"encoding/json"
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

func (r *StateHub) OnTargetAppChange(tenantId string, id int64, data any) {
	r.publishIntId(tenantId, "target_app", id, data)
}

func (r *StateHub) OnTargetHookChange(tenantId string, id int64, data any) {
	r.publishIntId(tenantId, "target_hook", id, data)
}

func (r *StateHub) OnResourceChange(tenantId, id string, data any) {
	r.publish(tenantId, "resource", id, data)
}

func (r *StateHub) OnUserGroupChange(tenantId, id string, data any) {
	r.publish(tenantId, "user_group", id, data)
}

func (r *StateHub) OnDataGroupChange(tenantId, gid string, data any) {
	r.publish(tenantId, "user_group", gid, data)
}

func (r *StateHub) OnDataTableChange(tenantId, gid, tid string, data any) {

}

func (r *StateHub) OnDataColumnChange(tenantId, gid, tid, cid string, data any) {

}

func (r *StateHub) OnTenantChange(id string, data any) {

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
		r.logger.
			Err(err).
			Str("path", msg.Path).
			Str("topic", msg.Topic).
			Msg(logid.StatehubMsgMarshelErr)
	}

	msg.Data = string(out)
	_, err = r.msgbus.Submit("tenant", msg)
	if err != nil {
		r.logger.
			Err(err).
			Str("path", msg.Path).
			Str("topic", msg.Topic).
			Str("data", msg.Data).
			Msg(logid.StatehubErr)
	}

}

func (r *StateHub) OnDomainChange(tenantId string, id int64, data any) {
	r.publishIntId(tenantId, "domain", id, data)
}

func (r *StateHub) publishIntId(tenantId, topic string, id int64, data any) {
	sid := ""
	if id != 0 {
		sid = fmt.Sprint(id)
	}

	r.publish(tenantId, topic, sid, data)
}

func (r *StateHub) publish(tenantId, topic, id string, data any) {

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
		r.logger.
			Err(err).
			Str("path", msg.Path).
			Str("topic", msg.Topic).
			Msg(logid.StatehubMsgMarshelErr)
	}

	msg.Data = string(out)
	_, err = r.msgbus.Submit("tenant", msg)
	if err != nil {
		r.logger.
			Err(err).
			Str("path", msg.Path).
			Str("topic", msg.Topic).
			Str("data", msg.Data).
			Msg(logid.StatehubErr)
	}

}
