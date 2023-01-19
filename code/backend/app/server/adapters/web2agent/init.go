package web2agent

import (
	"fmt"
	"html/template"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/server/adapters/common"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/tidwall/gjson"
)

type WAState struct {
	templates     map[string]string
	templateFuncs map[string]string
	routes        map[string]string
	template      *template.Template
}

func (w *Web2Agent) init() {
	deps := w.app.GetDeps()
	ch := deps.CoreHub().(store.CoreHub)
	rhub := deps.RepoHub().(repox.Hub)

	hooks, err := ch.ListTargetHookByType(w.tenantId, entities.TargetHookTypeDomainHook, fmt.Sprintf("%d", w.domain.Id))
	if err != nil {
		w.initError = err.Error()
		return
	}

	if len(hooks) == 0 {
		w.initError = "hooks empty"
		return
	}

	w.mainHook = hooks[0]
	handlePlug := w.mainHook.PlugId
	handleAgent := w.mainHook.AgentId

	plug, err := ch.PlugGet(w.tenantId, handlePlug)
	if err != nil {
		w.initError = "plug get err :" + err.Error()
		return
	}

	agent, err := ch.AgentGet(w.tenantId, handlePlug, handleAgent)
	if err != nil {
		w.initError = "agent get err :" + err.Error()
		return
	}

	w.intOk = true

	out, err := rhub.BprintGetBlob(w.tenantId, plug.BprintId, agent.IfaceFile)
	if err != nil {
		return
	}

	def := gjson.GetBytes(out, "definations.web2agent")

	templateFiles := make([]string, 0)

	for defkey, def := range def.Get("templates").Map() {
		tmplFile := def.String()

		w.state.templates[defkey] = tmplFile
		templateFiles = append(templateFiles, tmplFile)
	}

	for defkey, def := range def.Get("template_funcs").Map() {
		w.state.templateFuncs[defkey] = def.String()
	}

	for defkey, def := range def.Get("routes").Map() {
		w.state.routes[defkey] = def.String()
	}

	fs := common.NewLazyFS(common.LazyFSOptions{
		Tenant: w.tenantId,
		Files:  make(map[string]struct{}),
		Handler: func(tenantId, file string) ([]byte, error) {
			return []byte(`<h1>todo</h1>`), nil
		},
	})

	t, err := template.New("web2agent").
		ParseFS(fs, templateFiles...)
	if err != nil {
		pp.Println("@paring templates error")
		return
	}

	w.state.template = t
}
