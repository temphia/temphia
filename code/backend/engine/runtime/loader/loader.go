package loader

import (
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/models/vmodels"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func Load(syncer store.SyncDB, tenantId, plugId, agentId string) (*vmodels.ExecutorData, error) {
	plug, err := syncer.PlugGet(tenantId, plugId)
	if err != nil {
		pp.Println("err@plug")
		return nil, err
	}

	agent, err := syncer.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		pp.Println("err@agent")

		return nil, err
	}

	pp.Println("AGENT & PLUG loaded")

	// ignoring bprint err so bprint might be nil
	// we have to check that in binder
	bprint, _ := syncer.BprintGet(tenantId, plug.BprintId)

	return &vmodels.ExecutorData{
		Plug:   plug,
		Agent:  agent,
		Bprint: bprint,
	}, nil

}
