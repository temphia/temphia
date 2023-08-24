package notz

import (
	"github.com/temphia/temphia/code/backend/app/server/notz/agent"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz"
)

type Notz struct {
	agent *agent.AgentNotz
}

func New(app xtypes.App) *Notz {
	deps := app.GetDeps()

	ehub := deps.EngineHub().(etypes.EngineHub)
	corehub := deps.CoreHub().(store.CoreHub)
	cabinet := deps.Cabinet().(store.CabinetHub)

	return &Notz{
		agent: agent.New(ehub, corehub, cabinet),
	}
}

func (n *Notz) Start() error {
	return n.agent.Start()
}

func (n *Notz) HandleAgent(ctx xnotz.Context) {
	n.agent.Render(ctx)
}

func (n *Notz) HandleDomain(ctx xnotz.Context) {

}
