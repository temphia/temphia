package rpool

import (
	"math/rand"
	"sync"

	"github.com/temphia/temphia/code/backend/engine/binder"
)

type Pool struct {
	binders    map[string]*binder.Binder
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
		binders:            make(map[string]*binder.Binder),
		plugIndex:          make(map[string]*plugSet),
		agentIndex:         make(map[string]*agentSet),
		slock:              sync.Mutex{},
		totalAccessCounter: 0,
		totalCount:         0,
		maxTotalCount:      100,
		maxSingleCount:     20,
	}
}

func (p *Pool) Borrow(plugId, agentId string) (*binder.Binder, int) {
	p.slock.Lock()
	defer p.slock.Unlock()

	agentEventSet, ok := p.agentIndex[plugId+agentId]
	if !ok {
		return nil, 0
	}

	if agentEventSet.isEmpty() {
		return nil, 0
	}

	event := agentEventSet.binders.Pop()

	if agentEventSet.isEmpty() {
		delete(p.agentIndex, plugId+agentId)
	}

	bind := p.binders[event]
	delete(p.binders, event)

	p.totalAccessCounter = p.totalAccessCounter + 1
	p.totalCount = p.totalCount - 1
	agentEventSet.counter = agentEventSet.counter + 1

	return bind, len(agentEventSet.binders)

}

func (p *Pool) Destroy(plugId string, agentIds []string) {
	p.slock.Lock()
	defer p.slock.Unlock()

	for _, agentId := range agentIds {
		agentSet := p.agentIndex[plugId+agentId]
		if agentSet == nil {
			continue
		}

		for bdr := range agentSet.binders {
			delete(p.binders, bdr)
		}

		agentSet.epochVersion = agentSet.epochVersion + 1
	}

}

func (p *Pool) Return(b *binder.Binder) {
	p.slock.Lock()
	defer p.slock.Unlock()

	pSet, ok := p.plugIndex[b.PlugId]
	if !ok {
		p.plugIndex[b.PlugId] = &plugSet{
			agents: Set{
				b.AgentId: struct{}{},
			},
			bprintId: b.BprintId,
		}
	} else {
		pSet.agents.Push(b.AgentId)
	}

	aset, ok := p.agentIndex[b.PlugId+b.AgentId]
	if !ok {
		aset = &agentSet{
			binders: Set{
				b.EventId: struct{}{},
			},
			counter:      0,
			epochVersion: b.Epoch,
		}
		p.agentIndex[b.PlugId+b.AgentId] = aset
	} else {
		if aset.epochVersion > b.Epoch {
			return
		}

		count := p.calculate(aset.counter)
		if len(aset.binders) > int(count) {
			return
		}

		if p.totalCount > p.maxTotalCount {
			// delete random element
			p.getRandom(b.PlugId, b.AgentId)
			p.totalCount = p.totalCount - 1
			delete(p.binders, b.EventId)
		}

		aset.binders.Push(b.EventId)
	}

	p.totalCount = p.totalCount + 1
	p.binders[b.EventId] = b
}

func (p *Pool) SetEpoch(plug, agent string, e int64) {
	p.slock.Lock()
	defer p.slock.Unlock()

	aset, ok := p.agentIndex[plug+agent]
	if !ok {
		return
	}

	aset.epochVersion = e

	for ev := range aset.binders {
		delete(aset.binders, ev)
		delete(p.binders, ev)
	}
}

// private

func (p *Pool) getRandom(currPlug, currAgent string) *binder.Binder {

	elem := rand.Int31n(int32(len(p.binders)))
	i := 0
	for _, cval := range p.binders {
		i = i + 1
		if i >= int(elem) {
			if cval.PlugId == currPlug || cval.AgentId == currAgent {
				break
			}

			aset := p.agentIndex[cval.PlugId+cval.AgentId]
			if len(aset.binders) < 1 {
				continue
			}

			return cval
		}
	}

	for _, cval := range p.binders {
		if cval.PlugId == currPlug || cval.AgentId == currAgent {
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
