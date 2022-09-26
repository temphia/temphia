package apiadmin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (a *ApiAdmin) userGroupAPI(rg *gin.RouterGroup) {

	rg.GET("/", a.X(a.ListUserGroup))
	rg.POST("/", a.X(a.AddUserGroup))
	rg.GET("/:ugroup", a.X(a.GetUserGroup))
	rg.POST("/:ugroup", a.X(a.UpdateUserGroup))
	rg.DELETE("/:ugroup", a.X(a.RemoveUserGroup))

	rg.GET("/:ugroup/data", a.X(a.ListUserGroupData))
	rg.POST("/:ugroup/data", a.X(a.AddUserGroupData))
	rg.GET("/:ugroup/data/:id", a.X(a.GetUserGroupData))
	rg.POST("/:ugroup/data/:id", a.X(a.UpdateUserGroupData))
	rg.DELETE("/:ugroup/data/:id", a.X(a.RemoveUserGroupData))

	rg.GET("/:ugroup/plug", a.X(a.ListUserGroupPlug))
	rg.POST("/:ugroup/plug", a.X(a.AddUserGroupPlug))
	rg.GET("/:ugroup/plug/:id", a.X(a.GetUserGroupPlug))
	rg.POST("/:ugroup/plug/:id", a.X(a.UpdateUserGroupPlug))
	rg.DELETE("/:ugroup/plug/:id", a.X(a.RemoveUserGroupPlug))

	rg.GET("/:ugroup/auth", a.X(a.ListUserGroupAuth))
	rg.POST("/:ugroup/auth", a.X(a.AddUserGroupAuth))
	rg.GET("/:ugroup/auth/:id", a.X(a.GetUserGroupAuth))
	rg.POST("/:ugroup/auth/:id", a.X(a.UpdateUserGroupAuth))
	rg.DELETE("/:ugroup/auth/:id", a.X(a.RemoveUserGroupAuth))

}

func (r *ApiAdmin) AddUserGroup(ctx httpx.Request) {
	group := &entities.UserGroup{}

	err := ctx.Http.BindJSON(group)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	group.TenantID = ctx.Session.TenentId

	r.rutil.WriteJSON(ctx.Http, nil, r.cAdmin.AddUserGroup(ctx.Session, group))
}

func (r *ApiAdmin) GetUserGroup(ctx httpx.Request) {
	resp, err := r.cAdmin.GetUserGroup(ctx.Session, ctx.Http.Param("ugroup"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) ListUserGroup(ctx httpx.Request) {
	resp, err := r.cAdmin.ListUserGroup(ctx.Session)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) UpdateUserGroup(ctx httpx.Request) {
	data := make(map[string]interface{})
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.UpdateUserGroup(ctx.Session, ctx.Http.Param("ugroup"), data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) RemoveUserGroup(ctx httpx.Request) {
	r.rutil.WriteJSON(ctx.Http, nil, r.cAdmin.RemoveUserGroup(ctx.Session, ctx.Http.Param("ugroup")))
}

// auth user group meta

func (r *ApiAdmin) AddUserGroupAuth(ctx httpx.Request) {
	data := &entities.UserGroupAuth{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.AddUserGroupAuth(ctx.Session, ctx.Http.Param("ugroup"), data),
	)
}

func (r *ApiAdmin) UpdateUserGroupAuth(ctx httpx.Request) {
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

	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.UpdateUserGroupAuth(ctx.Session, ctx.Http.Param("ugroup"), id, data),
	)
}

func (r *ApiAdmin) ListUserGroupAuth(ctx httpx.Request) {
	data, err := r.cAdmin.ListUserGroupAuth(ctx.Session, ctx.Http.Param("ugroup"))
	r.rutil.WriteJSON(ctx.Http, data, err)
}

func (r *ApiAdmin) GetUserGroupAuth(ctx httpx.Request) {

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupAuth(ctx.Session, ctx.Http.Param("ugroup"), id)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}
func (r *ApiAdmin) RemoveUserGroupAuth(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.RemoveUserGroupAuth(ctx.Session, ctx.Http.Param("ugroup"), id)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

// plug

func (r *ApiAdmin) AddUserGroupPlug(ctx httpx.Request) {
	data := &entities.UserGroupPlug{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.AddUserGroupPlug(ctx.Session, ctx.Http.Param("ugroup"), data),
	)
}

func (r *ApiAdmin) UpdateUserGroupPlug(ctx httpx.Request) {
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

	err = r.cAdmin.UpdateUserGroupPlug(ctx.Session, ctx.Http.Param("ugroup"), id, data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) ListUserGroupPlug(ctx httpx.Request) {
	resp, err := r.cAdmin.ListUserGroupPlug(ctx.Session, ctx.Http.Param("ugroup"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) GetUserGroupPlug(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupPlug(ctx.Session, ctx.Http.Param("ugroup"), id)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) RemoveUserGroupPlug(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.cAdmin.RemoveUserGroupPlug(ctx.Session, ctx.Http.Param("ugroup"), id)
}

// data

func (r *ApiAdmin) AddUserGroupData(ctx httpx.Request) {
	data := &entities.UserGroupData{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.AddUserGroupData(ctx.Session, ctx.Http.Param("ugroup"), data),
	)
}

func (r *ApiAdmin) UpdateUserGroupData(ctx httpx.Request) {
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

	err = r.cAdmin.UpdateUserGroupData(ctx.Session, ctx.Http.Param("ugroup"), id, data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) ListUserGroupData(ctx httpx.Request) {
	resp, err := r.cAdmin.ListUserGroupData(ctx.Session, ctx.Http.Param("ugroup"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) GetUserGroupData(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupData(ctx.Session, ctx.Http.Param("ugroup"), id)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) RemoveUserGroupData(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.cAdmin.RemoveUserGroupData(ctx.Session, ctx.Http.Param("ugroup"), id)
}
