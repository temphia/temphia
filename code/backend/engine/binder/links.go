package binder

import (
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
)

func (b *Binder) selfInLinks() ([]bindx.Link, error) {

	resp, err := b.Deps.Corehub.AgentLinkListReverse(b.Namespace, b.PlugId, b.AgentId)
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

func (b *Binder) selfOutLinks() ([]bindx.Link, error) {
	b.loadLinks()

	links := make([]bindx.Link, 0, len(b.links))

	for _, al := range b.links {
		links = append(links, bindx.Link{
			Name:    al.Name,
			Type:    "",
			PlugId:  al.ToPlug,
			AgentId: al.ToAgent,
		})
	}

	return links, nil
}

func (b *Binder) selfLinkExec(name, method string, data xtypes.LazyData) (xtypes.LazyData, error) {

	alink, ok := b.links[name]
	if !ok {
		return nil, easyerr.Error(etypes.LinkNotFound)
	}

	out, err := data.AsJsonBytes()
	if err != nil {
		return nil, err
	}

	resp, err := b.Deps.Engine.RPXecute(etypes.Execution{
		TenantId: b.Namespace,
		PlugId:   alink.ToPlug,
		AgentId:  alink.ToAgent,
		Action:   method,
		Payload:  out,
		Invoker:  nil, // fixme
	})

	if err != nil {
		return nil, err
	}

	return lazydata.NewJsonData(resp), nil

}

func (b *Binder) selfForkExec(eid string, method string, data []byte) error {

	newEventId := xid.New().String()

	b.logInfo().
		Str("new_event_id", newEventId).
		Msg(logid.BinderExecutionForked)

	go func() {

		b.Deps.Engine.RPXecute(etypes.Execution{
			TenantId: b.Namespace,
			PlugId:   b.PlugId,
			AgentId:  b.AgentId,
			Action:   method,
			Payload:  data,
			Invoker:  nil, // fixme
		})

	}()

	return nil

}
