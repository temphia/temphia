package core

import "time"

var Debug = true

func (s *Sockd) debug() {
	if !Debug {
		return
	}

	connAndTags := make(map[int64][]string)

	for {
		time.Sleep(time.Second * 10)

		for _, room := range s.rooms {

			// clear the old values
			for k := range connAndTags {
				delete(connAndTags, k)
			}

			for tagkey, conns := range room.tags {
				for _, v := range conns {
					old, ok := connAndTags[v]
					if !ok {
						old = []string{tagkey}
					} else {
						old = append(old, tagkey)
					}
					connAndTags[v] = old
				}
			}

			s.logger.Debug().
				Str("tenant_id", room.ns).
				Str("room", room.name).
				Interface("connections", connAndTags).Msg("debug_connections")

		}

	}

}
