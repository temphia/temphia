package web2agent

import (
	"io"
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

func (w *WATarget) serveAction(action string) {
	out, _ := io.ReadAll(w.http.Request.Body)
	out, err := w.adapter.engine.Execute(etypes.Execution{
		TenantId: w.adapter.tenantId,
		PlugId:   w.adapter.mainHook.PlugId,
		AgentId:  w.adapter.mainHook.AgentId,
		Action:   action,
		Payload:  out,
		Invoker:  w,
	})

	if err != nil {
		w.error(err.Error())
		return
	}

	w.http.Data(http.StatusOK, "", out)
}
