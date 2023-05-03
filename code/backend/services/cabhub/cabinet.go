package cabhub

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
	"github.com/tidwall/gjson"
)

type CabHub struct {
	defaultProvider store.CabinetSource
	sources         map[string]store.CabinetSource
	defName         string
	modChan         chan xplane.Message
}

func New(sources map[string]store.CabinetSource, defprovider string) *CabHub {
	ch := &CabHub{
		sources:         sources,
		defaultProvider: sources[defprovider],
		defName:         defprovider,
		modChan:         make(chan xplane.Message),
	}
	return ch
}

func (c *CabHub) Start(mb xplane.MsgBus) error {

	mb.Subscribe("tenant", c.modChan)

	go func() {
		msg := <-c.modChan

		switch msg.Path {
		case "":
			// fixme => path support in systemevent
			fallthrough
		case "create":
			c.defaultProvider.InitilizeTenent(gjson.Get(msg.Data, "slug").String(), store.DefaultFolders)
		default:
			pp.Println("skipping event")
		}
	}()

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
		panic(easyerr.NotFound("cabinet source"))
	}

	return &cabinetSourced{
		source:   source,
		tenantId: tenant,
		provider: provider,
	}
}

func (c *CabHub) DefaultName(tenant string) string { return c.defName }
