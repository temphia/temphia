package rpool

import (
	"math/rand"
	"sync"

	"github.com/temphia/temphia/code/backend/engine/binders/standard"
)

type Pool struct {
	binders    map[string]*standard.Binder
	plugIndex  map[string]*plugSet  // plug_id => []agent_id
	agentIndex map[string]*agentSet // (plug_id + agent_id ) => []event_id
	slock      sync.Mutex

	totalAccessCounter int64
	totalCount         int64
	maxTotalCount      int64
	maxSingleCount     int64
}

func NewPool() Pool {
	return Pool{
		binders:            make(map[string]*standard.Binder),
		plugIndex:          make(map[string]*plugSet),
		agentIndex:         make(map[string]*agentSet),
		slock:              sync.Mutex{},
		totalAccessCounter: 0,
		totalCount:         0,
		maxTotalCount:      100,
		maxSingleCount:     20,
	}
}

func (p *Pool) Borrow(plugId, agentId string) (*standard.Binder, int) {
	p.slock.Lock()
	defer p.slock.Unlock()

	agentEventSet, ok := p.agentIndex[plugId+agentId]
	if !ok {
		return nil, 0
	}

	if agentEventSet.isEmpty() {
		return nil, 0
	}

	event := agentEventSet.events.Pop()

	if agentEventSet.isEmpty() {
		delete(p.agentIndex, plugId+agentId)
	}

	bind := p.binders[event]
	delete(p.binders, event)

	p.totalAccessCounter = p.totalAccessCounter + 1
	p.totalCount = p.totalCount - 1
	agentEventSet.counter = agentEventSet.counter + 1

	return bind, len(agentEventSet.events)

}

func (p *Pool) Return(b *standard.Binder) {
	p.slock.Lock()
	defer p.slock.Unlock()

	pSet, ok := p.plugIndex[b.Handle.PlugId]
	if !ok {
		p.plugIndex[b.Handle.PlugId] = &plugSet{
			agents: Set{
				b.Handle.AgentId: struct{}{},
			},
			bprintId: b.Handle.BprintId,
		}
	} else {
		pSet.agents.Push(b.Handle.AgentId)
	}

	aset, ok := p.agentIndex[b.Handle.PlugId+b.Handle.AgentId]
	if !ok {
		aset = &agentSet{
			events: Set{
				b.Handle.EventId: struct{}{},
			},
			counter:      0,
			epochVersion: b.Epoch,
		}
		p.agentIndex[b.Handle.PlugId+b.Handle.AgentId] = aset
	} else {
		if aset.epochVersion > b.Epoch {
			return
		}

		count := p.calculate(aset.counter)
		if len(aset.events) > int(count) {
			return
		}

		if p.totalCount > p.maxTotalCount {
			// delete random element
			p.getRandom(b.Handle.PlugId, b.Handle.AgentId)
			p.totalCount = p.totalCount - 1
			delete(p.binders, b.Handle.EventId)
		}

		aset.events.Push(b.Handle.EventId)
	}

	p.totalCount = p.totalCount + 1
	p.binders[b.Handle.EventId] = b
}

func (p *Pool) SetEpoch(plug, agent string, e int64) {
	p.slock.Lock()
	defer p.slock.Unlock()

	aset, ok := p.agentIndex[plug+agent]
	if !ok {
		return
	}

	aset.epochVersion = e

	for ev := range aset.events {
		delete(aset.events, ev)
		delete(p.binders, ev)
	}
}

// private

func (p *Pool) getRandom(currPlug, currAgent string) *standard.Binder {

	elem := rand.Int31n(int32(len(p.binders)))
	i := 0
	for _, cval := range p.binders {
		i = i + 1
		if i >= int(elem) {
			if cval.Handle.PlugId == currPlug || cval.Handle.AgentId == currAgent {
				break
			}

			aset := p.agentIndex[cval.Handle.PlugId+cval.Handle.AgentId]
			if len(aset.events) < 1 {
				continue
			}

			return cval
		}
	}

	for _, cval := range p.binders {
		if cval.Handle.PlugId == currPlug || cval.Handle.AgentId == currAgent {
			continue
		}

		return cval
	}

	return nil
}

func (p *Pool) calculate(singleCounter int64) int64 {
	if p.totalAccessCounter == 0 {
		return 0
	}

	count := (singleCounter * p.maxTotalCount) / p.totalAccessCounter

	if count > p.maxSingleCount {
		return p.maxSingleCount
	}

	if count < 1 {
		return 1
	}

	return count

}
