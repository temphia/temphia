package apiadmin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (a *ApiAdmin) plugAPI(rg *gin.RouterGroup) {

}

func (r *ApiAdmin) NewPlug(ctx httpx.Request) {
	data := &entities.Plug{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	data.TenantId = ctx.Session.TenentId

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

	r.cAdmin.PlugUpdate(ctx.Session, ctx.Http.Param("plug_id"), data)

}

func (r *ApiAdmin) GetPlug(ctx httpx.Request) {
	plug, err := r.cAdmin.PlugGet(ctx.Session, ctx.Http.Param("plug_id"))
	r.rutil.WriteJSON(ctx.Http, plug, err)
}
func (r *ApiAdmin) DelPlug(ctx httpx.Request) {
	err := r.cAdmin.PlugDel(ctx.Session, ctx.Http.Param("plug_id"))
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

	data.TenantId = ctx.Session.TenentId

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

	r.cAdmin.AgentUpdate(ctx.Session, ctx.Http.Param("plug_id"), ctx.Http.Param("agent_id"), data)
}

func (r *ApiAdmin) GetAgent(ctx httpx.Request) {
	agent, err := r.cAdmin.AgentGet(ctx.Session, ctx.Http.Param("plug_id"), ctx.Http.Param("agent_id"))
	r.rutil.WriteJSON(ctx.Http, agent, err)
}

func (r *ApiAdmin) DelAgent(ctx httpx.Request) {
	err := r.cAdmin.AgentDel(ctx.Session, ctx.Http.Param("plug_id"), ctx.Http.Param("agent_id"))
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) ListAgent(ctx httpx.Request) {
	agents, err := r.cAdmin.AgentList(ctx.Session, ctx.Http.Param("plug_id"))
	r.rutil.WriteJSON(ctx.Http, agents, err)
}

func (r *ApiAdmin) PairAgentToken(ctx httpx.Request) {

}

// link

func (r *ApiAdmin) AgentLinkNew(ctx httpx.Request) {
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")

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

	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
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
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.AgentLinkGet(ctx.Session, plugId, agentId, id)
	r.rutil.WriteJSON(ctx.Http, resp, err)

}

func (r *ApiAdmin) AgentLinkDel(ctx httpx.Request) {
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")

	id, _ := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)

	err := r.cAdmin.AgentLinkDel(ctx.Session, plugId, agentId, id)
	r.rutil.WriteFinal(ctx.Http, err)

}

func (r *ApiAdmin) AgentLinkList(ctx httpx.Request) {
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")

	resp, err := r.cAdmin.AgentLinkList(ctx.Session, plugId, agentId)
	r.rutil.WriteJSON(ctx.Http, resp, err)

}

// extension

func (r *ApiAdmin) AgentExtensionNew(ctx httpx.Request) {
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")

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

	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
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
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.AgentExtensionGet(ctx.Session, plugId, agentId, id)
	r.rutil.WriteJSON(ctx.Http, resp, err)

}

func (r *ApiAdmin) AgentExtensionDel(ctx httpx.Request) {
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	err = r.cAdmin.AgentExtensionDel(ctx.Session, plugId, agentId, id)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) AgentExtensionList(ctx httpx.Request) {
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")

	resp, err := r.cAdmin.AgentExtensionList(ctx.Session, plugId, agentId)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

// resource

func (r *ApiAdmin) AgentResourceNew(ctx httpx.Request) {
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")

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

	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")
	id := ctx.Http.Param("id")

	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.AgentResourceUpdate(ctx.Session, plugId, agentId, id, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) AgentResourceGet(ctx httpx.Request) {

	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")
	id := ctx.Http.Param("id")

	resp, err := r.cAdmin.AgentResourceGet(ctx.Session, plugId, agentId, id)
	r.rutil.WriteJSON(ctx.Http, resp, err)

}

func (r *ApiAdmin) AgentResourceDel(ctx httpx.Request) {
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")
	id := ctx.Http.Param("id")

	err := r.cAdmin.AgentResourceDel(ctx.Session, plugId, agentId, id)
	r.rutil.WriteFinal(ctx.Http, err)

}

func (r *ApiAdmin) AgentResourceList(ctx httpx.Request) {
	plugId := ctx.Http.Param("plug_id")
	agentId := ctx.Http.Param("agent_id")

	resp, err := r.cAdmin.AgentResourceList(ctx.Session, plugId, agentId)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}
