package notz

import (
	"fmt"
	"io"
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

const CacheBudget = 1 << 20

const (
	NotzRendererDynamicSPA = "dspa.v1"
	NotzRendererRawHttp    = "raw_http.v1"
	NotzRendererServe      = "serve.v1"
	NotzRendererRouter     = "router.v1"
)

func (a *Notz) HandleAgent(ctx xnotz.Context) {
	as := a.ecache.GetAgent(ctx.TenantId, ctx.PlugId, ctx.AgentId)
	if as == nil {
		pp.Println("@agent_not_found", ctx.AgentId, ctx.PlugId)
		return
	}

	switch as.Renderer {
	case NotzRendererRouter:
		a.simpleRouterRendererV1(ctx, as)
	case NotzRendererDynamicSPA:
		a.dynamicSPARenderV1(ctx, as)
	case NotzRendererRawHttp, "":
		a.engine.WebRawXecute(etypes.WebRawXecuteOptions{
			TenantId: ctx.TenantId,
			PlugId:   ctx.PlugId,
			AgentId:  ctx.AgentId,
			Writer:   ctx.Writer,
			Request:  ctx.Request,
		})

	case NotzRendererServe:

		plug := a.ecache.GetPlug(ctx.TenantId, ctx.PlugId)

		path := ctx.Request.URL.Path
		subfolder := as.WebOptions["serve_folder"]
		file := ""

		if path[len(path)-1] == '/' {
			file = "index.html"
		} else {
			paths := strings.Split(path, "/")
			file = paths[len(paths)-1]

			if !strings.Contains(file, ".") && as.WebOptions["serve_append_html"] == "true" {
				file = fmt.Sprintf("%s.html", file)
			}

		}

		a.serveFromBprint1(ctx, &router.RouteResponse{
			Mode:   router.RouteItemModeServe,
			Target: subfolder,
			File:   file,
		}, cconf.bprintId)

	default:
		panic("not implemented")
	}

}

func (a *Notz) dynamicSPARenderV1(ctx xnotz.Context, agent *entities.Agent) {

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

func (a *Notz) simpleRouterRendererV1(ctx xnotz.Context, agent *entities.Agent) {

	path := ctx.Request.URL.Path

	cconf := a.getRouteConfig(ctx.TenantId, ctx.PlugId, ctx.AgentId)
	if cconf == nil || cconf.config == nil {
		pp.Println("@bailing_out_conf_not found", cconf)
		return
	}

	route := router.SimpleRoutePick(cconf.config, path, ctx.Request.Method)
	if route == nil {
		return
	}

	switch route.Mode {
	case router.RouteItemModeRPX:
		// fixme check auth token ?

		out, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			pp.Println("err reading rpx payload", err)
			return
		}

		a.engine.RPXecute(etypes.RPXecuteOptions{
			TenantId: ctx.TenantId,
			PlugId:   ctx.PlugId,
			AgentId:  ctx.AgentId,
			Action:   route.Target,
			Payload:  out,
			Invoker:  nil,
		})
	case router.RouteItemModeServe:
		a.serveFromBprint1(ctx, route, cconf.bprintId)
	case router.RouteItemModeRaw:
		a.engine.WebRawXecute(etypes.WebRawXecuteOptions{
			TenantId: ctx.TenantId,
			PlugId:   ctx.PlugId,
			AgentId:  ctx.AgentId,
			Writer:   ctx.Writer,
			Request:  ctx.Request,
		})

	default:
		panic("Unknown Renderer")
	}

}

func (a *Notz) serveFromBprint1(ctx xnotz.Context, route *router.RouteResponse, bprintid string) {

	// fixme => implement caching with cache budget

	folder := fmt.Sprintf("%s/%s/%s", xtypes.BprintBlobFolder, bprintid, route.Target)

	out, err := a.cabinet.GetBlob(ctx.Request.Context(), ctx.TenantId, folder, route.File)
	if err != nil {
		return
	}

	ffiles := strings.Split(route.File, ".")

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
