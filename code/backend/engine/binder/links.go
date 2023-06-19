package binder

import (
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/engine/invokers/forked"
	"github.com/temphia/temphia/code/backend/engine/invokers/linked"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/job"
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

	links := make([]bindx.Link, 0, len(b.Links))

	for _, al := range b.Links {
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

	alink, ok := b.Links[name]
	if !ok {
		return nil, easyerr.Error(etypes.LinkNotFound)
	}

	out, err := data.AsJsonBytes()
	if err != nil {
		return nil, err
	}

	resp, err := b.Deps.Runtime.Preform(&job.Job{
		PlugId:      alink.ToPlug,
		AgentId:     alink.ToAgent,
		EventId:     xid.New().String(),
		EventAction: method,
		Namespace:   b.Namespace,
		Payload:     out,
		Invoker:     linked.New(b.EventId, b.PlugId, b.AgentId, nil),
	})

	if err != nil {
		return nil, err
	}

	return lazydata.NewJsonData(resp.Payload), nil

}

func (b *Binder) selfForkExec(method string, data []byte) error {

	newEventId := xid.New().String()

	b.logInfo().
		Str("new_event_id", newEventId).
		Msg(logid.BinderExecutionForked)

	j := b.Job

	go func() {
		b.Deps.Runtime.Preform(&job.Job{
			PlugId:            b.PlugId,
			AgentId:           b.AgentId,
			EventId:           newEventId,
			EventAction:       method,
			Namespace:         b.Namespace,
			Payload:           data,
			PendingPrePolicy:  true,
			PendingPostPolicy: true,
			Loaded:            j.Loaded,
			Plug:              j.Plug,
			Agent:             j.Agent,
			Invoker:           forked.New(b.EventId),
		})
	}()

	return nil

}
