package core

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"
	"github.com/tidwall/gjson"
)

func (c *Conn) processPacket2(msg []byte) {
	msgFromid := gjson.GetBytes(msg, "from_id").Int()
	msgType := gjson.GetBytes(msg, "type").String()
	msgXid := gjson.GetBytes(msg, "xid").String()

	c.parent.getLogger().Info().
		Int64("conn_id", int64(c.conn.Id())).
		Str("msg_type", msgType).Int64("from_id", msgFromid).
		Str("xid", msgXid).
		Msg(logid.SockdMsgReceived)

	c.parent.getLogger().Debug().
		Str("xid", msgXid).
		RawJSON("data", msg).
		Msg(logid.SockdMsgReceivedDebug)

	if msgFromid != int64(c.conn.Id()) {
		c.parent.getLogger().Warn().
			Str("xid", msgXid).
			Msg(logid.SockdMsgInvalidId)
		return
	}

	switch msgType {

	case sockdx.MESSAGE_CLIENT_DIRECT:
		rawTargets := gjson.GetBytes(msg, "target_ids").Array()
		if len(rawTargets) == 0 {
			c.parent.getLogger().Warn().
				Str("xid", msgXid).
				Msg(logid.SockdMsgEmptyTargetIds)
			return
		}

		for _, t := range rawTargets {
			c, ok := c.parent.connections[int64(t.Int())]
			if !ok {
				if c.parent.parent.syncer != nil {
					c.parent.parent.syncer.SyncMessage(c.parent.ns, c.parent.name, msgType, msg)
				}
				continue
			}

			c.write(msg)
		}

	case sockdx.MESSAGE_CLIENT_PUBLISH:
		rawTargets := gjson.GetBytes(msg, "target_tags").Array()
		targets := make([]string, 0, len(rawTargets))

		if len(rawTargets) == 0 {
			c.parent.getLogger().Warn().
				Str("xid", msgXid).
				Msg(logid.SockdMsgEmptyTargetTags)
		}

		for _, t := range rawTargets {
			targets = append(targets, t.String())
		}

		tset := c.parent.cidFromTags(targets, []int64{c.conn.Id()})
		for ci := range tset {
			c, ok := c.parent.connections[int64(ci)]
			if !ok {
				continue
			}
			c.write(msg)
		}

		if c.parent.parent.syncer != nil {
			c.parent.parent.syncer.SyncMessage(c.parent.ns, c.parent.name, msgType, msg)
		}

	case sockdx.MESSAGE_CLIENT_BROADCAST:
		for cid, conn := range c.parent.connections {
			if cid == c.conn.Id() {
				continue
			}
			conn.write(msg)
		}
		if c.parent.parent.syncer != nil {
			c.parent.parent.syncer.SyncMessage(c.parent.ns, c.parent.name, msgType, msg)
		}
	case sockdx.MESSAGE_CLIENT_SYSTEM:
		c.handleClientSystem(msgFromid, msg)
	default:
		c.parent.getLogger().Warn().
			Str("xid", msgXid).
			Msg(logid.SockdMsgInvalidMType)
	}

}

func (c *Conn) handleClientSystem(cid int64, data []byte) {

	sysType := gjson.GetBytes(data, "payload.type").String()
	token := gjson.GetBytes(data, "payload.token").String()

	switch sysType {
	case "update_tags":
	case "add_tags":

		pp.Println("FIXME ", token)
	default:
		pp.Println("NOT IMPLEMENTED")
	}

}
