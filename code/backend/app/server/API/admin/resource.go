package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/models/vmodels"
	"github.com/temphia/temphia/code/backend/xtypes/xnotz/httpx"
)

func (a *ApiAdmin) resourceAPI(rg *gin.RouterGroup) {

	rg.GET("/", a.X(a.ResourceList))
	rg.POST("/", a.X(a.ResourceCreate))
	rg.GET("/:slug", a.X(a.ResourceGet))
	rg.POST("/:slug", a.X(a.ResourceUpdate))
	rg.DELETE("/:slug", a.X(a.ResourceRemove))

}

func (r *ApiAdmin) ResourceCreate(ctx httpx.Request) {
	res := &entities.Resource{}
	err := ctx.Http.BindJSON(res)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	res.TenantId = ctx.Session.TenantId

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
