package apiadmin

import (
	"strconv"

	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (r *ApiAdmin) TenantRepoNew(ctx httpx.Request) {
	data := &entities.Repo{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.RepoNew(ctx.Session, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) TenantRepoList(ctx httpx.Request) {
	res, err := r.cAdmin.RepoList(ctx.Session)
	r.rutil.WriteJSON(ctx.Http, res, err)
}

func (r *ApiAdmin) TenantRepoGet(ctx httpx.Request) {

	rid, err := strconv.ParseInt(ctx.Http.Param("rid"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	res, err := r.cAdmin.RepoGet(ctx.Session, rid)
	r.rutil.WriteJSON(ctx.Http, res, err)
}

func (r *ApiAdmin) TenantRepoUpdate(ctx httpx.Request) {

	data := make(map[string]interface{})

	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	rid, err := strconv.ParseInt(ctx.Http.Param("rid"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.RepoUpdate(ctx.Session, rid, data)
	r.rutil.WriteFinal(ctx.Http, err)

}

func (r *ApiAdmin) TenantRepoDelete(ctx httpx.Request) {

	rid, err := strconv.ParseInt(ctx.Http.Param("rid"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.RepoDel(ctx.Session, rid)
	r.rutil.WriteFinal(ctx.Http, err)
}
