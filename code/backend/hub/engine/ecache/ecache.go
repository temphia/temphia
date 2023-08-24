package ecache

import (
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type ecache struct {
	corehub store.CoreHub
	agents  map[string]*entities.Agent
	aLock   sync.RWMutex
	aChan   chan agentstate

	// plugs       map[string]*entities.Plug
	// targethooks map[int64]*entities.TargetHook
	// targetapps  map[int64]*entities.TargetApp
	// hookIndex   map[string][]*entities.TargetHook
}

func New(corehub store.CoreHub) *ecache {
	ec := &ecache{
		corehub: corehub,
		agents:  make(map[string]*entities.Agent),
	}

	go ec.evLoop()

	return ec
}

func (e *ecache) GetAgent(tenantId, plug, agent string) *entities.Agent {
	key := plug + agent

	e.aLock.RLock()
	state := e.agents[key]
	e.aLock.RUnlock()

	if state == nil {
		return state
	}

	wchan := make(chan *entities.Agent)

	e.aChan <- agentstate{
		wchan:    wchan,
		plug:     plug,
		agent:    agent,
		tenantId: tenantId,
	}

	return <-wchan
}
