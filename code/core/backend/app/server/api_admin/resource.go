package apiadmin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/vmodels"
)

func (r *ApiAdmin) ResourceCreate(ctx httpx.Request) {
	res := &entities.Resource{}
	err := ctx.Http.BindJSON(res)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	res.TenantId = ctx.Session.TenentId

	err = r.cAdmin.ResourceNew(ctx.Session, res)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) ResourceGet(ctx httpx.Request) {
	resp, err := r.cAdmin.ResourceGet(ctx.Session, ctx.Http.Param("slug"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) ResourceUpdate(ctx httpx.Request) {
	data := make(map[string]interface{})
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	err = r.cAdmin.ResourceUpdate(ctx.Session, ctx.Http.Param("slug"), data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) ResourceRemove(ctx httpx.Request) {
	err := r.cAdmin.ResourceDel(ctx.Session, ctx.Http.Param("slug"))
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) ResourceList(ctx httpx.Request) {
	plugId := ctx.Http.Query("plug_id")
	if plugId != "" {
		resp, err := r.cAdmin.ResourceListByPlug(ctx.Session, plugId)
		r.rutil.WriteJSON(ctx.Http, resp, err)
		return
	}

	resp, err := r.cAdmin.ResourceList(ctx.Session)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) ResourceAgentList(ctx httpx.Request) {
	query := &vmodels.ResourceQuery{}
	err := ctx.Http.BindJSON(&query)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.ResourceAgentList(ctx.Session, query)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}
