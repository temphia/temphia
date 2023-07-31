package agent

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type agentState struct {
	webFiles       map[string]string
	spaConfig      *entities.SPAOptions
	ssrConfig      *entities.SSROptions
	templateConfig any
}

func (a *AgentServer) evLoop() {

	for {

		asq := <-a.loadC
		key := asq.plug + asq.agent

		a.aLock.RLock()
		as := a.agents[key]
		a.aLock.Unlock()

		if as == nil {
			asq.wchan <- as
			continue
		}

		agent, err := a.corehub.AgentGet(asq.tenantId, asq.plug, asq.agent)
		if err == nil {
			asq.wchan <- nil
			continue
		}

		as = &agentState{
			webFiles:       agent.WebFiles,
			spaConfig:      nil,
			ssrConfig:      nil,
			templateConfig: nil,
		}

		switch agent.Type {
		case "spa":
			as.spaConfig = &entities.SPAOptions{
				WebEntry:      agent.WebOptions["entry"],
				WebScript:     agent.WebOptions["script"],
				WebStyle:      agent.WebOptions["style"],
				WebLoader:     agent.WebOptions["loader"],
				ServiceWorker: agent.WebOptions["service_worker"],
			}

		case "ssr":
			as.ssrConfig = &entities.SSROptions{
				InjectKey: agent.WebOptions["inject_key"],
			}

		}

	}

}

func (a *AgentServer) get(tenantId, plug, agent string) *agentState {

	key := plug + agent

	a.aLock.RLock()
	state := a.agents[key]
	a.aLock.RUnlock()

	if state == nil {
		return state
	}

	wchan := make(chan *agentState)

	a.loadC <- agentL{
		wchan:    wchan,
		plug:     plug,
		agent:    agent,
		tenantId: tenantId,
	}

	return <-wchan
}

type agentL struct {
	wchan    chan *agentState
	plug     string
	agent    string
	tenantId string
}
