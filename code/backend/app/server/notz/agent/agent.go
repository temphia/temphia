package agent

import (
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz"
)

type AgentNotz struct {
	agents map[string]*agentState
	aLock  sync.RWMutex

	loadC chan agentL

	ehub    etypes.EngineHub
	corehub store.CoreHub
	cabinet store.CabinetHub
}

func New(ehub etypes.EngineHub, corehub store.CoreHub, cabinet store.CabinetHub) *AgentNotz {
	as := &AgentNotz{
		agents: make(map[string]*agentState),
		aLock:  sync.RWMutex{},

		loadC:   make(chan agentL, 10),
		corehub: corehub,
		cabinet: cabinet,
		ehub:    ehub,
	}

	go as.evLoop()

	return as
}

func (a *AgentNotz) Render(ctx xnotz.Context) {

	as := a.get(ctx.TenantId, ctx.PlugId, ctx.AgentId)
	if as == nil {
		return
	}

	a.spaRender(ctx, as)
}
