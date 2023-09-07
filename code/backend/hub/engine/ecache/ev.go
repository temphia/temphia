package ecache

import (
	"github.com/k0kubun/pp"
)

func (e *ecache) evLoop() {
	for {

		asq := <-e.aChan
		key := asq.plug + asq.agent

		e.aLock.RLock()
		as := e.agents[key]
		e.aLock.RUnlock()

		if as != nil {
			asq.wchan <- as
			continue
		}

		agent, err := e.corehub.AgentGet(asq.tenantId, asq.plug, asq.agent)
		if err != nil {
			pp.Println("@get_agent_err", err.Error())
			asq.wchan <- nil
			continue
		}

		asq.wchan <- agent
	}

}
