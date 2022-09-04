package engine

import (
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/core/backend/engine/invoker/web"
	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/job"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/vmodels"

	"github.com/ztrue/tracerr"
)

func (e *Engine) clientLaunchExec(tenantId, plugId, agentId, mode string, ctx *gin.Context) {
	if mode == "ssr" {
		e.clientLaunchExecSSR(tenantId, plugId, agentId, ctx)
		return
	}

	agent, err := e.syncer.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		httpx.WriteErr(ctx, err.Error())
		return
	}

	resp := &vmodels.LoaderOptions{
		BaseURL:      baseURL(ctx),
		Token:        "",
		EntryName:    agent.WebEntry,
		ExecLoader:   agent.WebLoader,
		JSPlugScript: agent.WebScript,
		StyleFile:    agent.WebStyle,
		Plug:         plugId,
		Agent:        agentId,
		ExtScripts:   nil,
	}

	httpx.WriteJSON(ctx, resp, nil)

}

func (e *Engine) plugAction(tenantId, plugId, agentId, action string, ctx *gin.Context) {

	e.logger.Debug().Str("tenant_id", tenantId).
		Str("plug_id", plugId).
		Str("agent_id", agentId).
		Str("action", action).Msg("action start")

	out, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		httpx.WriteErr(ctx, err.Error())
		return
	}

	ictx := web.NewWeb(ctx, e.signer)

	j := &job.Job{
		PlugId:            plugId,
		AgentId:           agentId,
		Namespace:         tenantId,
		EventId:           xid.New().String(),
		EventAction:       action,
		Payload:           (out),
		PendingPrePolicy:  true,
		PendingPostPolicy: true,
		Plug:              nil,
		Agent:             nil,
		Loaded:            false,
		Invoker:           ictx,
	}

	eresp, err := e.runtime.Preform(j)
	if err != nil {
		tracerr.PrintSourceColor(err, 10)
		httpx.WriteErr(ctx, err.Error())
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.Writer.Write(eresp.Payload)
}

func (e *Engine) servePlugFile(tenantId, plugId, agentId, file string, ctx *gin.Context) {
	plug, err := e.syncer.PlugGet(tenantId, plugId)
	if err != nil {
		e.logFileServe(tenantId, plugId, agentId, file, "error reteriving plug")
		httpx.WriteErr(ctx, err.Error())
		return
	}

	agent, err := e.syncer.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		e.logFileServe(tenantId, plugId, agentId, file, "error reteriving agent")
		httpx.WriteErr(ctx, err.Error())
		return
	}

	actualFile := agent.WebFiles[file]

	if actualFile == "" {
		e.logFileServe(tenantId, plugId, agentId, file, "empty file mappings")
		return
	}

	out, err := e.pacman.BprintGetBlob(tenantId, plug.BprintId, actualFile)
	if err != nil {
		e.logFileServe(tenantId, plugId, agentId, file, "file not found in cabinet")
		httpx.WriteErr(ctx, err.Error())
		return
	}

	e.writeFile(file, out, ctx)
}

func (e *Engine) serveExecutorFile(tenantId, plugId, agentId, file string, ctx *gin.Context) {
	agent, err := e.syncer.AgentGet(tenantId, plugId, agentId)
	if err != nil {
		e.logExecFileServe(tenantId, plugId, agentId, file, "error reteriving agent")
		return
	}

	builder, ok := e.builders[agent.Executor]
	if !ok {
		e.logExecFileServe(tenantId, plugId, agentId, file, "exec_builder not found")
		return
	}

	out, err := builder.ExecFile(file)
	if err != nil {
		e.logExecFileServe(tenantId, plugId, agentId, file, "file not found")
		return
	}

	e.writeFile(file, out, ctx)
}

func (e *Engine) writeFile(file string, data []byte, ctx *gin.Context) {

	ffiles := strings.Split(file, ".")

	switch ffiles[1] {
	case "js":
		ctx.Writer.Header().Set("Content-Type", "application/javascript")
	default:
		ctx.Writer.Header().Set("Content-Type", "text/css")
	}
	ctx.Writer.Write(data)
}

func (e *Engine) logExecFileServe(tenantId, plugId, agentId, file, message string) {
	e.logger.Debug().Str("tenant_id", tenantId).
		Str("plug_id", plugId).
		Str("agent_id", agentId).
		Str("action", "exec_file_serve").
		Str("file", file).Msg(message)
}

func (e *Engine) logFileServe(tenantId, plugId, agentId, file, message string) {
	e.logger.Debug().Str("tenant_id", tenantId).
		Str("plug_id", plugId).
		Str("agent_id", agentId).
		Str("action", "file_serve").
		Str("file", file).Msg(message)
}
