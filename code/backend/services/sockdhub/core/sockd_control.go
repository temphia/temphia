package core

import (
	"github.com/temphia/temphia/code/backend/xtypes/service/sockdx"
	"github.com/thoas/go-funk"
)

func (s *Sockd) localListConns(ns string) (map[int64]string, error) {

	rooms := s.tenantRooms[ns]

	resp := make(map[int64]string)

	for _, room := range rooms {
		r := s.roomGet(ns, room, false)
		if r == nil {
			continue
		}

		for ci := range r.connections {
			resp[ci] = room
		}
	}

	return resp, nil
}

func (s *Sockd) localListRoomConns(ns, room string) (map[int64][]string, error) {
	r := s.roomGet(ns, room, false)
	if r == nil {
		return nil, sockdx.ErrRoomNotFound
	}

	resp := make(map[int64][]string, 0)

	for tag, cids := range r.tags {
		for _, cid := range cids {

			old, ok := resp[cid]
			if ok {
				old = []string{tag}
			} else {
				old = append(old, tag)
			}
			resp[cid] = old
		}
	}

	return resp, nil
}

func (s *Sockd) localKickRoomConn(ns, room string, cid int64) error {
	r := s.roomGet(ns, room, false)
	if r == nil {
		return sockdx.ErrRoomNotFound
	}

	r.kickRoomConn(cid)

	return nil
}

func (s *Sockd) localCloseRoom(ns, room string) error {
	s.roomLock.Lock()

	r, ok := s.rooms[room]
	if !ok {
		return sockdx.ErrRoomNotFound
	}

	delete(s.rooms, room)

	s.tenantRooms[ns] = funk.FilterString(s.tenantRooms[ns], func(s string) bool {
		return s == room
	})

	s.roomLock.Unlock()

	r.close()

	return nil
}
