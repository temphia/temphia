package ecache

import (
	"github.com/k0kubun/pp"
)

func (e *ecache) evLoop() {

	go e.elAgent()
	go e.elPlug()

}

func (e *ecache) elAgent() {
	for {

		asq := <-e.aChan

		key := asq.plug + asq.agent

		agent, err := e.corehub.AgentGet(asq.tenantId, asq.plug, asq.agent)
		if err != nil {
			pp.Println("@get_agent_err", err.Error())
			asq.wchan <- nil
			continue
		}

		e.aLock.Lock()
		e.agents[key] = agent
		e.aLock.Unlock()

		asq.wchan <- agent

	}

}

func (e *ecache) elPlug() {

	for {
		asq := <-e.pChan

		plug, err := e.corehub.PlugGet(asq.tenantId, asq.plug)
		if err != nil {
			pp.Println("@get_agent_err", err.Error())
			asq.wchan <- nil
			continue
		}

		e.aLock.Lock()
		e.plugs[asq.plug] = plug
		e.aLock.Unlock()

		asq.wchan <- plug
	}

}
