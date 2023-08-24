package ecache

func (e *ecache) evLoop() {

	for {

		asq := <-e.aChan
		key := asq.plug + asq.agent

		e.aLock.RLock()
		as := e.agents[key]
		e.aLock.Unlock()

		if as == nil {
			asq.wchan <- as
			continue
		}

		agent, err := e.corehub.AgentGet(asq.tenantId, asq.plug, asq.agent)
		if err != nil {
			asq.wchan <- nil
			continue
		}

		asq.wchan <- agent
	}

}
