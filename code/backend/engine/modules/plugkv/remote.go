package plugkv

import (
	"encoding/json"
	"errors"

	"github.com/temphia/temphia/code/backend/engine/invokers/linked"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

var ErrNoLink = errors.New("No Link found")

type RCAgent struct {
	tenantId string
	agentId  string
	plugId   string

	corehub store.CoreHub
	loaded  bool
	links   []*entities.AgentLink
	engine  etypes.Engine
}

func NewRCAgent(tenantId, plugId, agentId string, corehub store.CoreHub) *RCAgent {
	return &RCAgent{
		tenantId: tenantId,
		agentId:  agentId,
		plugId:   plugId,
		corehub:  corehub,
		loaded:   false,
	}

}

type emitOptions struct {
	Name  string          `json:"name,omitempty"`
	Data  json.RawMessage `json:"data,omitempty"`
	Async bool            `json:"async,omitempty"`
}

func (r *RCAgent) EmitBatch(opts *emitOptions) ([]any, error) {
	if !r.loaded {
		err := r.load()
		if err != nil {
			return nil, err
		}
	}

	for _, link := range r.links {

		go func(link *entities.AgentLink) {
			r.emit(opts.Name, opts.Data, link)

		}(link)
	}

	return nil, nil
}

func (r *RCAgent) EmitBatchAsync(opts *emitOptions) error {
	if !r.loaded {
		err := r.load()
		if err != nil {
			return err
		}
	}

	for _, link := range r.links {
		go func(link *entities.AgentLink) {
			r.emit(opts.Name, opts.Data, link)
		}(link)
	}

	return nil
}

func (r *RCAgent) EmitOne(opts *emitOptions) (any, error) {
	if !r.loaded {
		err := r.load()
		if err != nil {
			return nil, err
		}
	}

	return r.emit(opts.Name, opts.Data, r.links[0])
}

func (r *RCAgent) EmitOneAysnc(opts *emitOptions) error {
	if !r.loaded {
		err := r.load()
		if err != nil {
			return err
		}
	}

	go r.emit(opts.Name, opts.Data, r.links[0])

	return nil
}

// private

func (r *RCAgent) emit(method string, data []byte, link *entities.AgentLink) (any, error) {
	return r.engine.Execute(etypes.Execution{
		TenantId: r.tenantId,
		PlugId:   link.ToPlug,
		AgentId:  link.ToAgent,
		Action:   method,
		Payload:  data,
		Invoker:  linked.New("", r.agentId, r.plugId, nil),
	})
}

func (r *RCAgent) load() error {
	links, err := r.corehub.AgentLinkList(r.tenantId, r.plugId, r.agentId)
	if err != nil {
		return err
	}

	if len(links) == 0 {
		return ErrNoLink
	}

	r.links = links
	r.loaded = true
	return nil
}
