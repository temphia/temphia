package web2agent

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type TemplateContext struct {
	Rid        int64          `json:"rid,omitempty"`
	Path       string         `json:"path,omitempty"`
	DomainName string         `json:"domain_name,omitempty"`
	Data       map[string]any `json:"data,omitempty"`
}

type Request struct {
	Rid  int64  `json:"rid,omitempty"`
	Path string `json:"path,omitempty"`
}

func (w *WATarget) serveTemplate(file string) {
	tpl := w.adapter.state.template
	if tpl == nil {
		return
	}

	ctx := TemplateContext{
		Rid:        w.rid,
		Path:       w.http.Request.URL.Path,
		DomainName: w.adapter.domain.Name,
		Data:       map[string]any{},
	}

	if action, ok := w.adapter.state.routes[file]; ok {

		req := Request{
			Rid:  w.rid,
			Path: w.http.Request.URL.Path,
		}

		rout, err := json.Marshal(&req)
		if err != nil {
			pp.Println("request unmarshel err", err.Error())
			return
		}

		out, err := w.adapter.engine.Execute(etypes.Execution{
			TenantId: w.adapter.tenantId,
			PlugId:   w.adapter.mainHook.PlugId,
			AgentId:  w.adapter.mainHook.AgentId,
			Action:   action,
			Payload:  rout,
			Invoker:  w,
		})

		if err != nil {
			pp.Println("could not laod conext data from action", err.Error())
		}

		json.Unmarshal(out, &ctx.Data)

	}

	tpl.Execute(w.http.Writer, ctx)

}
