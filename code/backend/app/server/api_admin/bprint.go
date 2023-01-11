package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/core/backend/controllers/admin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/service/repox"
)

func (a *ApiAdmin) bprintAPI(rg *gin.RouterGroup) {

	rg.GET("/", a.X(a.BprintList))
	rg.POST("/", a.X(a.BprintCreate))
	rg.PUT("/", a.X(a.BprintImport))

	rg.GET("/:id", a.X(a.BprintGet))
	rg.POST("/:id", a.X(a.BprintUpdate))
	rg.DELETE("/:id", a.X(a.BprintRemove))
	rg.GET("/:id/file", a.X(a.BprintListFiles))
	rg.GET("/:id/file/:file_id", a.X(a.BprintGetFile))
	rg.POST("/:id/file/:file_id", a.X(a.BprintNewBlob))
	rg.PATCH("/:id/file/:file_id", a.X(a.BprintUpdateBlob))
	rg.DELETE("/:id/file/:file_id", a.X(a.BprintDelFile))

	rg.POST("/:id/instance", a.X(a.BprintInstance))
	rg.POST("/:id/issue", a.X(a.DevIssueTkt))
	rg.POST("/:id/issue/encoded", a.X(a.DevIssueTktEncoded))
}

func (r *ApiAdmin) BprintList(ctx httpx.Request) {

	rep, err := r.cAdmin.BprintList(ctx.Session, "")
	r.rutil.WriteJSON(ctx.Http, rep, err)
}

func (r *ApiAdmin) BprintCreate(ctx httpx.Request) {
	bprint := &entities.BPrint{}
	err := ctx.Http.BindJSON(bprint)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	id, err := r.cAdmin.BprintCreate(ctx.Session, bprint)
	r.rutil.WriteJSON(ctx.Http, id, err)
}

func (r *ApiAdmin) BprintGet(ctx httpx.Request) {
	resp, err := r.cAdmin.BprintGet(ctx.Session, ctx.Http.Param("id"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) BprintUpdate(ctx httpx.Request) {
	bprint := &entities.BPrint{}
	err := ctx.Http.BindJSON(bprint)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	err = r.cAdmin.BprintUpdate(ctx.Session, bprint)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) BprintRemove(ctx httpx.Request) {
	err := r.cAdmin.BprintRemove(ctx.Session, ctx.Http.Param("id"))
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) BprintListFiles(ctx httpx.Request) {
	resp, err := r.cAdmin.BprintListBlobs(ctx.Session, ctx.Http.Param("id"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) BprintNewBlob(ctx httpx.Request) {
	bytes, err := httpx.ReadForm(ctx.Http)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.BprintNewBlob(ctx.Session, ctx.Http.Param("id"), ctx.Http.Param("file_id"), bytes)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) BprintUpdateBlob(ctx httpx.Request) {
	pp.Println("@@@@@")

	bytes, err := httpx.ReadForm(ctx.Http)
	if err != nil {
		pp.Println("aaaa", err)
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.BprintUpdateBlob(ctx.Session, ctx.Http.Param("id"), ctx.Http.Param("file_id"), bytes)

	pp.Println(err)

	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) BprintGetFile(ctx httpx.Request) {
	out, err := r.cAdmin.BprintGetBlob(ctx.Session, ctx.Http.Param("id"), ctx.Http.Param("file_id"))
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	httpx.WriteBinary(ctx.Http, out)
}

func (r *ApiAdmin) BprintDelFile(ctx httpx.Request) {
	err := r.cAdmin.BprintDeleteBlob(ctx.Session, ctx.Http.Param("id"), ctx.Http.Param("file_id"))
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

// instance

func (r *ApiAdmin) BprintInstance(ctx httpx.Request) {
	opts := &admin.InstanceOptions{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.BprintInstance(ctx.Session, ctx.Http.Param("id"), opts)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

// import

func (r *ApiAdmin) BprintImport(ctx httpx.Request) {

	opts := &repox.RepoImportOpts{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		return
	}

	resp, err := r.cAdmin.BprintImport(ctx.Session, opts)

	r.rutil.WriteJSON(ctx.Http, resp, err)
}

// issue

func (r *ApiAdmin) DevIssueTkt(ctx httpx.Request) {
	rdata := admin.DevIssueReq{}

	err := ctx.Http.BindJSON(&rdata)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	rdata.BprintId = ctx.MustParam("id")
	tkt, err := r.cAdmin.DevIssueTktEncoded(ctx.Session, ctx.Http.Request.Host, rdata)

	r.rutil.WriteJSON(ctx.Http, tkt, err)
}

func (r *ApiAdmin) DevIssueTktEncoded(ctx httpx.Request) {
	rdata := admin.DevIssueReq{}

	err := ctx.Http.BindJSON(&rdata)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	rdata.BprintId = ctx.MustParam("id")

	tkt, err := r.cAdmin.DevIssueTktEncoded(ctx.Session, ctx.Http.Request.Host, rdata)

	r.rutil.WriteJSON(ctx.Http, tkt, err)
}
