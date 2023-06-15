package core

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
)

var _ sockdx.SockdCore = (*Sockd)(nil)

type Options struct {
	ServerIdent string
	Syncer      sockdx.PeerSync
	SysHelper   sockdx.SystemHelper
	Logger      zerolog.Logger
}

type Sockd struct {
	serverIdent string

	rooms       map[string]*room
	tenantRooms map[string][]string
	roomLock    sync.RWMutex

	syncer sockdx.PeerSync
	system sockdx.SystemHelper
	logger zerolog.Logger
}

func New(opts sockdx.Options) *Sockd {
	s := &Sockd{
		rooms:       make(map[string]*room),
		tenantRooms: make(map[string][]string),
		serverIdent: opts.ServerIdent,
		roomLock:    sync.RWMutex{},
		syncer:      opts.Syncer,
		system:      opts.SysHelper,
		logger:      opts.Logger,
	}

	go s.debug()
	return s
}

func (s *Sockd) NewConnection(opts sockdx.ConnOptions) error {
	return s.newConnection(opts)
}

func (s *Sockd) SendDirect(ns, room string, connId int64, payload []byte) error {
	return s.sendDirect(ns, room, connId, payload)
}

func (s *Sockd) SendDirectBatch(ns, room string, conns []int64, payload []byte) error {
	return s.sendDirectBatch(ns, room, conns, payload)
}

func (s *Sockd) SendBroadcast(ns, room string, ignores []int64, payload []byte) error {
	return s.sendBroadcast(ns, room, ignores, payload)
}

func (s *Sockd) SendTagged(ns, room string, tags []string, ignores []int64, payload []byte) error {
	return s.sendTagged(ns, room, tags, ignores, payload)
}

func (s *Sockd) RoomUpdateTags(ns, roomId string, opts sockdx.UpdateTagOptions) error {
	return s.roomUpdateTags(ns, roomId, opts)
}

func (s *Sockd) Poll(ns, roomId string, max, min int64) (*sockdx.PollResponse, error) {
	return s.poll(ns, roomId, max, min)
}

func (s *Sockd) LocalListConns(ns string) (map[int64]string, error) {
	return s.localListConns(ns)
}

func (s *Sockd) LocalListRoomConns(ns, room string) (map[int64][]string, error) {
	return s.localListRoomConns(ns, room)
}

func (s *Sockd) LocalKickRoomConn(ns, room string, cid int64) error {
	return s.localKickRoomConn(ns, room, cid)
}

func (s *Sockd) LocalCloseRoom(ns, room string) error {
	return s.localCloseRoom(ns, room)
}
