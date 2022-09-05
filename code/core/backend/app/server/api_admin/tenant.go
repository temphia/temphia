package apiadmin

import (
	"strconv"

	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

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

// widget

func (r *ApiAdmin) AddDomainWidget(ctx httpx.Request) {
	did, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	data := &entities.DomainWidget{}
	err = ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	data.DomainId = did
	err = r.cAdmin.AddDomainWidget(ctx.Session, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) UpdateDomainWidget(ctx httpx.Request) {
	wid, err := strconv.ParseInt(ctx.Http.Param("wid"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	data := map[string]interface{}{}

	ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.UpdateDomainWidget(ctx.Session, wid, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) GetDomainWidget(ctx httpx.Request) {
	wid, err := strconv.ParseInt(ctx.Http.Param("wid"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.GetDomainWidget(ctx.Session, wid)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) RemoveDomainWidget(ctx httpx.Request) {
	wid, err := strconv.ParseInt(ctx.Http.Param("wid"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.RemoveDomainWidget(ctx.Session, wid)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) ListDomainWidget(ctx httpx.Request) {
	did, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.ListDomainWidget(ctx.Session, did)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}
