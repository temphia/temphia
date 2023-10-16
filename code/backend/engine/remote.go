package engine

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

func (e *Engine) GetRemoteHandler() any {

	rh := &RemoteHandler{
		engine: e,
	}

	return rh
}

type RemoteHandler struct {
	engine *Engine
}

func (r *RemoteHandler) Log(uclaim *claim.RemoteExec, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	out, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	b.Log(c.Query("eid"), kosher.Str(out))
}

func (r *RemoteHandler) LazyLog(uclaim *claim.RemoteExec, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)
	logs := []string{}

	err := c.BindJSON(&logs)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	b.LazyLog(c.Query("eid"), logs)
}

func (r *RemoteHandler) GetSelfFile(uclaim *claim.RemoteExec, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	out, _, err := b.GetFileWithMeta(c.Param("name"))
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	c.Data(http.StatusOK, httpx.CtypeBin, out)
}

func (r *RemoteHandler) ListResources(uclaim *claim.RemoteExec, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	resp, err := b.ListResources()
	httpx.WriteJSON(c, resp, err)
}

func (r *RemoteHandler) GetResource(uclaim *claim.RemoteExec, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	resp, err := b.GetResource(c.Param("name"))
	httpx.WriteJSON(c, resp, err)
}

func (r *RemoteHandler) InLinks(uclaim *claim.RemoteExec, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	resp, err := b.InLinks()
	httpx.WriteJSON(c, resp, err)
}

func (r *RemoteHandler) OutLinks(uclaim *claim.RemoteExec, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	resp, err := b.OutLinks()
	httpx.WriteJSON(c, resp, err)
}

func (r *RemoteHandler) LinkExec(uclaim *claim.RemoteExec, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	out, err := io.ReadAll(c.Request.Body)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	resp, err := b.LinkExec(c.Param("name"), c.Param("method"), lazydata.NewJsonData(out))
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	out, err = resp.AsJsonBytes()
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	c.Data(http.StatusOK, httpx.CtypeBin, out)
}

func (r *RemoteHandler) LinkExecEmit(uclaim *claim.RemoteExec, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	out, err := io.ReadAll(c.Request.Body)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	err = b.LinkExecEmit(c.Param("name"), c.Param("method"), lazydata.NewJsonData(out))
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	httpx.WriteOk(c)
}

func (r *RemoteHandler) NewModule(uclaim *claim.RemoteExec, c *gin.Context) {

	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	out, err := io.ReadAll(c.Request.Body)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	mid, err := b.NewModule(c.Param("name"), lazydata.NewJsonData(out))
	httpx.WriteJSON(c, mid, err)

}

func (r *RemoteHandler) ModuleTicket(uclaim *claim.RemoteExec, c *gin.Context) {

	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	out, err := io.ReadAll(c.Request.Body)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	resp, err := b.ModuleTicket(c.Param("name"), lazydata.NewJsonData(out))
	httpx.WriteJSON(c, resp, err)

}

func (r *RemoteHandler) ModuleExec(uclaim *claim.RemoteExec, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	out, err := io.ReadAll(c.Request.Body)
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	mid, _ := strconv.ParseInt(c.Param("mid"), 10, 64)

	resp, err := b.ModuleExec(int32(mid), c.Param("name"), lazydata.NewJsonData(out))
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	out, err = resp.AsJsonBytes()
	if err != nil {
		httpx.WriteErr(c, err)
		return
	}

	c.Data(http.StatusOK, httpx.CtypeJSON, out)

}
