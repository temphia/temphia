package lsock

import (
	"encoding/json"
	"sync"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/xserver"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz"
)

var (
	_ xserver.LSock = (*LSock)(nil)
)

type LSock struct {
	subs  map[int64]xserver.LSubcriber
	sLock sync.Mutex

	wsconns map[int64]*lsockWs
	wLock   sync.RWMutex

	counter int64

	notz   xnotz.Notz
	signer service.Signer
}

func New(notz xnotz.Notz, signer service.Signer) *LSock {
	return &LSock{
		subs:    make(map[int64]xserver.LSubcriber),
		sLock:   sync.Mutex{},
		wsconns: make(map[int64]*lsockWs),
		wLock:   sync.RWMutex{},
		counter: 1,
		notz:    notz,
		signer:  signer,
	}
}

func (l *LSock) API(g *gin.RouterGroup) {

	g.POST("/rpc/:action", l.apiRPCAction)
	g.GET("/ws", l.apiWS)
	g.POST("/register", l.apiRegister)

}

func (l *LSock) Register(s xserver.LSubcriber) int64 {

	sid := atomic.AddInt64(&l.counter, 1)

	l.sLock.Lock()
	l.subs[sid] = s
	l.sLock.Unlock()

	return sid
}

type LSPacket struct {
	Name string          `json:"name,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

func (l *LSock) SendWS(sid int64, name string, data []byte) error {
	l.wLock.RLock()
	wcon := l.wsconns[sid]
	l.wLock.RUnlock()

	pak := &LSPacket{
		Name: name,
		Data: data,
	}

	out, err := json.Marshal(pak)
	if err != nil {
		return err
	}

	wcon.wChan <- out
	return nil

}
