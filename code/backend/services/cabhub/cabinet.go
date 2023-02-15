package cabhub

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

type CabHub struct {
	defaultProvider store.CabinetSource
	sources         map[string]store.CabinetSource
	defName         string
}

func New(sources map[string]store.CabinetSource, defprovider string) *CabHub {
	ch := &CabHub{
		sources:         sources,
		defaultProvider: sources[defprovider],
		defName:         defprovider,
	}
	return ch
}

func (c *CabHub) Start(eventbus any) error {
	eb := eventbus.(xplane.EventBus)

	eb.OnTenantChange(func(tenant, event string, data *entities.Tenant) {
		go func() {
			switch event {
			case xplane.EventCreateTenant:
				c.defaultProvider.InitilizeTenent(tenant, store.DefaultFolders)
			default:
				pp.Println("skipping event")
			}
		}()
	})

	return nil
}

func (c *CabHub) Default(tenant string) store.CabinetSourced {
	return &cabinetSourced{
		source:   c.defName,
		tenantId: tenant,
		provider: c.defaultProvider,
	}
}

func (c *CabHub) ListSources(tenant string) ([]string, error) {
	sources := make([]string, 0, len(c.sources))
	for k := range c.sources {
		sources = append(sources, k)
	}
	return sources, nil
}

func (c *CabHub) GetSource(source, tenant string) store.CabinetSourced {

	pp.Println(source, tenant)
	fmt.Println("@sources", c.sources)

	provider, ok := c.sources[source]
	if !ok {
		panic(easyerr.NotFound())
	}

	return &cabinetSourced{
		source:   source,
		tenantId: tenant,
		provider: provider,
	}
}

func (c *CabHub) DefaultName(tenant string) string { return c.defName }
