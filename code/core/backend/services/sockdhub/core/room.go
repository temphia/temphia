package core

import (
	"strings"
	"sync"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/core/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/sockdx"

	"github.com/thoas/go-funk"
)

type room struct {
	parent      *Sockd
	ns          string
	name        string
	connections map[int64]*Conn
	tags        map[string][]int64

	rlock sync.Mutex
}

func (r *room) sendDirect(connId int64, payload []byte) error {

	msg := sockdx.Message{
		Room:        r.name,
		Type:        sockdx.MESSAGE_SERVER_DIRECT,
		ServerIdent: r.parent.serverIdent,
		XID:         xid.New().String(),
		Payload:     payload,
	}

	out, err := msg.JSON()
	if err != nil {
		r.parent.logger.Warn().
			Str("id", msg.XID).
			Str("mtype", sockdx.MESSAGE_SERVER_DIRECT).
			Str("room", r.name).
			Bytes("data", payload).
			Msg(logid.SockdSendMarshelErr)
		return err
	}

	r.parent.logger.Info().
		Str("id", msg.XID).
		Int64("target", int64(connId)).
		Msg(logid.SockdSendDirect)

	c, ok := r.connections[connId]
	if !ok {
		if r.parent.syncer != nil {
			r.parent.syncer.SyncMessage(r.ns, r.name, sockdx.MESSAGE_SERVER_DIRECT, msg)
		}
		return nil
	}

	c.write(out)

	return nil

}

func (r *room) sendDirectBatch(conns []int64, payload []byte) error {

	msg := sockdx.Message{
		Room:        r.name,
		Type:        sockdx.MESSAGE_SERVER_DIRECT,
		ServerIdent: r.parent.serverIdent,
		XID:         xid.New().String(),
		Payload:     payload,
	}

	out, err := msg.JSON()
	if err != nil {
		r.parent.logger.Warn().
			Str("id", msg.XID).
			Str("mtype", sockdx.MESSAGE_SERVER_DIRECT).
			Str("room", r.name).
			Bytes("data", payload).
			Msg(logid.SockdSendMarshelErr)
		return err
	}

	r.parent.logger.Info().
		Str("id", msg.XID).
		Interface("targets", conns).
		Msg(logid.SockdSendDirectBatch)

	pending := make([]int64, 0)
	for _, cid := range conns {
		c, ok := r.connections[cid]
		if !ok {
			pending = append(pending, cid)
		}
		c.write(out)
	}

	if len(pending) != 0 && r.parent.syncer != nil {
		msg.TargetIds = pending
		r.parent.syncer.SyncMessage(r.ns, r.name, sockdx.MESSAGE_SERVER_DIRECT, msg)
	}

	return nil
}

func (r *room) sendBroadcast(ignores []int64, payload []byte) error {

	msg := sockdx.Message{
		Room:        r.name,
		Type:        sockdx.MESSAGE_SERVER_BROADCAST,
		ServerIdent: r.parent.serverIdent,
		XID:         xid.New().String(),
		Payload:     payload,
	}

	out, err := msg.JSON()
	if err != nil {
		r.parent.logger.Warn().
			Str("id", msg.XID).
			Str("mtype", sockdx.MESSAGE_SERVER_BROADCAST).
			Str("room", r.name).
			Bytes("data", payload).
			Msg(logid.SockdSendMarshelErr)
		return err
	}

	r.parent.logger.Info().
		Str("id", msg.XID).
		Interface("ignores", ignores).
		Msg(logid.SockdSendBroadcast)

	for cid, conn := range r.connections {
		if containsCid(ignores, cid) {
			continue
		}
		conn.write(out)
	}

	if r.parent.syncer != nil {
		r.parent.syncer.SyncMessage(r.ns, r.name, sockdx.MESSAGE_SERVER_BROADCAST, out)
	}

	return nil

}

