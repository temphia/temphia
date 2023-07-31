package agent

import (
	"net/http"
	"sync"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type AgentServer struct {
	agents map[string]*agentState
	aLock  sync.RWMutex

	loadC chan agentL

	ehub    etypes.EngineHub
	app     xtypes.App
	corehub store.CoreHub
}

func New(ehub etypes.EngineHub, app xtypes.App) *AgentServer {
	as := &AgentServer{
		agents: make(map[string]*agentState),
		aLock:  sync.RWMutex{},

		loadC:   make(chan agentL, 10),
		corehub: nil,

		ehub: ehub,
		app:  app,
	}

	go as.evLoop()

	return as
}

type Context struct {
	Writer   http.ResponseWriter
	Request  *http.Request
	TenantId string
	PlugId   string
	AgentId  string
	DomainId int64
}

func (a *AgentServer) Render(ctx Context) {

	as := a.get(ctx.TenantId, ctx.PlugId, ctx.AgentId)
	if as == nil {
		return
	}

	pp.Println(as)

}
