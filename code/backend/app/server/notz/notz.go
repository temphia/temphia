package notz

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Notz struct {
	ehub    etypes.EngineHub
	corehub store.CoreHub
	cabinet store.CabinetHub
	ecache  etypes.Ecache
}

func New(app xtypes.App) *Notz {
	deps := app.GetDeps()

	ehub := deps.EngineHub().(etypes.EngineHub)
	corehub := deps.CoreHub().(store.CoreHub)
	cabinet := deps.Cabinet().(store.CabinetHub)

	return &Notz{
		ehub:    ehub,
		corehub: corehub,
		cabinet: cabinet,
	}
}

func (n *Notz) Start() error {
	ecahe := n.ehub.GetCache()
	if ecahe == nil {
		return easyerr.Error("ecache not found")
	}

	n.ecache = ecahe

	return nil
}
