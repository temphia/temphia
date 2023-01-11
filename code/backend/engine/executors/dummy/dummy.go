package dummy

import (
	"encoding/json"
	"time"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"
)

var _ etypes.Executor = (*Dummy)(nil)

func New(opts etypes.ExecutorOption) (etypes.Executor, error) {
	return &Dummy{
		Binder:   opts.Binder,
		PlugId:   opts.PlugId,
		AgentId:  opts.AgentId,
		ExecType: opts.ExecType,
	}, nil
}

type Dummy struct {
	Binder   bindx.Bindings
	PlugId   string
	AgentId  string
	Slug     string
	ExecType string
}

type PingResp struct {
	Action  string          `json:"action,omitempty"`
	Time    time.Time       `json:"time,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (d *Dummy) Process(ev *event.Request) (*event.Response, error) {
	var out []byte
	switch ev.Name {
	case "ping":
		out1, err := d.actionPing(ev)
		if err != nil {
			return nil, err
		}
		out = out1
	default:
		return nil, easyerr.Error("action not found")
	}

	return &event.Response{
		Payload: (out),
	}, nil
}

func (d *Dummy) actionPing(ev *event.Request) ([]byte, error) {
	ping := PingResp{
		Action:  ev.Name,
		Time:    time.Now(),
		Message: "hello",
		Data:    nil,
	}

	return json.Marshal(&ping)
}
