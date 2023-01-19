package web2agent

import (
	"encoding/json"
	"text/template"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

type TemplateContext struct {
	Rid        int64  `json:"rid,omitempty"`
	Path       string `json:"path,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
	Data       any    `json:"data,omitempty"`
}

type Request struct {
	Rid  int64  `json:"rid,omitempty"`
	Path string `json:"path,omitempty"`
}

func (w *WATarget) serveTemplate(file string) {
	if w.adapter.state.template == nil {
		return
	}

	ctx := TemplateContext{
		Rid:        w.rid,
		Path:       w.http.Request.URL.Path,
		DomainName: w.adapter.domain.Name,
		Data:       nil,
	}

	if action, ok := w.adapter.state.routes[file]; ok {
		aresp, err := w.performAction(action)
		if err != nil {
			ctx.Data = aresp
		}
	}

	tpl, err := w.adapter.state.template.Clone()
	if err != nil {
		pp.Println("err when cloning")
		return
	}

	tpl.Funcs(template.FuncMap{
		"action": w.performAction,
	}).Execute(w.http.Writer, ctx)

}

func (w *WATarget) performAction(action string) (any, error) {

	var data any

	if action, ok := w.adapter.state.routes[action]; ok {
		req := Request{
			Rid:  w.rid,
			Path: w.http.Request.URL.Path,
		}

		rout, err := json.Marshal(&req)
		if err != nil {
			return nil, err
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
			return nil, err
		}

		err = json.Unmarshal(out, &data)
		if err != nil {
			return nil, err
		}
	}

	return data, nil

}
