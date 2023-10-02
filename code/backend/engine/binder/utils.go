package binder

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

func (h *Binder) initLogger() {

	h.Logger = h.Deps.LoggerBase.
		With().
		Str("tenant_id", h.Namespace).
		Str("plug_id", h.PlugId).
		Str("agent_id", h.AgentId).
		Str("bprint_id", h.BprintId).
		Str("event_id", "fixme").Logger()
}

func (b *Binder) logInfo() *zerolog.Event {
	return b.Logger.Info()
}

func (b *Binder) logErr() *zerolog.Event {
	return b.Logger.Info()
}

func (b *Binder) logDebug() *zerolog.Event {
	return b.Logger.Info()
}

func (b *Binder) logDebugRoom(msg *etypes.DebugMessage) {
	out, err := json.Marshal(msg)
	if err != nil {
		pp.Println(err)
		return
	}

	b.Deps.Sockd.SendBroadcast(b.Namespace, "plugs_dev", []int64{}, out)
}

func (h *Binder) loadResources() {
	if h.resources != nil {
		return
	}

	h.Logger.Info().Msg(logid.EngineResourcesLoading)

	corehub := h.Deps.App.GetDeps().CoreHub().(store.CoreHub)
	agentRes, err := corehub.ListResourcePairs(h.Namespace, h.PlugId, h.AgentId)
	if err != nil {
		panic(err)
	}

	rh := make(map[string]*entities.Resource)

	for _, rp := range agentRes {
		// fixme => overlay exta_meta|actions|policy etc from agentResource on top of resource
		rh[rp.AgentResource.Slug] = rp.Resource
	}

	h.resources = rh
	h.Logger.Info().
		Interface("resources", rh).
		Msg(logid.EngineResourcesLoaded)

}

func (h *Binder) loadLinks() {
	if h.links != nil {
		return
	}

	corehub := h.Deps.App.GetDeps().CoreHub().(store.CoreHub)
	links, err := corehub.AgentLinkList(h.Namespace, h.PlugId, h.AgentId)
	if err != nil {
		panic(err)
	}

	h.links = make(map[string]*entities.AgentLink, len(links))

	for _, al := range links {
		h.links[al.Name] = al
	}

}
