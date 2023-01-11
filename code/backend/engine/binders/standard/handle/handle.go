package handle

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/core/backend/engine/binders/standard/deps"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

// handle is a shared ctx between
// different bindings component
type Handle struct {
	Deps      *deps.Deps
	Namespace string
	PlugId    string
	AgentId   string
	BprintId  string

	Context  context.Context
	Executor etypes.Executor
	Logger   zerolog.Logger

	Job *job.Job

	EventId string
	Resp    []byte

	// lazy loaded
	Resources map[string]*entities.Resource
	Links     map[string]*entities.AgentLink
}

func New(ns, pid, aid, bid string, deps *deps.Deps) Handle {
	return Handle{
		Deps:      deps,
		Namespace: ns,
		PlugId:    pid,
		AgentId:   aid,
		BprintId:  bid,
	}
}

func (h *Handle) InitLogger() {

	h.Logger = h.Deps.LoggerBase.
		With().
		Str("tenant_id", h.Namespace).
		Str("plug_id", h.PlugId).
		Str("agent_id", h.AgentId).
		Str("bprint_id", h.BprintId).
		Str("event_id", h.EventId).Logger()
}

func (h *Handle) LoadResources() {
	if h.Resources != nil {
		return
	}

	corehub := h.Deps.App.GetDeps().CoreHub().(store.CoreHub)
	ress, err := corehub.ResourceListByAgent(h.Namespace, h.PlugId, h.AgentId)
	if err != nil {
		panic(err)
	}

	h.Resources = make(map[string]*entities.Resource, len(ress))
	for _, r := range ress {
		h.Resources[r.Name] = r
	}
}

func (h *Handle) LoadLinks() {
	if h.Links != nil {
		return
	}

	corehub := h.Deps.App.GetDeps().CoreHub().(store.CoreHub)
	links, err := corehub.AgentLinkList(h.Namespace, h.PlugId, h.AgentId)
	if err != nil {
		panic(err)
	}

	h.Links = make(map[string]*entities.AgentLink, len(links))

	for _, al := range links {
		h.Links[al.Name] = al
	}

}
