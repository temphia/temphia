package apiadmin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (a *ApiAdmin) tenantAPI(rg *gin.RouterGroup) {

	rg.POST("/", a.X(a.UpdateTenant))
	rg.GET("/", a.X(a.GetTenant))

	rg.GET("/domain", a.X(a.ListTenantDomain))
	rg.POST("/domain", a.X(a.AddTenantDomain))
	rg.GET("/domain/:id", a.X(a.GetTenantDomain))
	rg.POST("/domain/:id", a.X(a.UpdateTenantDomain))
	rg.DELETE("/domain/:id", a.X(a.RemoveTenantDomain))

}

func (r *ApiAdmin) UpdateTenant(ctx httpx.Request) {

	data := make(map[string]interface{})

	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.UpdateTenant(ctx.Session, data)
	httpx.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) GetTenant(ctx httpx.Request) {
	resp, err := r.cAdmin.GetTenant(ctx.Session)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

// domain

func (r *ApiAdmin) AddTenantDomain(ctx httpx.Request) {
	data := &entities.TenantDomain{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.AddDomain(ctx.Session, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) ListTenantDomain(ctx httpx.Request) {
	resp, err := r.cAdmin.ListDomain(ctx.Session)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) GetTenantDomain(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	resp, err := r.cAdmin.GetDomain(ctx.Session, id)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) UpdateTenantDomain(ctx httpx.Request) {
	data := map[string]interface{}{}
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.UpdateDomain(ctx.Session, id, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) RemoveTenantDomain(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.RemoveDomain(ctx.Session, id)
	r.rutil.WriteFinal(ctx.Http, err)
}
