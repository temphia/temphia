package binder

import (
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/engine/invokers/linked"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
)

func (b *SelfBindings) selfInLinks() ([]bindx.Link, error) {
	resp, err := b.db.AgentLinkListReverse(b.handle.Namespace, b.handle.PlugId, b.handle.AgentId)
	if err != nil {
		panic(err)
	}

	links := make([]bindx.Link, 0, len(resp))

	for _, al := range resp {
		links = append(links, bindx.Link{
			Name:    al.Name,
			Type:    "",
			PlugId:  al.FromPlug,
			AgentId: al.FromAgent,
		})
	}

	return links, nil
}

func (b *SelfBindings) selfOutLinks() ([]bindx.Link, error) {
	b.handle.LoadLinks()

	links := make([]bindx.Link, 0, len(b.handle.Links))

	for _, al := range b.handle.Links {
		links = append(links, bindx.Link{
			Name:    al.Name,
			Type:    "",
			PlugId:  al.ToPlug,
			AgentId: al.ToAgent,
		})
	}

	return links, nil
}

func (b *SelfBindings) selfLinkExec(name, method string, data xtypes.LazyData, async, detached bool) (xtypes.LazyData, error) {

	alink, ok := b.handle.Links[name]
	if !ok {
		return nil, easyerr.Error(etypes.LinkNotFound)
	}

	out, err := data.AsJsonBytes()
	if err != nil {
		return nil, err
	}

	if async {
		resp, err := b.runtime.Preform(&job.Job{
			PlugId:      alink.ToPlug,
			AgentId:     alink.ToAgent,
			EventId:     xid.New().String(),
			EventAction: method,
			Namespace:   b.handle.Namespace,
			Payload:     out,
			Invoker:     linked.New(b.handle.EventId, b.handle.PlugId, b.handle.AgentId, nil),
		})

		if err != nil {
			return nil, err
		}

		return lazydata.NewJsonData(resp.Payload), nil
	}

	if detached {

		return nil, nil
	}

	panic("Not implemented")
}
