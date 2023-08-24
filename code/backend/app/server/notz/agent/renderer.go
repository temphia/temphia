package agent

import (
	"net/http"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz/httpx"
)

/*

renderer_type
	era => Executor Rendered App
	spa => Single Page Application
	gossr => golang Server Side Rendered
	static => Static

*/

func (a *AgentNotz) spaRender(ctx xnotz.Context, agent *entities.Agent) {

	builder := SpaBuilder{
		opts: SpaBuilderOptions{
			Plug:         ctx.PlugId,
			Agent:        ctx.AgentId,
			TenantID:     ctx.TenantId,
			APIBaseURL:   httpx.ApiBaseURL(ctx.Request.Host, ctx.TenantId),
			EntryName:    agent.WebOptions["web_entry"],
			ExecLoader:   agent.WebOptions["exec_loader"],
			JsPlugScript: agent.WebOptions["web_script"],
			StyleFile:    agent.WebOptions["web_style"],
			ExtScripts:   make(map[string]interface{}),
		},
	}

	out := builder.Build()

	ctx.Writer.Write(out)
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "text/html")

}
