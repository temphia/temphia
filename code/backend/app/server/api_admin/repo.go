package apiadmin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (a *ApiAdmin) repoAPI(rg *gin.RouterGroup) {

	rg.GET("/", a.X(a.RepoList))
	rg.POST("/", a.X(a.RepoNew))
	rg.GET("/:rid", a.X(a.RepoGet))
	rg.POST("/:rid", a.X(a.RepoUpdate))
	rg.DELETE("/:rid", a.X(a.RepoDelete))

}

func (r *ApiAdmin) RepoNew(ctx httpx.Request) {
	data := &entities.Repo{}

	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.RepoNew(ctx.Session, data)
	r.rutil.WriteFinal(ctx.Http, err)
}

func (r *ApiAdmin) RepoList(ctx httpx.Request) {
	res, err := r.cAdmin.RepoList(ctx.Session)
	r.rutil.WriteJSON(ctx.Http, res, err)
}

func (r *ApiAdmin) RepoGet(ctx httpx.Request) {

	rid, err := strconv.ParseInt(ctx.Http.Param("rid"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	res, err := r.cAdmin.RepoGet(ctx.Session, rid)
	r.rutil.WriteJSON(ctx.Http, res, err)
}

func (r *ApiAdmin) RepoUpdate(ctx httpx.Request) {

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

func (r *ApiAdmin) RepoDelete(ctx httpx.Request) {

	rid, err := strconv.ParseInt(ctx.Http.Param("rid"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.RepoDel(ctx.Session, rid)
	r.rutil.WriteFinal(ctx.Http, err)
}
