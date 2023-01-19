package web2agent

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/tidwall/gjson"
)

type WAState struct {
	templates     map[string]string
	templateFuncs map[string]string
	routes        map[string]string
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

	for defkey, def := range def.Get("templates").Map() {
		w.state.templates[defkey] = def.String()
	}

	for defkey, def := range def.Get("template_funcs").Map() {
		w.state.templateFuncs[defkey] = def.String()
	}

	for defkey, def := range def.Get("routes").Map() {
		w.state.routes[defkey] = def.String()
	}

}
