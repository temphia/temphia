package apiadmin

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/instance"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/vmodels"
	"github.com/temphia/temphia/code/core/backend/xtypes/service"
)

func (a *ApiAdmin) bprintAPI(rg *gin.RouterGroup) {

	/*

		bAPI := adminApi.Group("/bprint")
		bAPI.GET("/", r.Authed(r.BprintList))
		bAPI.POST("/", r.Authed(r.BprintCreate))

		bAPI.GET("/:id", r.Authed(r.BprintGet))
		bAPI.POST("/:id", r.Authed(r.BprintUpdate))
		bAPI.DELETE("/:id", r.Authed(r.BprintRemove))

		bAPI.POST("/:id/install", r.Authed(r.BprintInstall)) // fixme => remove this all the way down

		bAPI.POST("/:id/instance", r.Authed(r.BprintInstance))

		bAPI.GET("/:id/file", r.Authed(r.BprintListFiles))
		bAPI.GET("/:id/file/:file_id", r.Authed(r.BprintGetFile))

		bAPI.POST("/:id/file/:file_id", r.Authed(r.BprintNewBlob))
		bAPI.PATCH("/:id/file/:file_id", r.Authed(r.BprintUpdateBlob))
		bAPI.DELETE("/:id/file/:file_id", r.Authed(r.BprintDelFile))

		adminApi.POST("/import_bprint", r.Authed(r.BprintImport))
		adminApi.POST("/dev_plug_issue_tkt", r.Authed(r.DevIssuePlugTkt))

		adminApi.GET("/check_slug/bprint/:bid", r.Authed(r.CheckBprint))
		adminApi.GET("/check_slug/plug/:pid", r.Authed(r.CheckPlug))
		adminApi.GET("/check_slug/data_group/:source/:gid", r.Authed(r.CheckDataGroup))
		adminApi.GET("/check_slug/data_table/:source/:gid/:tid", r.Authed(r.CheckDataTable))


	*/

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

func (r *ApiAdmin) BprintInstall(ctx httpx.Request) {
	opts := &vmodels.RepoInstallOpts{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.BprintInstall(ctx.Session, opts)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

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

func (c *ApiAdmin) BprintImport(ctx httpx.Request) {

	opts := &service.RepoImportOpts{}
	err := ctx.Http.BindJSON(opts)
	if err != nil {
		return
	}

	resp, err := c.cAdmin.BprintImport(ctx.Session, opts)

	c.rutil.WriteJSON(ctx.Http, resp, err)
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
