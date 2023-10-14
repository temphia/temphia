package engine

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
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

func (r *RemoteHandler) Log(uclaim *claim.LSock, c *gin.Context) {
	b := r.engine.getBinding(uclaim.TenantId, uclaim.Plug, uclaim.Agent)

	out, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	b.Log(c.Param("eid"), kosher.Str(out))
}

func (r *RemoteHandler) LazyLog(c *gin.Context)       {}
func (r *RemoteHandler) GetSelfFile(c *gin.Context)   {}
func (r *RemoteHandler) ListResources(c *gin.Context) {}
func (r *RemoteHandler) GetResource(c *gin.Context)   {}
func (r *RemoteHandler) InLinks(c *gin.Context)       {}
func (r *RemoteHandler) OutLinks(c *gin.Context)      {}
func (r *RemoteHandler) LinkExec(c *gin.Context)      {}
func (r *RemoteHandler) LinkExecEmit(c *gin.Context)  {}
func (r *RemoteHandler) NewModule(c *gin.Context)     {}
func (r *RemoteHandler) ModuleTicket(c *gin.Context)  {}
func (r *RemoteHandler) ModuleExec(c *gin.Context)    {}
