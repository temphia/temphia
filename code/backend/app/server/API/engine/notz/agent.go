package notz

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/k0kubun/pp"

	"github.com/temphia/temphia/code/backend/app/server/API/engine/notz/spatpl"
	"github.com/temphia/temphia/code/backend/app/server/API/engine/notz/static"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (a *Notz) HandleAgent(ctx xnotz.Context) {
	as := a.ecache.GetAgent(ctx.TenantId, ctx.PlugId, ctx.AgentId)
	if as == nil {
		return
	}

	pp.Println("@got_agent", as)

	switch as.Renderer {
	case "era":
		eb := a.ehub.GetExecutorBuilder(as.Executor)
		if eb == nil {
			ctx.Writer.Write([]byte(`<h1>Executor not found</h1>`))
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			return
		}

		pp.Println("@era", eb)

		// er, ok := eb.(etypes.ExecutorRenderer)
		// if !ok {
		// 	ctx.Writer.Write([]byte(`<h1>Executor is not renderer.</h1>`))
		// 	ctx.Writer.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// er.Handle(&etypes.ERContext{
		// 	Writer:  ctx.Writer,
		// 	Request: ctx.Request,
		// 	PlugId:  ctx.PlugId,
		// 	AgentId: ctx.AgentId,
		// })
	case "static":
		a.staticRenderer(ctx, as)
	case "spa":
		a.spaRender(ctx, as)
	case "raw":
		a.rawRender(ctx, as)
	default:
		panic("not implemented")
	}

}

/*

renderer_type
	era => Executor Rendered App
	spa => Single Page Application
	gossr => golang Server Side Rendered
	static => Static

*/

func (a *Notz) spaRender(ctx xnotz.Context, agent *entities.Agent) {

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

func (a *Notz) staticRenderer(ctx xnotz.Context, agent *entities.Agent) {

	plug := a.ecache.GetPlug(ctx.TenantId, ctx.PlugId)
	bprintid := plug.BprintId

	path := ctx.Request.URL.Path

	fprefix, file := static.ExtractPath(path, agent)

	folder := fmt.Sprintf("%s/%s/%s", xtypes.BprintBlobFolder, bprintid, fprefix)

	pp.Println("@folder/file", folder, file)

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

func (a *Notz) rawRender(ctx xnotz.Context, agent *entities.Agent) {

	key := ctx.PlugId + ctx.AgentId

	a.rLock.RLock()
	rpxy := a.rawProxies[key]
	a.rLock.RUnlock()

	if rpxy == nil {
		a.laLock.Lock()
		addr := a.laddrs[key]
		a.laLock.Unlock()

		if addr == "" {
			return
		}

		u, err := url.Parse(fmt.Sprintf("http://%s", addr))
		if err != nil {
			return
		}

		rpxy = httputil.NewSingleHostReverseProxy(u)

		a.rLock.Lock()
		if _, ok := a.rawProxies[key]; !ok {
			a.rawProxies[key] = rpxy
		}
		a.rLock.Unlock()
	}

	rpxy.ServeHTTP(ctx.Writer, ctx.Request)

}
