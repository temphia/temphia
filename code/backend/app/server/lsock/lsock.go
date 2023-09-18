package lsock

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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

	remotes map[int64]*xserver.REInfo
	rLock   sync.RWMutex

	counter int64

	notz   xnotz.Notz
	signer service.Signer
}

func New(notz xnotz.Notz, signer service.Signer) *LSock {
	return &LSock{
		subs:  make(map[int64]xserver.LSubcriber),
		sLock: sync.Mutex{},

		remotes: make(map[int64]*xserver.REInfo),
		rLock:   sync.RWMutex{},

		counter: 1,
		notz:    notz,
		signer:  signer,
	}
}

func (l *LSock) API(g *gin.RouterGroup) {

	g.POST("/rpc/:action", l.apiRPCAction)
	g.POST("/register", l.apiRegister)

}

func (l *LSock) Register(s xserver.LSubcriber) int64 {

	sid := atomic.AddInt64(&l.counter, 1)

	l.sLock.Lock()
	l.subs[sid] = s
	l.sLock.Unlock()

	return sid
}

func (l *LSock) SendRPC(iid int64, name string, data []byte) ([]byte, error) {

	l.rLock.RLock()
	rinfo := l.remotes[iid]
	l.rLock.RUnlock()

	prefix := rinfo.RPCPrefix
	if prefix == "" {
		prefix = "/"
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s%s/%s", rinfo.Addr, rinfo.RPCPrefix, name), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", rinfo.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)

}
