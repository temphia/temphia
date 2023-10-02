package ecache

import (
	"sync"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type ecache struct {
	corehub store.CoreHub

	agents map[string]*entities.Agent
	aLock  sync.RWMutex
	aChan  chan agentState

	plugs map[string]*entities.Plug
	pLock sync.RWMutex
	pChan chan plugState

	// targethooks map[int64]*entities.TargetHook
	// targetapps  map[int64]*entities.TargetApp
	// hookIndex   map[string][]*entities.TargetHook
}

func New(corehub store.CoreHub) *ecache {
	ec := &ecache{
		corehub: corehub,
		agents:  make(map[string]*entities.Agent),
		aLock:   sync.RWMutex{},
		aChan:   make(chan agentState),

		plugs: make(map[string]*entities.Plug),
		pLock: sync.RWMutex{},
		pChan: make(chan plugState),
	}

	go ec.evLoop()

	return ec
}

func (e *ecache) GetPlug(tenantId, plugId string) *entities.Plug {

	e.pLock.RLock()
	state := e.plugs[plugId]
	e.pLock.RUnlock()

	if state != nil {
		return state
	}

	wchan := make(chan *entities.Plug)

	e.pChan <- plugState{
		wchan:    wchan,
		plug:     plugId,
		tenantId: tenantId,
	}

	return <-wchan
}

func (e *ecache) GetAgent(tenantId, plug, agent string) *entities.Agent {
	key := plug + agent

	e.aLock.RLock()
	state := e.agents[key]
	e.aLock.RUnlock()

	if state != nil {
		return state
	}

	wchan := make(chan *entities.Agent)

	e.aChan <- agentState{
		wchan:    wchan,
		plug:     plug,
		agent:    agent,
		tenantId: tenantId,
	}

	return <-wchan
}
