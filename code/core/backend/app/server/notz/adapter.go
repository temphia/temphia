package notz

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func (m *Notz) get(tenantId, host string) httpx.Adapter {
	m.rlock.Lock()
	r, ok := m.renderers[tenantId+host]
	if ok {
		return r
	}

	m.rlock.Unlock()

	corehub := m.app.GetDeps().CoreHub().(store.CoreHub)

	td, err := corehub.GetDomainByName(tenantId, host)
	if err != nil {
		return nil
	}

	rbuilder, ok := m.rendererBuilders[td.AdapterType]
	if !ok {
		return nil
	}

	renderer, err := rbuilder(httpx.BuilderOptions{
		App:      m.app,
		TenantId: tenantId,
		Domain:   td,
	})
	if err != nil {
		return nil
	}

	m.rlock.Lock()
	r, ok = m.renderers[tenantId+host]
	if ok {
		return r
	} else {
		m.renderers[tenantId+host] = renderer
	}
	m.rlock.Unlock()
	return renderer
}