func (r *room) sendTagged(tags []string, ignores []int64, payload []byte) error {
	tagSet := r.cidFromTags(tags, ignores)
	if len(tagSet) == 0 {
		return nil
	}

	msg := sockdx.Message{
		Room:        r.name,
		Type:        sockdx.MESSAGE_SERVER_PUBLISH,
		ServerIdent: r.parent.serverIdent,
		XID:         xid.New().String(),
		Payload:     payload,
	}

	out, err := msg.JSON()
	if err != nil {

		r.parent.logger.Warn().
			Str("id", msg.XID).
			Str("mtype", sockdx.MESSAGE_SERVER_PUBLISH).
			Str("room", r.name).
			Bytes("data", payload).
			Msg(logid.SockdSendMarshelErr)

		return err
	}

	r.parent.logger.Info().
		Str("id", msg.XID).
		Interface("targets", tags).
		Interface("ignores", ignores).
		Msg(logid.SockdSendTagged)

	for ci := range tagSet {
		conn, ok := r.connections[ci]
		if !ok {
			continue
		}

		conn.write(out)
	}

	if r.parent.syncer != nil {
		msg.TargetTags = tags
		msg.IgnoreConns = ignores
		r.parent.syncer.SyncMessage(r.ns, r.name, sockdx.MESSAGE_SERVER_PUBLISH, msg)
	}

	return nil

}

func (r *room) roomUpdateTags(opts sockdx.UpdateTagOptions) bool {
	r.rlock.Lock()
	defer r.rlock.Unlock()

	_, ok := r.connections[opts.Id]
	if !ok {
		return false
	}

	for tagkey, conns := range r.tags {
		if strings.HasPrefix(tagkey, "@") {
			continue
		}

		if opts.ClearOld {
			if containsCid(conns, opts.Id) {
				r.tags[tagkey] = filterCid(conns, opts.Id)
			}
		} else {
			if funk.ContainsString(opts.RemoveTags, tagkey) {
				r.tags[tagkey] = filterCid(conns, opts.Id)
			}
		}
	}

	for _, tag := range opts.AddTags {
		old, ok := r.tags[tag]
		if ok {
			old = []int64{opts.Id}
		} else {
			old = append(old, opts.Id)
		}

		r.tags[tag] = old
	}

	return true

}

func (r *room) AddConn(conn sockdx.Conn, tags []string) {
	r.rlock.Lock()

	id := conn.Id()

	oldconn, ok := r.connections[id]
	if ok {
		oldconn.close(false)
		r.clearConnTags(id)
	}

	c := &Conn{
		parent:  r,
		conn:    conn,
		closed:  false,
		failed:  false,
		writeCh: make(chan []byte),
	}

	for _, tag := range tags {
		old, ok := r.tags[tag]
		if ok {
			old = []int64{id}
		} else {
			old = append(old, id)
		}

		r.tags[tag] = old
	}

	c.start()
	r.connections[conn.Id()] = c

	r.rlock.Unlock()

	r.parent.logger.Info().
		Int64("conn_id", int64(conn.Id())).
		Msg(logid.SockdNewConnection)

}

func (r *room) kickRoomConn(cid int64) {
	r.rlock.Lock()
	defer r.rlock.Unlock()

	conn, ok := r.connections[cid]
	if !ok {
		return
	}

	conn.close(false)
	delete(r.connections, cid)
	r.clearConnTags(cid)
}

func (r *room) close() {
	r.rlock.Lock()
	defer r.rlock.Unlock()

	for _, c := range r.connections {
		c.close(false)
	}
}

// private

func (r *room) getLogger() *zerolog.Logger {
	return &r.parent.logger
}

func (r *room) cidFromTags(tags []string, ignores []int64) map[int64]struct{} {
	tagSet := make(map[int64]struct{}, 10)

	r.rlock.Lock()
	defer r.rlock.Unlock()

	head, ok := r.tags[tags[0]]
	if !ok {
		return nil
	}

	for _, cid := range head {
		tagSet[cid] = struct{}{}
	}

	for _, tag := range tags {
		cids, ok := r.tags[tag]
		if !ok {
			return nil
		}

		for ci := range tagSet {
			if !containsCid(cids, ci) {
				delete(tagSet, ci)
			}
		}
	}

	return tagSet
}

func (r *room) clearConnTags(cid int64) {

	for k, all := range r.tags {
		if !containsCid(all, (cid)) {
			continue
		}
		r.tags[k] = filterCid(all, cid)
	}

}

func containsCid(s []int64, v int64) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func filterCid(s []int64, v int64) []int64 {
	resp := make([]int64, 0, len(s)-1)
	for _, vv := range s {
		if vv == v {
			continue
		}
		resp = append(resp, vv)
	}

	return resp
}
