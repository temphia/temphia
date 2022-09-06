package apiadmin

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/instance"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
)

func (a *ApiAdmin) bprintAPI(rg *gin.RouterGroup) {

	rg.GET("/", a.X(a.BprintList))
	rg.POST("/", a.X(a.BprintCreate))
	rg.GET("/:id", a.X(a.BprintGet))
	rg.POST("/:id", a.X(a.BprintUpdate))
	rg.DELETE("/:id", a.X(a.BprintRemove))
	rg.GET("/:id/file", a.X(a.BprintListFiles))
	rg.GET("/:id/file/:file_id", a.X(a.BprintGetFile))
	rg.POST("/:id/file/:file_id", a.X(a.BprintNewBlob))
	rg.PATCH("/:id/file/:file_id", a.X(a.BprintUpdateBlob))
	rg.DELETE("/:id/file/:file_id", a.X(a.BprintDelFile))

	rg.POST("/:id/instance", a.X(a.BprintInstance))

	// adminApi.POST("/import_bprint", r.Authed(r.BprintImport))
	// adminApi.POST("/dev_plug_issue_tkt", r.Authed(r.DevIssuePlugTkt))

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
	bytes, err := io.ReadAll(ctx.Http.Request.Body)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.BprintUpdateBlob(ctx.Session, ctx.Http.Param("id"), ctx.Http.Param("file_id"), bytes)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) BprintGetFile(ctx httpx.Request) {
	out, err := r.cAdmin.BprintGetBlob(ctx.Session, ctx.Http.Param("id"), ctx.Http.Param("file_id"))
	if err != nil {
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
	opts := &instance.RepoOptions{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.BprintInstance(ctx.Session, opts)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

// import

func (c *ApiAdmin) BprintImport(ctx httpx.Request) {

	opts := &service.RepoImportOpts{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		return
	}

	resp, err := c.cAdmin.BprintImport(ctx.Session, opts)

	c.rutil.WriteJSON(ctx.Http, resp, err)
}
