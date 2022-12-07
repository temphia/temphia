package invokers

import (
	"github.com/temphia/temphia/code/core/backend/xtypes"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
)

type Handle struct {
	invoker *Invoker
}

func (h *Handle) Parse(tenantId, token string) (*claim.Session, error) {
	signer := h.invoker.app.GetDeps().Signer().(service.Signer)

	return signer.ParseSession(tenantId, token)
}

func (h *Handle) ExecuteModule(module, action string, data xtypes.LazyData) (xtypes.LazyData, error) {
	return h.invoker.ExecuteModule(module, action, data)
}

func (h *Handle) ListModule() []string {
	return h.invoker.ListModules()
}

func (h *Handle) GetApp() interface{} {
	return h.invoker.app
}
