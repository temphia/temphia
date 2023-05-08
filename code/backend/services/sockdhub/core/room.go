package core

import (
	"sync"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/thoas/go-funk"
)

type room struct {
	parent      *Sockd
	ns          string
	name        string
	connections map[int64]*Conn
	rlock       sync.Mutex

	// fixme => implement this
	// debugWatch      bool
	// debugWatchConns []*Conn
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
		if funk.ContainsInt64(ignores, cid) {
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

	pp.Println("@pushing_to_tags", tagSet)

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

	conn, ok := r.connections[opts.Id]
	if !ok {
		return false
	}

	if opts.ClearOld {
		for k := range conn.tags {
			delete(conn.tags, k)
		}
	}

	for _, ov := range opts.RemoveTags {
		delete(conn.tags, ov)
	}

	for _, v := range opts.AddTags {
		conn.tags[v] = struct{}{}
	}

	return true

}

func (r *room) AddConn(conn sockdx.Conn, tags []string) {
	r.rlock.Lock()

	id := conn.Id()

	pp.Println("@conn start", id, tags)

	oldconn, ok := r.connections[id]
	if ok {

		pp.Println("@closing old conn")

		oldconn.close(false)
		r.clearConnTags(id)
	}

	iTags := make(map[string]struct{})

	for _, v := range tags {
		iTags[v] = struct{}{}
	}

	c := &Conn{
		parent:  r,
		conn:    conn,
		closed:  false,
		failed:  false,
		writeCh: make(chan []byte),
		tags:    iTags,
	}

	c.start()
	r.connections[id] = c

	r.rlock.Unlock()

	pp.Println("@conn end", id, tags)
	pp.Println("@total_conn", len(r.connections))

	r.parent.logger.Info().
		Int64("conn_id", int64(id)).
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

	for cid, c := range r.connections {
		if funk.ContainsInt64(ignores, cid) {
			continue
		}
		if c.hasTags(tags) {
			tagSet[cid] = struct{}{}
		}
	}

	return tagSet
}

func (r *room) clearConnTags(cid int64) {

}
