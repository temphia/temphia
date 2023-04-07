package groupd

import (
	"strings"
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

type Groupd struct {
	groups map[string]*instance
	gLock  sync.RWMutex

	eventChan    chan xplane.Message
	newGroupChan chan [2]string
}

func New(msgbus xplane.MsgBus) *Groupd {

	echan := make(chan xplane.Message, 8)

	_, err := msgbus.Subscribe("user_group", echan)
	if err != nil {
		panic(err)
	}

	return &Groupd{
		groups:    map[string]*instance{},
		eventChan: echan,
		gLock:     sync.RWMutex{},
	}
}

func (g *Groupd) RunLoop() {
	for {

		select {
		case msg := <-g.eventChan:
			path := strings.Split(msg.Path, ".")

			g.gLock.RLock()
			ist, ok := g.groups[msg.Tenant+path[1]]
			g.gLock.RUnlock()

			if !ok {
				continue
			}

			go ist.handle(path[0], msg)

		case ginfo := <-g.newGroupChan:

			g.gLock.RLock()
			_, ok := g.groups[ginfo[0]+ginfo[1]]
			g.gLock.RUnlock()
			if ok {
				return
			}

			inst := &instance{
				corehub:  nil,
				tenantId: ginfo[0],
				group:    ginfo[1],
			}

			inst.fetchLatest()

			g.gLock.Lock()
			_, ok = g.groups[ginfo[0]+ginfo[1]]
			if ok {
				g.gLock.Unlock()
				return
			}

			g.groups[ginfo[0]+ginfo[1]] = inst
			g.gLock.Unlock()

		}

	}

}
