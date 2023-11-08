package notz

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/server/API/engine/notz/spatpl"
	"github.com/temphia/temphia/code/backend/app/server/API/engine/router"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

const (
	NotzRendererDynamicSPA = "notz.dspa"
	NotzRendererStandardV1 = "notz.std.v1"
)

func (a *Notz) HandleAgent(ctx xnotz.Context) {
	as := a.ecache.GetAgent(ctx.TenantId, ctx.PlugId, ctx.AgentId)
	if as == nil {
		return
	}

	switch as.Renderer {
	case NotzRendererStandardV1, "":
		a.stdRendererV1(ctx, as)
	case NotzRendererDynamicSPA:
		a.dynamicSPARender(ctx, as)
	default:
		panic("not implemented")
	}

}

func (a *Notz) dynamicSPARender(ctx xnotz.Context, agent *entities.Agent) {

	builder := spatpl.New(spatpl.SpaBuilderOptions{
		Plug:         ctx.PlugId,
		Agent:        ctx.AgentId,
		TenantID:     ctx.TenantId,
		APIBaseURL:   httpx.ApiBaseURL(ctx.Request.Host, ctx.TenantId),
		EntryName:    agent.WebOptions["web_entry"],
		ExecLoader:   agent.WebOptions["exec_loader"],
		JsPlugScript: agent.WebOptions["web_script"],
		StyleFile:    agent.WebOptions["web_style"],
		ExtScripts:   make(map[string]interface{}),
	})

	out := builder.Build()

	ctx.Writer.Write(out)
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "text/html")

}

func (a *Notz) stdRendererV1(ctx xnotz.Context, agent *entities.Agent) {

	path := ctx.Request.URL.Path

	cconf := a.getRouteConfig(ctx.TenantId, ctx.PlugId, ctx.AgentId)
	if cconf == nil || cconf.config == nil {
		pp.Println("@bailing_out_conf_not found", cconf)
		return
	}

	engine := a.ehub.GetEngine()

	for _, item := range cconf.config.Items {

		if strings.HasPrefix(path, item.Path) {
			continue
		}

		switch item.Mode {
		case router.RouteItemModeRaw:

			engine.WebRawXecute(etypes.WebRawXecuteOptions{
				TenantId: ctx.TenantId,
				PlugId:   ctx.PlugId,
				AgentId:  ctx.AgentId,
				Writer:   ctx.Writer,
				Request:  ctx.Request,
			})

		case router.RouteItemModeServe:
			a.serve(ctx, cconf.bprintId)

		case router.RouteItemModeRPX:

			engine.RPXecute(etypes.RPXecuteOptions{
				TenantId: ctx.TenantId,
				PlugId:   ctx.PlugId,
				AgentId:  ctx.AgentId,
				Action:   item.Target,
				Payload:  nil,
				Invoker:  nil,
			})

		default:
			panic("Unknow Renderer")
		}

	}

}

func (a *Notz) serve(ctx xnotz.Context, bprintid string) {

	// fprefix, file := static.ExtractPath(path, agent)

	file := ""
	fprefix := ""

	folder := fmt.Sprintf("%s/%s/%s", xtypes.BprintBlobFolder, bprintid, fprefix)

	out, err := a.cabinet.GetBlob(ctx.Request.Context(), ctx.TenantId, folder, file)
	if err != nil {
		return
	}

	ffiles := strings.Split(file, ".")

	ctype := ""
	switch ffiles[1] {
	case "js":
		ctype = httpx.CtypeJS
	case "css":
		ctype = httpx.CtypeCSS
	default:
		ctype = http.DetectContentType(out)
	}

	ctx.Writer.Header().Set("Context-Type", ctype)
	ctx.Writer.Write(out)

}
