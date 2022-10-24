package apiadmin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (a *ApiAdmin) TargetAPI(rg *gin.RouterGroup) {
	rg.POST("/app/:target_type", a.X(a.AddTargetApp))
	rg.GET("/app/:target_type", a.X(a.ListTargetAppByType))
	rg.GET("/app", a.X(a.ListTargetApp))
	rg.GET("/app/:target_type/:id", a.X(a.GetTargetApp))
	rg.POST("/app/:target_type/:id", a.X(a.UpdateTargetApp))
	rg.DELETE("/app/:target_type/:id", a.X(a.RemoveTargetApp))

	rg.POST("/hook/:target_type", a.X(a.AddTargetHook))
	rg.GET("/hook/:target_type", a.X(a.ListTargetHookByType))
	rg.GET("/hook", a.X(a.ListTargetHook))
	rg.GET("/hook/:target_type/:id", a.X(a.GetTargetHook))
	rg.POST("/hook/:target_type/:id", a.X(a.UpdateTargetHook))
	rg.DELETE("/hook/:target_type/:id", a.X(a.RemoveTargetHook))

}

// app

func (r *ApiAdmin) AddTargetApp(ctx httpx.Request) {
	data := &entities.TargetApp{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	data.TargetType = ctx.MustParam("target_type")
	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.AddTargetApp(ctx.Session, data),
	)

}

func (r *ApiAdmin) UpdateTargetApp(ctx httpx.Request) {
	data := make(map[string]interface{})
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

	err = r.cAdmin.UpdateTargetApp(ctx.Session, ctx.MustParam("target_type"), id, data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) ListTargetApp(ctx httpx.Request) {
	resp, err := r.cAdmin.ListTargetApp(ctx.Session)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) ListTargetAppByType(ctx httpx.Request) {
	resp, err := r.cAdmin.ListTargetAppByType(
		ctx.Session,
		ctx.MustParam("target_type"),
	)

	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) GetTargetApp(ctx httpx.Request) {

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	resp, err := r.cAdmin.GetTargetApp(ctx.Session, ctx.MustParam("target_type"), id)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) RemoveTargetApp(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	r.cAdmin.RemoveTargetApp(ctx.Session, ctx.MustParam("target_type"), id)
}

// hook

func (r *ApiAdmin) AddTargetHook(ctx httpx.Request) {
	data := &entities.TargetHook{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	data.TargetType = ctx.MustParam("target_type")
	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.AddTargetHook(ctx.Session, data),
	)

}

func (r *ApiAdmin) UpdateTargetHook(ctx httpx.Request) {
	data := make(map[string]interface{})
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

	err = r.cAdmin.UpdateTargetHook(ctx.Session, ctx.MustParam("target_type"), id, data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) ListTargetHook(ctx httpx.Request) {
	resp, err := r.cAdmin.ListTargetHook(ctx.Session)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) ListTargetHookByType(ctx httpx.Request) {
	resp, err := r.cAdmin.ListTargetHookByType(
		ctx.Session,
		ctx.MustParam("target_type"),
	)

	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) GetTargetHook(ctx httpx.Request) {

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	resp, err := r.cAdmin.GetTargetHook(ctx.Session, ctx.MustParam("target_type"), id)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) RemoveTargetHook(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	r.cAdmin.RemoveTargetHook(ctx.Session, ctx.MustParam("target_type"), id)
}
