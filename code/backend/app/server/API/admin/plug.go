package apiadmin

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/admin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (a *ApiAdmin) plugAPI(rg *gin.RouterGroup) {

	rg.GET("/", a.X(a.ListPlug))
	rg.POST("/", a.X(a.NewPlug))
	rg.GET("/:plug_id", a.X(a.GetPlug))
	rg.POST("/:plug_id", a.X(a.UpdatePlug))
	rg.DELETE("/:plug_id", a.X(a.DelPlug))

	rg.GET("/:plug_id/agent/", a.X(a.ListAgent))
	rg.POST("/:plug_id/agent/", a.X(a.NewAgent))
	rg.GET("/:plug_id/agent/:agent_id", a.X(a.GetAgent))
	rg.POST("/:plug_id/agent/:agent_id", a.X(a.UpdateAgent))
	rg.DELETE("/:plug_id/agent/:agent_id", a.X(a.DelAgent))

	rg.GET("/:plug_id/state/", a.X(a.ListPlugState))
	rg.POST("/:plug_id/state/", a.X(a.NewPlugState))
	rg.GET("/:plug_id/state/:key", a.X(a.GetPlugState))
	rg.POST("/:plug_id/state/:key", a.X(a.UpdatePlugState))
	rg.DELETE("/:plug_id/state/:key", a.X(a.DelPlugState))

	rg.GET("/:plug_id/state_export", a.X(a.ExportPlugState))
	rg.POST("/:plug_id/state_import", a.X(a.ImportPlugState))

	rg.GET("/:plug_id/flowmap", a.X(a.PlugFlowmap))

	rg.GET("/:plug_id/agent/:agent_id/link", a.X(a.AgentLinkList))
	rg.POST("/:plug_id/agent/:agent_id/link", a.X(a.AgentLinkNew))
	rg.POST("/:plug_id/agent/:agent_id/link/:id", a.X(a.AgentLinkUpdate))
	rg.GET("/:plug_id/agent/:agent_id/link/:id", a.X(a.AgentLinkGet))
	rg.DELETE("/:plug_id/agent/:agent_id/link/:id", a.X(a.AgentLinkDel))

	rg.GET("/:plug_id/agent/:agent_id/extension", a.X(a.AgentExtensionList))
	rg.POST("/:plug_id/agent/:agent_id/extension", a.X(a.AgentExtensionNew))
	rg.POST("/:plug_id/agent/:agent_id/extension/:id", a.X(a.AgentExtensionUpdate))
	rg.GET("/:plug_id/agent/:agent_id/extension/:id", a.X(a.AgentExtensionGet))
	rg.DELETE("/:plug_id/agent/:agent_id/extension/:id", a.X(a.AgentExtensionDel))

	rg.GET("/:plug_id/agent/:agent_id/resource", a.X(a.AgentResourceList))
	rg.POST("/:plug_id/agent/:agent_id/resource", a.X(a.AgentResourceNew))
	rg.POST("/:plug_id/agent/:agent_id/resource/:id", a.X(a.AgentResourceUpdate))
	rg.GET("/:plug_id/agent/:agent_id/resource/:id", a.X(a.AgentResourceGet))
	rg.DELETE("/:plug_id/agent/:agent_id/resource/:id", a.X(a.AgentResourceDel))

	// 		adminApi.POST("/agent_resources", r.Authed(r.ResourceAgentList))

}

