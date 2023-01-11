package datahub

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
)

/*
	hooks types
		- on_row_ctx
		- on_table_ctx
		- on_before_mod
		- on_after_mod
*/

type DataEventReq struct {
	Id    int64          `json:"id,omitempty"`
	Type  string         `json:"type,omitempty"`
	Group string         `json:"group,omitempty"`
	Table string         `json:"table,omitempty"`
	Data  map[string]any `json:"data,omitempty"`
}

func (d *dynSource) OnBeforeMod(tenant, plug, agent, handler string, event DataEventReq) (map[string]any, error) {

	out, err := json.Marshal(&event)
	if err != nil {
		return nil, err
	}

	j := &job.Job{
		PlugId:            plug,
		AgentId:           agent,
		Namespace:         tenant,
		EventId:           xid.New().String(),
		EventAction:       handler,
		Payload:           (out),
		PendingPrePolicy:  true,
		PendingPostPolicy: true,
		Invoker:           nil, // FIXME => ?
	}

	runtime := d.hub.engine.GetRuntime()

	resp, err := runtime.Preform(j)
	if err != nil {
		return nil, err
	}

	pp.Println(resp)

	return nil, nil
}

func (d *dynSource) OnAfterMod(event DataEventReq) error {
	return nil
}
