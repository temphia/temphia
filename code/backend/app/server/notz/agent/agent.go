package agent

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz"
)

type AgentNotz struct {
	ehub    etypes.EngineHub
	corehub store.CoreHub
	cabinet store.CabinetHub
	ecache  etypes.Ecache
}

func New(ehub etypes.EngineHub, corehub store.CoreHub, cabinet store.CabinetHub) *AgentNotz {
	as := &AgentNotz{
		corehub: corehub,
		cabinet: cabinet,
		ehub:    ehub,
	}

	return as
}

func (a *AgentNotz) Start() error {
	ecahe := a.ehub.GetCache()
	if ecahe == nil {
		return easyerr.Error("ecache not found")
	}

	a.ecache = ecahe

	return nil
}

func (a *AgentNotz) Render(ctx xnotz.Context) {
	as := a.ecache.GetAgent(ctx.TenantId, ctx.PlugId, ctx.AgentId)
	if as == nil {
		return
	}

	a.spaRender(ctx, as)
}