func (r *ApiAdmin) NewPlug(ctx httpx.Request) {
	data := &entities.Plug{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	data.TenantId = ctx.Session.TenantId

	err = r.cAdmin.PlugNew(ctx.Session, data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) UpdatePlug(ctx httpx.Request) {
	data := make(map[string]interface{})
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.PlugUpdate(ctx.Session, ctx.MustParam("plug_id"), data)
	r.rutil.WriteFinal(ctx.Http, err)

}

func (r *ApiAdmin) GetPlug(ctx httpx.Request) {
	plug, err := r.cAdmin.PlugGet(ctx.Session, ctx.MustParam("plug_id"))
	r.rutil.WriteJSON(ctx.Http, plug, err)
}
func (r *ApiAdmin) DelPlug(ctx httpx.Request) {
	err := r.cAdmin.PlugDel(ctx.Session, ctx.MustParam("plug_id"))
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) ListPlug(ctx httpx.Request) {
	bid := ctx.Http.Query("bprint_id")

	if bid != "" {
		pgs, err := r.cAdmin.PlugListByBprint(ctx.Session, bid)
		r.rutil.WriteJSON(ctx.Http, pgs, err)
		return
	}

	pgs, err := r.cAdmin.PlugList(ctx.Session)
	r.rutil.WriteJSON(ctx.Http, pgs, err)
}

func (r *ApiAdmin) NewAgent(ctx httpx.Request) {
	data := &entities.Agent{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	data.TenantId = ctx.Session.TenantId

	err = r.cAdmin.AgentNew(ctx.Session, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) UpdateAgent(ctx httpx.Request) {
	data := make(map[string]interface{})
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.AgentUpdate(ctx.Session, ctx.MustParam("plug_id"), ctx.MustParam("agent_id"), data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) GetAgent(ctx httpx.Request) {
	agent, err := r.cAdmin.AgentGet(ctx.Session, ctx.MustParam("plug_id"), ctx.MustParam("agent_id"))
	r.rutil.WriteJSON(ctx.Http, agent, err)
}

func (r *ApiAdmin) DelAgent(ctx httpx.Request) {
	err := r.cAdmin.AgentDel(ctx.Session, ctx.MustParam("plug_id"), ctx.MustParam("agent_id"))
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) ListAgent(ctx httpx.Request) {
	agents, err := r.cAdmin.AgentList(ctx.Session, ctx.MustParam("plug_id"))
	r.rutil.WriteJSON(ctx.Http, agents, err)
}

// plug state

func (r *ApiAdmin) NewPlugState(ctx httpx.Request) {
	data := admin.PlugStateNew{}
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.PlugStateNew(ctx.Session, ctx.MustParam("plug_id"), data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) UpdatePlugState(ctx httpx.Request) {
	data := admin.PlugStateUpdate{}
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.cAdmin.PlugStateUpdate(ctx.Session, ctx.MustParam("plug_id"), ctx.MustParam("key"), data)
}

func (r *ApiAdmin) GetPlugState(ctx httpx.Request) {
	key, err := r.cAdmin.PlugStateGet(ctx.Session, ctx.MustParam("plug_id"), ctx.MustParam("key"))
	r.rutil.WriteJSON(ctx.Http, key, err)
}

func (r *ApiAdmin) DelPlugState(ctx httpx.Request) {
	err := r.cAdmin.PlugStateDel(ctx.Session, ctx.MustParam("plug_id"), ctx.MustParam("key"))
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) ExportPlugState(ctx httpx.Request) {
	file, err := r.cAdmin.PlugKvExport(ctx.Session, ctx.MustParam("plug_id"))
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	ctx.Http.File(file)
	os.Remove(file)
}

func (r *ApiAdmin) ImportPlugState(ctx httpx.Request) {
	err := r.cAdmin.PlugKvImport(ctx.Session, ctx.MustParam("plug_id"), ctx.Http.Query("clear") == "true", ctx.Http.Request.Body)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) ListPlugState(ctx httpx.Request) {
	page, _ := strconv.ParseUint(ctx.Http.Query("page"), 10, 64)

	agents, err := r.cAdmin.PlugStateList(ctx.Session, ctx.MustParam("plug_id"), ctx.Http.Query("key_cursor"), uint(page))
	r.rutil.WriteJSON(ctx.Http, agents, err)
}

// link

func (r *ApiAdmin) AgentLinkNew(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")

	data := &entities.AgentLink{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	data.FromPlug = plugId
	data.FromAgent = agentId

	err = r.cAdmin.AgentLinkNew(ctx.Session, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) AgentLinkUpdate(ctx httpx.Request) {
	data := make(map[string]interface{})

	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")
	id, err := strconv.ParseInt(ctx.MustParam("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.AgentLinkUpdate(ctx.Session, plugId, agentId, id, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) AgentLinkGet(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")
	id, err := strconv.ParseInt(ctx.MustParam("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.AgentLinkGet(ctx.Session, plugId, agentId, id)
	r.rutil.WriteJSON(ctx.Http, resp, err)

}

func (r *ApiAdmin) AgentLinkDel(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")

	id, _ := strconv.ParseInt(ctx.MustParam("id"), 10, 64)

	err := r.cAdmin.AgentLinkDel(ctx.Session, plugId, agentId, id)
	r.rutil.WriteFinal(ctx.Http, err)

}

func (r *ApiAdmin) AgentLinkList(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")

	resp, err := r.cAdmin.AgentLinkList(ctx.Session, plugId, agentId)
	r.rutil.WriteJSON(ctx.Http, resp, err)

}

// extension

func (r *ApiAdmin) AgentExtensionNew(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")

	data := &entities.AgentExtension{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	data.Plug = plugId
	data.Agent = agentId

	err = r.cAdmin.AgentExtensionNew(ctx.Session, data)
	r.rutil.WriteFinal(ctx.Http, err)

}

func (r *ApiAdmin) AgentExtensionUpdate(ctx httpx.Request) {
	data := make(map[string]interface{})

	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")
	id, err := strconv.ParseInt(ctx.MustParam("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.AgentExtensionUpdate(ctx.Session, plugId, agentId, id, data)
	r.rutil.WriteFinal(ctx.Http, err)

}

func (r *ApiAdmin) AgentExtensionGet(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")
	id, err := strconv.ParseInt(ctx.MustParam("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.AgentExtensionGet(ctx.Session, plugId, agentId, id)
	r.rutil.WriteJSON(ctx.Http, resp, err)

}

func (r *ApiAdmin) AgentExtensionDel(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")

	id, err := strconv.ParseInt(ctx.MustParam("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	err = r.cAdmin.AgentExtensionDel(ctx.Session, plugId, agentId, id)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) AgentExtensionList(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")

	resp, err := r.cAdmin.AgentExtensionList(ctx.Session, plugId, agentId)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

// resource

func (r *ApiAdmin) AgentResourceNew(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")

	data := &entities.AgentResource{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	data.PlugId = plugId
	data.AgentId = agentId

	err = r.cAdmin.AgentResourceNew(ctx.Session, data)
	r.rutil.WriteFinal(ctx.Http, err)

}

func (r *ApiAdmin) AgentResourceUpdate(ctx httpx.Request) {
	data := make(map[string]interface{})

	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")
	id := ctx.MustParam("id")

	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.AgentResourceUpdate(ctx.Session, plugId, agentId, id, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) AgentResourceGet(ctx httpx.Request) {

	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")
	id := ctx.MustParam("id")

	resp, err := r.cAdmin.AgentResourceGet(ctx.Session, plugId, agentId, id)
	r.rutil.WriteJSON(ctx.Http, resp, err)

}

func (r *ApiAdmin) AgentResourceDel(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")
	id := ctx.MustParam("id")

	err := r.cAdmin.AgentResourceDel(ctx.Session, plugId, agentId, id)
	r.rutil.WriteFinal(ctx.Http, err)

}

func (r *ApiAdmin) AgentResourceList(ctx httpx.Request) {
	plugId := ctx.MustParam("plug_id")
	agentId := ctx.MustParam("agent_id")

	resp, err := r.cAdmin.AgentResourceList(ctx.Session, plugId, agentId)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) PlugFlowmap(ctx httpx.Request) {
	resp, err := r.cAdmin.PlugFlowmap2(ctx.Session, ctx.MustParam("plug_id"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}
