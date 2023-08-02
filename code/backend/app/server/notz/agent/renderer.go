package agent

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz"
)

func (a *AgentNotz) spaRender(ctx xnotz.Context, state *agentState) {

	builder := SpaBuilder{
		opts: SpaBuilderOptions{
			Plug:         ctx.PlugId,
			Agent:        ctx.AgentId,
			APIBaseURL:   httpx.ApiBaseURL(ctx.Request.Host, ctx.TenantId),
			EntryName:    state.spaConfig.WebEntry,
			ExecLoader:   "",
			TenantID:     ctx.TenantId,
			JsPlugScript: state.spaConfig.WebScript,
			StyleFile:    state.spaConfig.WebStyle,
			ExtScripts:   make(map[string]interface{}),
		},
	}

	out := builder.Build()

	ctx.Writer.Write(out)
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "text/html")

}
