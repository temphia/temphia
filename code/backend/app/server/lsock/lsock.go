package lsock

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/xserver"
)

type LSock struct {
	subs  map[int64]xserver.LSubcriber
	sLock sync.Mutex

	wsconns map[int64]*lsockWs
	wLock   sync.RWMutex

	counter int64
}

func New() *LSock {
	return &LSock{
		subs:    make(map[int64]xserver.LSubcriber),
		sLock:   sync.Mutex{},
		wsconns: make(map[int64]*lsockWs),
		wLock:   sync.RWMutex{},
		counter: 1,
	}
}

func (l *LSock) API(g *gin.RouterGroup) {
	/*

	  /z/lsock/rpc/:action
	  /z/lsock/ws?token=xxxyyy

	*/
}

func (l *LSock) Register(s xserver.LSubcriber) int64 {
	return 0
}

func (l *LSock) Send(eid int64, name string, data []byte) {

}

func (l *LSock) SendWS(eid, tid int64, name string, data []byte) {

}
